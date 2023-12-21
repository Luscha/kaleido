package automation

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.pitagora/pkg/action.go"
	"github.pitagora/pkg/storage"
	"gitlab.com/technity/go-x/pkg/connection"
	"gitlab.com/technity/go-x/pkg/message"
)

type AutomationManager struct {
	tenant     string
	conn       *connection.ConnectionManager[*storage.Client]
	automation Automation
}

func NewAutomationManager(tenant string, conn *connection.ConnectionManager[*storage.Client]) *AutomationManager {
	return &AutomationManager{
		tenant: tenant,
		conn:   conn,
		automation: Automation{
			Trigger:  Trigger{},
			Manifest: action.ActionRoot{},
		},
	}
}

func (t *AutomationManager) Load(ctx context.Context, name string) error {
	cl, err := t.conn.Borrow(ctx, t.tenant)
	if err != nil {
		panic(err)
	}
	defer t.conn.Release(ctx, t.tenant)
	automation, err := cl.GetAutomation(ctx, name)
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(automation.Trigger), &t.automation.Trigger); err != nil {
		panic(err)
	}

	if err := json.Unmarshal([]byte(automation.Manifest), &t.automation.Manifest); err != nil {
		panic(err)
	}
	return nil
}

func (t *AutomationManager) SetAutomation(ctx context.Context, auto Automation) {
	t.automation = auto
}

func (t *AutomationManager) Enable(ctx context.Context) error {
	switch t.automation.Trigger.Type {
	case TRIGGER_TYPE_CRON:
		{
			ser, err := json.Marshal(t.automation.Trigger.Manifest)
			if err != nil {
				panic(err)
			}

			var manifest TriggerCron
			err = json.Unmarshal(ser, &manifest)
			if err != nil {
				panic(err)
			}

			cronName := fmt.Sprintf("%s_%s", t.tenant, t.automation.Trigger.Name)
			additionalLoadFn := make([]func(o *config.LoadOptions) error, 0)
			if os.Getenv("AWS_ACCESS_KEY_ID") != "" && os.Getenv("AWS_SECRET_ACCESS_KEY") != "" {
				additionalLoadFn = append(additionalLoadFn,
					config.WithCredentialsProvider(
						credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
					),
				)
			}

			additionalLoadFn = append(additionalLoadFn, config.WithRegion(os.Getenv("AWS_REGION")))

			clientSettings, err := config.LoadDefaultConfig(ctx, additionalLoadFn...)
			if err != nil {
				panic("sqs configuration error, " + err.Error())
			}

			// Create EventBridge client
			ebClient := eventbridge.NewFromConfig(clientSettings)

			cronExpression := fmt.Sprintf("cron(%s)", manifest.Expression)
			snsTopicArn := os.Getenv("SNS_PUBLISH_TOPIC")
			payload := message.NewCommandMessage(
				message.NewCommandTMN("kaleido", "automation"),
				message.NewTRN("me", "me", "me", "me"),
				message.WithArguments(
					message.NewCommandMessageArgument(t.automation.Manifest),
				),
			)

			// Create CloudWatch Events rule
			createRuleOutput, err := ebClient.PutRule(ctx, &eventbridge.PutRuleInput{
				Name:               aws.String(cronName),
				ScheduleExpression: aws.String(cronExpression),
				State:              types.RuleStateEnabled,
			})
			if err != nil {
				log.Fatalf("failed to create rule, %v", err)
			}
			fmt.Printf("Rule ARN: %s\n", *createRuleOutput.RuleArn)

			// Prepare SNS payload
			payloadBytes, err := json.Marshal(payload)
			if err != nil {
				log.Fatalf("failed to marshal payload, %v", err)
			}
			payloadString := string(payloadBytes)

			// Create target for the rule
			_, err = ebClient.PutTargets(ctx, &eventbridge.PutTargetsInput{
				Rule: aws.String(cronName),
				Targets: []types.Target{
					{
						Arn:   aws.String(snsTopicArn),
						Id:    aws.String(cronName),
						Input: aws.String(payloadString),
						// ! NO: the role must exist before
						// to allow event bridge to send message, we need to change sns's access policy:
						/*
							{
							      "Sid": "AllowEventBridge",
							      "Effect": "Allow",
							      "Principal": {
							        "Service": "events.amazonaws.com"
							      },
							      "Action": "SNS:Publish",
							      "Resource": "arn:aws:sns:eu-west-1:466786686423:local-bridge"
							}
						*/

						//RoleArn: aws.String("arn:aws:iam::466786686423:role/test/test-eventbridge-sns"),
					},
				},
			})
			if err != nil {
				log.Fatalf("failed to put targets, %v", err)
			}
			fmt.Println("Target set for rule")
		}
	}
	return nil
}

func (t *AutomationManager) Disable(ctx context.Context) error {
	switch t.automation.Trigger.Type {
	case TRIGGER_TYPE_CRON:
		{
			cronName := fmt.Sprintf("%s_%s", t.tenant, t.automation.Trigger.Name)
			additionalLoadFn := make([]func(o *config.LoadOptions) error, 0)
			if os.Getenv("AWS_ACCESS_KEY_ID") != "" && os.Getenv("AWS_SECRET_ACCESS_KEY") != "" {
				additionalLoadFn = append(additionalLoadFn,
					config.WithCredentialsProvider(
						credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
					),
				)
			}

			additionalLoadFn = append(additionalLoadFn, config.WithRegion(os.Getenv("AWS_REGION")))

			clientSettings, err := config.LoadDefaultConfig(ctx, additionalLoadFn...)
			if err != nil {
				panic("sqs configuration error, " + err.Error())
			}

			// Create EventBridge client
			ebClient := eventbridge.NewFromConfig(clientSettings)

			// Remove target from the rule
			_, err = ebClient.RemoveTargets(ctx, &eventbridge.RemoveTargetsInput{
				Rule: aws.String(cronName),
				Ids:  []string{cronName}, // ID of the target to remove
			})
			if err != nil {
				log.Fatalf("failed to remove targets, %v", err)
			}
			log.Println("Target removed from rule")

			// Delete the rule
			_, err = ebClient.DeleteRule(ctx, &eventbridge.DeleteRuleInput{
				Name: aws.String(cronName),
			})
			if err != nil {
				log.Fatalf("failed to delete rule, %v", err)
			}
			log.Println("Rule deleted")
		}
	}
	return nil
}
