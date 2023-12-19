package procedure

import (
	"errors"
	"fmt"

	"github.pitagora/pkg/datasource"
	"github.pitagora/pkg/node"
	"github.pitagora/pkg/transformer"
)

type NodeDependency struct {
	Type node.NodeType
	Name string
}

type Node struct {
	Type         node.NodeType
	Procedure    transformer.Procedure
	DataSource   datasource.DataSource
	Dependencies []NodeDependency
}

type DependencyTree map[string]*Node

func (dp *DependencyTree) Pop(key string) (*Node, bool) {
	value, exists := (*dp)[key]
	if exists {
		delete((*dp), key)
	}
	return value, exists
}

func buildNodeDependency(source string) (NodeDependency, error) {
	t, name := node.GetNameAndType(source)
	if len(string(t)) == 0 {
		return NodeDependency{}, errors.New(fmt.Sprintf("unknown procedure input: %s", source))
	}
	return NodeDependency{Type: t, Name: name}, nil
}

func buildDependencyTree(procedures []transformer.Procedure, sources []datasource.DataSource) (DependencyTree, error) {
	nodes := make(DependencyTree)

	// Create nodes for each procedure
	for _, proc := range procedures {
		nodes[proc.StepName] = &Node{Type: node.NODE_TYPE_PROCEDURE, Procedure: proc}
	}

	// Add dependencies
	for _, proc := range procedures {
		node := nodes[proc.StepName]

		if dataArgument, ok := proc.Arguments["data"]; ok {
			switch source := dataArgument.Value.(type) {
			case string:
				dep, err := buildNodeDependency(source)
				if nil != err {
					return nil, err
				}
				node.Dependencies = append(node.Dependencies, dep)
			case []interface{}:
				// Multiple dependencies
				for _, src := range source {
					strSrc := src.(string)
					dep, err := buildNodeDependency(strSrc)
					if nil != err {
						return nil, err
					}
					node.Dependencies = append(node.Dependencies, dep)
				}
			default:
				// TODO proper error
				panic("unknow data")
			}
		}
	}

	for _, source := range sources {
		if !source.DependsOnSomething() {
			continue
		}

		nodes[source.Name] = &Node{Type: node.NODE_TYPE_DATA, DataSource: source}
		node := nodes[source.Name]

		for _, d := range source.Depends {
			dep, err := buildNodeDependency(d.Value)
			if nil != err {
				return nil, err
			}
			node.Dependencies = append(node.Dependencies, dep)
		}
	}

	return nodes, nil
}
