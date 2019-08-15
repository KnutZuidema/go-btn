package model

import (
	"encoding/json"
	"strconv"
	"time"
)

type ForumThreadPost struct {
	Author string    `json:"Author"`
	ID     int       `json:"ID"`
	Time   time.Time `json:"Time"`
}

type rawForumThreadPost struct {
	Author string `json:"Author"`
	ID     string `json:"ID"`
	Time   int    `json:"Time"`
}

func (p *ForumThreadPost) UnmarshalJSON(data []byte) error {
	var raw rawForumThreadPost
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	id, err := strconv.Atoi(raw.ID)
	if err != nil {
		return err
	}
	t := time.Unix(int64(raw.Time), 0)
	*p = ForumThreadPost{
		Author: raw.Author,
		ID:     id,
		Time:   t,
	}
	return nil
}
