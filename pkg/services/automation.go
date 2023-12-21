package services

import (
	"github.com/gin-gonic/gin"
	"github.pitagora/pkg/automation"
)

func (s *Server) EnableAutomation(c *gin.Context) {
	ctx := c.Request.Context()
	manager := automation.NewAutomationManager("tenant-abc06d5-28d8-45a3-a272-f577db014f67", s.conn)
	err := manager.Load(ctx, c.Param("name"))
	if err != nil {
		panic(err)
	}

	err = manager.Enable(ctx)
}

func (s *Server) DisableAutomation(c *gin.Context) {
	ctx := c.Request.Context()
	manager := automation.NewAutomationManager("tenant-abc06d5-28d8-45a3-a272-f577db014f67", s.conn)
	err := manager.Load(ctx, c.Param("name"))
	if err != nil {
		panic(err)
	}

	err = manager.Disable(ctx)
}
