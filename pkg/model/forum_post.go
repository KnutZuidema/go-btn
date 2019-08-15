package model

import (
	"encoding/json"
	"strconv"
	"time"
)

type ForumPost struct {
	ID       int       `json:"ID"`
	Time     time.Time `json:"Time"`
	Title    string    `json:"Title"`
	TopicID  int       `json:"TopicID"`
	UserID   int       `json:"UserID"`
	Username string    `json:"Username"`
}

type rawForumPost struct {
	ID       string `json:"ID"`
	Time     int    `json:"Time"`
	Title    string `json:"Title"`
	TopicID  string `json:"TopicID"`
	UserID   string `json:"UserID"`
	Username string `json:"Username"`
}

func (p *ForumPost) UnmarshalJSON(data []byte) error {
	var raw rawForumPost
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	id, err := strconv.Atoi(raw.ID)
	if err != nil {
		return err
	}
	t := time.Unix(int64(raw.Time), 0)
	topicID, err := strconv.Atoi(raw.TopicID)
	if err != nil {
		return err
	}
	userID, err := strconv.Atoi(raw.UserID)
	if err != nil {
		return err
	}
	*p = ForumPost{
		ID:       id,
		Time:     t,
		Title:    raw.Title,
		TopicID:  topicID,
		UserID:   userID,
		Username: raw.Username,
	}
	return nil
}
