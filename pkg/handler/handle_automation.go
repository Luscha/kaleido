package handler

import (
	"context"

	"github.pitagora/pkg/action.go"
	"gitlab.com/technity/go-x/pkg/logger"
	"gitlab.com/technity/go-x/pkg/message"
	"gitlab.com/technity/go-x/pkg/sqs"
)

func (handler *Handler) handleAutomation(ctx context.Context, data *sqs.Payload) (*sqs.HandleResult, bool) {
	if data.CommandMessage.Tmn.GetMethod() != "automation" {
		return nil, false
	}

	var commandMessage message.CommandMessage[*action.ActionRoot]
	if err := message.CastCommand(data.CommandMessage, &commandMessage); err != nil {
		logger.GetLogger(ctx).WithError(err).Error("failed to unmarshal command message payload")
		return &sqs.HandleResult{Error: err}, true
	}

	ah := action.NewActionHandler("tenant-abc06d5-28d8-45a3-a272-f577db014f67", handler.conn)
	// return ws.CheckIntegrationJobs(ctx, data.EventMessage.Trn.GetParent(), nil), true
	ah.Run(ctx, *commandMessage.Arguments.Presets)
	return &sqs.HandleResult{}, true
}
