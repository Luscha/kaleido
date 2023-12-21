package procedure

import (
	"context"
	"encoding/json"
	"sync"

	"github.pitagora/pkg/datasource"
	"github.pitagora/pkg/node"
	"github.pitagora/pkg/storage"
	"github.pitagora/pkg/template"
	"github.pitagora/pkg/transformer"
	"gitlab.com/technity/go-x/pkg/connection"
)

type Root struct {
	Data         []datasource.DataSource `json:"data"`
	Procedure    []transformer.Procedure `json:"procedure"`
	SubProcedure []SubProcedure          `json:"real_procedure"`
	Arguments    template.Arguments      `json:"arguments"`
}

type Bus struct {
	Type node.NodeType
	Name string
	data []byte
}

type Orchestrator struct {
	name   string
	tenant string
	conn   *connection.ConnectionManager[*storage.Client]
}

func NewOrchestrator(name, tenant string, conn *connection.ConnectionManager[*storage.Client]) *Orchestrator {
	return &Orchestrator{
		name:   name,
		tenant: tenant,
		conn:   conn,
	}
}

type ProcedureResult struct {
	Name   string
	Result []byte
	// TODO error
}

func (o *Orchestrator) Run(ctx context.Context, procedure Root) (ProcedureResult, error) {
	depTree, err := buildDependencyTree(procedure.Procedure, procedure.Data)
	if nil != err {
		panic(err)
	}

	var wg sync.WaitGroup
	dataCh := make(chan datasource.Result, len(procedure.Data))
	transCh := make(chan transformer.Result, len(procedure.Procedure))
	subprocedureCh := make(chan ProcedureResult, 1)
	subOrchestrator := make([]*Orchestrator, 0)
	transformers := make([]*transformer.MacroHandler, 0)
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
				err := template.Resolve(data, procedure.Arguments, template.ArgumentPrefix(node.NODE_TYPE_DATA, data.Name), &actualData)
				if nil != err {
					panic(err)
				}
			}
			datasource.FetchChrono(ctx, actualData, dataCh)
		}(data)
	}

	for _, data := range procedure.SubProcedure {
		// TODO dependencies
		// if data.DependsOnSomething() {
		// 	continue
		// }

		orc := NewOrchestrator(data.Name, o.tenant, o.conn)
		subOrchestrator = append(subOrchestrator, orc)
		go func(procedureName string, ochestrator *Orchestrator) {
			// fetch procedure
			cl, err := o.conn.Borrow(ctx, o.tenant)
			if err != nil {
				panic(err)
			}
			defer o.conn.Release(ctx, o.tenant)
			proc, err := cl.GetProcedure(ctx, procedureName)
			if err != nil {
				panic(err)
			}
			var procManifest Root
			err = json.Unmarshal([]byte(proc.Manifest), &procManifest)
			if err != nil {
				panic(err)
			}
			// inject arguments
			procManifest.Arguments = procedure.Arguments.GetArgumentSubprocedure(procedureName)
			res, err := orc.Run(ctx, procManifest)
			if err != nil {
				panic(err)
			}
			subprocedureCh <- res
		}(data.Reference, orc)
	}

	// add to waitgroup all procedures
	wg.Add(len(procedure.Data))
	wg.Add(len(procedure.Procedure))
	wg.Add(len(procedure.SubProcedure))

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
		for data := range subprocedureCh {
			// TODO erro handling
			// if data.Err != nil {
			// 	panic(data)
			// }
			bus <- Bus{Type: node.NODE_TYPE_SUB_PROCEDURE, Name: data.Name, data: data.Result}
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
						err := template.Resolve(actualProcedure, procedure.Arguments, prefix, &actualProcedure)
						if nil != err {
							panic(err)
						}
					}

					h := transformer.NewMacroHandler(o.conn, o.tenant)
					transformers = append(transformers, h)
					go h.Transform(ctx, actualProcedure, &results, transCh)
				} else if n.Type == node.NODE_TYPE_DATA {
					actualData := n.DataSource
					fullArgs, err := template.MergeArgumentsForData(actualData, procedure.Arguments, &results)
					if nil != err {
						panic(err)
					}
					prefix := template.ArgumentPrefix(node.NODE_TYPE_DATA, actualData.Name)
					err = template.Resolve(actualData, fullArgs, prefix, &actualData)
					if nil != err {
						panic(err)
					}
					datasource.FetchChrono(ctx, actualData, dataCh)
				}
				// TODO subprocedure dependencies
			}
		}
	}()

	wg.Wait()
	close(dataCh)
	close(transCh)
	close(bus)
	close(subprocedureCh)

	return ProcedureResult{Name: o.name, Result: res}, nil
}

func checkDependencies(dataMap *sync.Map, n Node) bool {
	for _, dep := range n.Dependencies {
		if _, exists := dataMap.Load(node.TypeAndStringKey(dep.Type, dep.Name)); !exists {
			return false
		}
	}
	return true
}
