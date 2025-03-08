package model

const (
	MessageTypeSwitchSpace = "SWITCH_SPACE"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}
