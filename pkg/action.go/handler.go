package action

import (
	"context"
	"fmt"
	"sync"

	"github.pitagora/pkg/node"
	"github.pitagora/pkg/storage"
	"github.pitagora/pkg/template"
	"gitlab.com/technity/go-x/pkg/connection"
	"gitlab.com/technity/go-x/pkg/message"
	"gitlab.com/technity/go-x/pkg/sns"
)

type ActionHandler struct {
	tenant string
	conn   *connection.ConnectionManager[*storage.Client]
}

func NewActionHandler(tenant string, conn *connection.ConnectionManager[*storage.Client]) *ActionHandler {
	return &ActionHandler{
		tenant: tenant,
		conn:   conn,
	}
}

func (a *ActionHandler) Run(ctx context.Context, root ActionRoot) error {
	// orc := procedure.NewOrchestrator("root", a.tenant, a.conn)
	// root.Procedure.Arguments = root.Arguments
	// res, err := orc.Run(ctx, root.Procedure)
	// if err != nil {
	// 	return err
	// }

	results := sync.Map{}
	// results.Store(node.TypeAndStringKey(node.NODE_TYPE_PROCEDURE, res.Name), res.Result)
	results.Store(node.TypeAndStringKey(node.NODE_TYPE_PROCEDURE, "root"), []byte(`["a", "b", "c"]`))

	for _, action := range root.Actions {
		actualAction := action
		fullArgs, err := MergeArgumentsForActions(actualAction, root.Arguments, &results)
		if nil != err {
			panic(err)
		}
		prefix := template.ArgumentPrefix(node.NODE_TYPE_ACTION, action.Name)
		err = template.Resolve(actualAction, fullArgs, prefix, &actualAction)
		if nil != err {
			panic(err)
		}

		err = a.Do(ctx, actualAction)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *ActionHandler) Do(ctx context.Context, action Action) error {
	publisher := sns.Factory.Spawn(ctx)
	msg := message.NewCommandMessage(
		message.NewCommandTMN("test", "email"),
		message.NewTRN("me", "me", "me", "me"),
		message.WithArguments(
			message.NewCommandMessageArgument(action.Manifest),
		),
	)
	return publisher.Publish(ctx, sns.SNSIngressTopic, msg)
}

func MergeArgumentsForActions(a Action, globalArguments template.Arguments, intermediateRes *sync.Map) (template.Arguments, error) {
	merged := template.Arguments{}
	for key, value := range globalArguments {
		merged[key] = value
	}

	for _, dep := range a.Depends {
		specialName := node.TypeAndStringKey(node.GetNameAndType(dep.Value))
		res, ok := intermediateRes.Load(specialName)
		if !ok {
			return template.Arguments{}, fmt.Errorf("%s not found in results", specialName)
		}

		merged[fmt.Sprintf("%s%s", template.ArgumentPrefix(node.NODE_TYPE_ACTION, a.Name), dep.Template)] = string(res.([]byte))
	}
	return merged, nil
}
