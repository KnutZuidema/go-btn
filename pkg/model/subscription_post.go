package model

import (
	"encoding/json"
	"strconv"
	"time"
)

type SubscriptionPost struct {
	AuthorID int       `json:"AuthorID"`
	ID       int       `json:"ID"`
	Time     time.Time `json:"Time"`
	Username string    `json:"Username"`
}

func (p *SubscriptionPost) UnmarshalJSON(data []byte) error {
	var raw rawSubscriptionPost
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	authorID, err := strconv.Atoi(raw.AuthorID)
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(raw.ID)
	if err != nil {
		return err
	}
	t := time.Unix(int64(raw.Time), 0)
	*p = SubscriptionPost{
		AuthorID: authorID,
		ID:       id,
		Time:     t,
		Username: raw.Username,
	}
	return nil
}

type rawSubscriptionPost struct {
	AuthorID string `json:"AuthorID"`
	ID       string `json:"ID"`
	Time     int    `json:"Time"`
	Username string `json:"Username"`
}
