package automation

import "github.pitagora/pkg/action.go"

type Automation struct {
	Trigger  Trigger           `json:"trigger"`
	Manifest action.ActionRoot `json:"manifest"`
}
