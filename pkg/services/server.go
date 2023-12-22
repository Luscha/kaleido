package services

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.pitagora/pkg/action.go"
	"github.pitagora/pkg/automation"
	"github.pitagora/pkg/procedure"
	"github.pitagora/pkg/storage"
	"gitlab.com/technity/go-x/pkg/connection"
	"gitlab.com/technity/go-x/pkg/logger"
	"gitlab.com/technity/go-x/pkg/tracing"
	"gitlab.com/technity/go-x/pkg/xhttp"
)

type ServerConfig struct {
	Port string
}

type Server struct {
	r    *gin.Engine
	cfg  *ServerConfig
	conn *connection.ConnectionManager[*storage.Client]
}

func NewServer(ctx context.Context, cfg *ServerConfig, conn *connection.ConnectionManager[*storage.Client]) *Server {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/"),
		gin.Recovery(),
	)

	// tracing
	r.Use(tracing.GinTracing())

	// logger
	r.Use(func(c *gin.Context) {
		ctx := c.Request.Context()
		log := logger.New(
			logger.WithService("vigile"),
			logger.WithTracingId(tracing.GetTracing(ctx)),
			logger.WithMinLevel(os.Getenv(logger.LOGGER_LEVEL_ENV)),
		)
		c.Request = c.Request.WithContext(logger.WithLogger(ctx, log))
		c.Next()
	})

	// client
	r.Use(func(c *gin.Context) {
		ctx := c.Request.Context()
		client := xhttp.NewClient(xhttp.WithAmzTracingId(tracing.GetTracing(ctx)))
		c.Request = c.Request.WithContext(xhttp.WithHttpClient(ctx, client))
		c.Next()
	})

	server := &Server{
		r:    r,
		cfg:  cfg,
		conn: conn,
	}

	server.r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, ResponseType, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	if err := server.buildRouter(ctx); err != nil {
		logrus.Panicf("error building routes the server: %v", err)
	}

	return server
}

func (s *Server) Close() error {
	return nil
}

func (s *Server) buildRouter(ctx context.Context) error {
	s.r.GET("/", s.HealthCheck)
	s.r.POST("/procedure", s.Procedure)
	s.r.POST("/action", s.Action)
	s.r.POST("/automation/:name", s.EnableAutomation)
	s.r.DELETE("/automation/:name", s.DisableAutomation)
	s.r.GET("/automation/list", s.ListAutomations)
	s.r.GET("/macro/list", s.ListMacros)
	s.r.GET("/procedure/list", s.ListProcedures)
	return nil
}

func (s *Server) Run(ctx context.Context, port string) error {
	return s.r.Run(port)
}

func (s *Server) HealthCheck(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (s *Server) Procedure(c *gin.Context) {
	ctx := c.Request.Context()

	var request procedure.Root
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	orc := procedure.NewOrchestrator("root", "tenant-abc06d5-28d8-45a3-a272-f577db014f67", s.conn)
	res, err := orc.Run(ctx, request)

	logger.GetLogger(ctx).WithFields(map[string]any{
		// "res": res,
		"err": err,
	}).Info("done")

	c.Data(http.StatusOK, "application/json", res.Result)
}

func (s *Server) Action(c *gin.Context) {
	ctx := c.Request.Context()

	var request action.ActionRoot
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	ah := action.NewActionHandler("tenant-abc06d5-28d8-45a3-a272-f577db014f67", s.conn)
	err := ah.Run(ctx, request)

	logger.GetLogger(ctx).WithFields(map[string]any{
		// "res": res,
		"err": err,
	}).Info("done")

	c.Status(http.StatusOK)
}

func (s *Server) ListMacros(c *gin.Context) {
	ctx := c.Request.Context()
	co, err := s.conn.Borrow(ctx, "tenant-abc06d5-28d8-45a3-a272-f577db014f67")
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	macros, err := co.ListMacros(ctx)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, macros)
}

func (s *Server) ListProcedures(c *gin.Context) {
	ctx := c.Request.Context()
	co, err := s.conn.Borrow(ctx, "tenant-abc06d5-28d8-45a3-a272-f577db014f67")
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	proc, err := co.ListProcedures(ctx)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, proc)
}

func (s *Server) ListAutomations(c *gin.Context) {
	ctx := c.Request.Context()
	co, err := s.conn.Borrow(ctx, "tenant-abc06d5-28d8-45a3-a272-f577db014f67")
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	automs := make([]map[string]any, 0)
	raw, err := co.ListAutomations(ctx)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	for _, a := range raw {
		var manifest action.ActionRoot
		if err := json.Unmarshal([]byte(a.Manifest), &manifest); err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		var trigger automation.Trigger
		if err := json.Unmarshal([]byte(a.Trigger), &trigger); err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		// rawManifest, err := json.Marshal(map[string]any{
		// 	"trigger":        trigger,
		// 	"action":         manifest.Actions,
		// 	"arguments":      manifest.Arguments,
		// 	"data":           manifest.Procedure.Data,
		// 	"procedure":      manifest.Procedure.Procedure,
		// 	"real_procedure": manifest.Procedure.SubProcedure,
		// })
		// if err != nil {
		// 	c.Status(http.StatusInternalServerError)
		// 	return
		// }
		autom := map[string]any{}
		autom["id"] = a.ID
		autom["name"] = a.Name
		autom["description"] = a.Description
		autom["manifest"] = map[string]any{
			"trigger":        trigger,
			"action":         manifest.Actions,
			"arguments":      manifest.Arguments,
			"data":           manifest.Procedure.Data,
			"procedure":      manifest.Procedure.Procedure,
			"real_procedure": manifest.Procedure.SubProcedure,
		}
		automs = append(automs, autom)
	}

	c.JSON(http.StatusOK, automs)
}
