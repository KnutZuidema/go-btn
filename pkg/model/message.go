package model

import (
	"encoding/json"
	"strconv"
	"time"
)

type Message struct {
	Date     time.Time `json:"Date"`
	ID       int       `json:"ID"`
	Sticky   bool      `json:"Sticky"`
	Subject  string    `json:"Subject"`
	Unread   bool      `json:"Unread"`
	UserID   int       `json:"UserID"`
	Username string    `json:"Username"`
}

type rawMessage struct {
	Date     string `json:"Date"`
	ID       string `json:"ID"`
	Sticky   string `json:"Sticky"`
	Subject  string `json:"Subject"`
	Unread   string `json:"Unread"`
	UserID   string `json:"UserID"`
	Username string `json:"Username"`
}

func (m *Message) UnmarshalJSON(data []byte) error {
	var raw rawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	ts, err := strconv.Atoi(raw.Date)
	if err != nil {
		return err
	}
	date := time.Unix(int64(ts), 0)
	id, err := strconv.Atoi(raw.ID)
	if err != nil {
		return err
	}
	sticky := raw.Sticky != "0"
	unread := raw.Unread != "0"
	userid, err := strconv.Atoi(raw.UserID)
	if err != nil {
		return err
	}
	*m = Message{
		Date:     date,
		ID:       id,
		Sticky:   sticky,
		Subject:  raw.Subject,
		Unread:   unread,
		UserID:   userid,
		Username: raw.Username,
	}
	return nil
}
