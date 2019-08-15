package model

import (
	"encoding/json"
	"strconv"
	"time"
)

type ConversationMessage struct {
	ID       int       `json:"ID"`
	Body     string    `json:"Body"`
	Time     time.Time `json:"Time"`
	Username string    `json:"Username"`
}

type Conversation struct {
	Messages []ConversationMessage `json:"Messages"`
	Sticky   bool                  `json:"Sticky"`
	Subject  string                `json:"Subject"`
	Unread   bool                  `json:"Unread"`
}

type rawConversation struct {
	Conversation map[string]struct {
		Body     string `json:"Body"`
		Time     string `json:"Time"`
		Username string `json:"Username"`
	} `json:"Conversation"`
	Sticky  string `json:"Sticky"`
	Subject string `json:"Subject"`
	Unread  string `json:"UnRead"`
}

func (c *Conversation) UnmarshalJSON(data []byte) error {
	var raw rawConversation
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	sticky := raw.Sticky != "0"
	unread := raw.Unread != "0"
	var messages []ConversationMessage
	for idStr, values := range raw.Conversation {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return err
		}
		ts, err := strconv.Atoi(values.Time)
		if err != nil {
			return err
		}
		t := time.Unix(int64(ts), 0)
		messages = append(messages, ConversationMessage{
			ID:       id,
			Body:     values.Body,
			Time:     t,
			Username: values.Username,
		})
	}
	*c = Conversation{
		Messages: messages,
		Sticky:   sticky,
		Subject:  raw.Subject,
		Unread:   unread,
	}
	return nil
}
