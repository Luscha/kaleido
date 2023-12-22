package handler

import (
	"context"

	"github.pitagora/pkg/storage"
	"gitlab.com/technity/go-x/pkg/connection"
	"gitlab.com/technity/go-x/pkg/sns"
	"gitlab.com/technity/go-x/pkg/sqs"
	"gitlab.com/technity/go-x/pkg/xerrors"
)

type Handler struct {
	conn *connection.ConnectionManager[*storage.Client]
}

func NewHandler(conn *connection.ConnectionManager[*storage.Client]) *Handler {
	return &Handler{
		conn: conn,
	}
}

func Router(handler *Handler) sqs.Router {
	r := sqs.NewRouterImpl(
		sqs.WithExtractor(func(ctx context.Context, data []byte, result *sqs.Payload) error {
			command, event, err := sqs.MessageExtractor[any](ctx, nil, data)
			if err != nil {
				return err
			}
			result.CommandMessage = command
			result.EventMessage = event
			return nil
		}),
	)

	commandGr := sqs.NewHandleGroup(
		func(ctx context.Context, data *sqs.Payload) bool {
			return data.CommandMessage != nil
		},
	)

	eventGr := sqs.NewHandleGroup(
		func(ctx context.Context, data *sqs.Payload) bool {
			return data.EventMessage != nil
		},
	)

	commandGr.AddHandlerFunctions(
		handler.handleAutomation,
	)

	eventGr.AddHandlerFunctions()

	r.AddHandleGroups(
		commandGr,
		eventGr,
	)

	r.AddPostHook(func(ctx context.Context, result *sqs.HandleResult) error {
		publisher := sns.Factory.Spawn(ctx)

		for _, message := range result.Messages {
			if err := publisher.BatchMessage(ctx, sns.SNSIngressTopic, message); err != nil {
				return xerrors.NewRetriableError(err)
			}
		}
		if err := publisher.BatchFlush(ctx); err != nil {
			return xerrors.NewRetriableError(err)
		}
		return nil
	})
	return r
}
