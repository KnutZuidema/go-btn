package model

import (
	"encoding/json"
	"strconv"
)

type ForumThread struct {
	Author   string          `json:"Author"`
	ID       int             `json:"ID"`
	Locked   bool            `json:"Locked"`
	Sticky   bool            `json:"Sticky"`
	LastPost ForumThreadPost `json:"LastPost"`
	Title    string          `json:"Title"`
}

type rawForumThread struct {
	Author   string          `json:"Author"`
	ID       string          `json:"ID"`
	Locked   string          `json:"IsLocked"`
	Sticky   string          `json:"IsSticky"`
	LastPost ForumThreadPost `json:"LastPost"`
	Title    string          `json:"Title"`
}

func (t *ForumThread) UnmarshalJSON(data []byte) error {
	var raw rawForumThread
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	id, err := strconv.Atoi(raw.ID)
	if err != nil {
		return err
	}
	locked := raw.Locked != "0"
	sticky := raw.Sticky != "0"
	*t = ForumThread{
		Author:   raw.Author,
		ID:       id,
		Locked:   locked,
		Sticky:   sticky,
		LastPost: raw.LastPost,
		Title:    raw.Title,
	}
	return nil
}
