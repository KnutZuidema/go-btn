package model

type Inbox struct {
	MaxPages int       `json:"MaxPages"`
	Messages []Message `json:"results"`
}
