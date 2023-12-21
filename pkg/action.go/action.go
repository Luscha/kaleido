package action

import (
	"github.pitagora/pkg/procedure"
	"github.pitagora/pkg/template"
)

// ! TODO conditionals to manage when an action is triggered
type ActionType string

const (
	ACTION_TYPE_EMAIL ActionType = "email"
)

type Action struct {
	Type     ActionType      `json:"type"`
	Name     string          `json:"name"`
	Manifest any             `json:"manifest"`
	Depends  []ActionDepends `json:"depends"`
}

type ActionDepends struct {
	Value    string `json:"value"`
	Template string `json:"template"`
}

type ActionRoot struct {
	Procedure procedure.Root     `json:"procedure"`
	Actions   []Action           `json:"action"`
	Arguments template.Arguments `json:"arguments"`
}
