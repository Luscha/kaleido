package automation

type TriggerType string

const (
	TRIGGER_TYPE_CRON TriggerType = "CRON"
)

type Trigger struct {
	Type     TriggerType `json:"type"`
	Name     string      `json:"name"`
	Manifest any         `json:"manifest"`
}

type TriggerCron struct {
	Expression string `json:"expression"`
}
