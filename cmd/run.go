package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.pitagora/pkg/handler"
	python "github.pitagora/pkg/python3"
	"github.pitagora/pkg/services"
	"github.pitagora/pkg/storage"
	"gitlab.com/technity/go-x/pkg/connection"
	"gitlab.com/technity/go-x/pkg/endpoints"
	"gitlab.com/technity/go-x/pkg/logger"
	"gitlab.com/technity/go-x/pkg/message"
	"gitlab.com/technity/go-x/pkg/secret_manager"
	"gitlab.com/technity/go-x/pkg/sns"
	"gitlab.com/technity/go-x/pkg/sqs"
	"gitlab.com/technity/go-x/pkg/tracing"
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

			connectionCfg := &connection.ConnectionConfig{
				Username: "",
				Password: "",
				Port:     5432,
				Engine:   "postgres",
				Host:     os.Getenv("DATABASE_ENDPOINT"),
			}

			scrtmgrConfig := &secret_manager.SecretManagerConfig{
				Region:    os.Getenv("AWS_REGION"),
				KeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
				KeySecret: os.Getenv("AWS_SECRET_ACCESS_KEY"),
			}

			scrtMngr := secret_manager.NewSecretManager(ctx, scrtmgrConfig)

			dbSecretCredArn := os.Getenv("DATABASE_SECRET_ARN_CRED")
			dbSecretHostArn := os.Getenv("DATABASE_SECRET_ARN_HOST")
			if dbSecretCredArn != "" && dbSecretHostArn != "" {
				connectionCredCfg := &connection.ConnectionConfig{}
				connectionHostCfg := &connection.ConnectionConfig{}

				if err := scrtMngr.GetSecret(ctx, dbSecretCredArn, &connectionCredCfg); err != nil {
					logger.Main.WithError(err).Panic("failed to get database cred connection secret")
				}
				if err := scrtMngr.GetSecret(ctx, dbSecretHostArn, &connectionHostCfg); err != nil {
					logger.Main.WithError(err).Panic("failed to get database host connection secret")
				}
				connectionCfg.Engine = connectionHostCfg.Engine
				connectionCfg.Host = connectionHostCfg.Host
				connectionCfg.Port = connectionHostCfg.Port
				connectionCfg.Username = connectionCredCfg.Username
				connectionCfg.Password = connectionCredCfg.Password
			}

			dbClientFactory := storage.NewClientFactoy()

			ctx := context.Background()
			// ctx, cancel := context.WithCancel(ctx)
			// intr := interrupt.New(func(os.Signal) {}, cancel)

			conn := connection.NewConnectionManager(ctx,
				&connection.ConnectionManagerConfig[*storage.Client]{
					Config:            connectionCfg,
					ConnectionFactory: dbClientFactory,
				})

			// sns
			publishCfg := &sns.PublishConfig{
				Region:    os.Getenv("AWS_REGION"),
				KeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
				KeySecret: os.Getenv("AWS_SECRET_ACCESS_KEY"),
				TopicArns: map[string]string{
					sns.SNSIngressTopic: os.Getenv("SNS_PUBLISH_TOPIC"),
				},
			}

			sns.NewFactory(publishCfg)

			// endpoints
			endpoints.Init(endpoints.Config{
				Maya: os.Getenv("MAYA_URL"),
			})

			// api server
			serverCfg := &services.ServerConfig{
				Port: fmt.Sprintf("0.0.0.0:%s", os.Getenv("API_SERVER_PORT")),
			}

			server := services.NewServer(ctx, serverCfg, conn)

			fmt.Println("running")
			g, ctx := errgroup.WithContext(ctx)
			g.Go(func() error { return server.Run(ctx, serverCfg.Port) })

			// sqs
			h := handler.NewHandler(conn)

			if os.Getenv("SQS_ENDPOINT") != "" {
				pullCfg := &sqs.PullConfig{
					Region:    os.Getenv("AWS_REGION"),
					KeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
					KeySecret: os.Getenv("AWS_SECRET_ACCESS_KEY"),
					PollerConfig: &sqs.PollerConfig{
						SQSEndpoint:           os.Getenv("SQS_ENDPOINT"),
						MaxNumberOfMessages:   10,
						MessageAttributeNames: []string{tracing.MESSAGE_TRACING_ATTRIBUTE},
					},
					MaxConcurrency: 1,
				}

				poller := sqs.NewPollerPool(ctx, pullCfg, handler.Router(h))
				// middleware to inject logger
				poller.UseMiddleware(func(ctx context.Context, message sqs.SQSMessage) context.Context {
					log := logger.New(
						logger.WithService("vigile"),
						logger.WithTracingId(tracing.GetTracing(ctx)),
						logger.WithMinLevel(os.Getenv(logger.LOGGER_LEVEL_ENV)),
					)

					return logger.WithLogger(ctx, log)
				})

				g.Go(func() error { return poller.Pull(ctx) })
			}

			if err := g.Wait(); err != nil {
				logger.Main.WithError(err).Panic("failed to run")
			}
		},
	}

	rootCmd.AddCommand(runCmd)
}
