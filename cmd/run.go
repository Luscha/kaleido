package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	python "github.pitagora/pkg/python3"
	"github.pitagora/pkg/services"
	"gitlab.com/technity/go-x/pkg/endpoints"
	"gitlab.com/technity/go-x/pkg/logger"
	"gitlab.com/technity/go-x/pkg/message"
	"golang.org/x/sync/errgroup"
)

func init() {
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Runs the server",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(python.PyGILState_Check())
			fmt.Println(python.PyEval_ThreadsInitialized())
			fmt.Println(python.Py_IsInitialized())
			// state := python.PyGILState_Ensure()
			// defer python.PyGILState_Release(state)

			logger.NewMainLogger(
				logger.WithService("vigile"),
				logger.WithMinLevel(os.Getenv(logger.LOGGER_LEVEL_ENV)),
			)

			message.InitializeDefaultServiceTrnProcessId("vigiles")

			// endpoints
			endpoints.Init(endpoints.Config{
				Maya: os.Getenv("MAYA_URL"),
			})

			// api server
			serverCfg := &services.ServerConfig{
				Port: fmt.Sprintf("0.0.0.0:%s", os.Getenv("API_SERVER_PORT")),
			}

			server := services.NewServer(ctx, serverCfg)

			fmt.Println("running")
			g, ctx := errgroup.WithContext(ctx)
			g.Go(func() error { return server.Run(ctx, serverCfg.Port) })

			if err := g.Wait(); err != nil {
				logger.Main.WithError(err).Panic("failed to run")
			}
		},
	}

	rootCmd.AddCommand(runCmd)
}
