package procedure

import (
	"context"
	"sync"

	"github.pitagora/pkg/datasource"
	"github.pitagora/pkg/node"
	"github.pitagora/pkg/template"
	"github.pitagora/pkg/transformer"
)

type Root struct {
	Data      []datasource.DataSource `json:"data"`
	Procedure []transformer.Procedure `json:"procedure"`
	Arguments template.Arguments      `json:"arguments"`
}

type Bus struct {
	Type node.NodeType
	Name string
	data []byte
}

type Orchestrator struct{}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{}
}

func (o *Orchestrator) Run(ctx context.Context, procedure Root) ([]byte, error) {
	depTree, err := buildDependencyTree(procedure.Procedure, procedure.Data)
	if nil != err {
		panic(err)
	}

	var wg sync.WaitGroup
	dataCh := make(chan datasource.Result, len(procedure.Data))
	transCh := make(chan transformer.Result, len(procedure.Procedure))
	bus := make(chan Bus, 1)
	results := sync.Map{}
	var res []byte

	for _, data := range procedure.Data {
		if data.DependsOnSomething() {
			continue
		}

		go func(data datasource.DataSource) {
			actualData := data
			if procedure.Arguments.HasArguments(node.NODE_TYPE_DATA, data.Name) {
				template.Resolve(data, procedure.Arguments, template.ArgumentPrefix(node.NODE_TYPE_DATA, data.Name), &actualData)
			}
			datasource.FetchChrono(ctx, actualData, dataCh)
		}(data)
	}

	// add to waitgroup all procedures
	wg.Add(len(procedure.Data))
	wg.Add(len(procedure.Procedure))

	// Process data sources and start transformations
	go func() {
		for data := range dataCh {
			if data.Err != nil {
				// TODO erro handling
				panic(data)
			}
			res = data.Body
			bus <- Bus{Type: node.NODE_TYPE_DATA, Name: data.ID, data: data.Body}
		}
	}()

	go func() {
		for data := range transCh {
			if data.Err != nil {
				// TODO erro handling
				panic(data)
			}
			bus <- Bus{Type: node.NODE_TYPE_PROCEDURE, Name: data.ID, data: data.Data}
		}
	}()

	go func() {
		for b := range bus {
			defer wg.Done()

			results.Store(node.TypeAndStringKey(b.Type, b.Name), b.data)
			res = b.data
			// if no more items in dep map -> we are done
			if len(depTree) == 0 {
				return
			}

			// check remaining transformers
			nextSteps := []string{}
			for key, node := range depTree {
				if checkDependencies(&results, *node) {
					nextSteps = append(nextSteps, key)
				}
			}

			for _, step := range nextSteps {
				n, _ := depTree.Pop(step)
				if n.Type == node.NODE_TYPE_PROCEDURE {
					actualProcedure := n.Procedure
					if procedure.Arguments.HasArguments(node.NODE_TYPE_PROCEDURE, actualProcedure.StepName) {
						prefix := template.ArgumentPrefix(node.NODE_TYPE_PROCEDURE, actualProcedure.StepName)
						template.Resolve(actualProcedure, procedure.Arguments, prefix, &actualProcedure)
					}
					go transformer.Transform(ctx, actualProcedure, &results, transCh)
				} else if n.Type == node.NODE_TYPE_DATA {
					actualData := n.DataSource
					fullArgs, err := template.MergeArgumentsForData(actualData, procedure.Arguments, &results)
					if nil != err {
						panic(err)
					}
					prefix := template.ArgumentPrefix(node.NODE_TYPE_DATA, actualData.Name)
					template.Resolve(actualData, fullArgs, prefix, &actualData)
					datasource.FetchChrono(ctx, actualData, dataCh)
				}
			}
		}
	}()

	wg.Wait()
	close(dataCh)
	close(transCh)
	close(bus)

	return res, nil
}

func checkDependencies(dataMap *sync.Map, n Node) bool {
	for _, dep := range n.Dependencies {
		if _, exists := dataMap.Load(node.TypeAndStringKey(dep.Type, dep.Name)); !exists {
			return false
		}
	}
	return true
}
