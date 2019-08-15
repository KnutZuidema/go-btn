package model

import (
	"encoding/json"
	"strconv"
)

type Forum struct {
	Description string    `json:"Description"`
	ForumID     int       `json:"ForumID"`
	ForumName   string    `json:"ForumName"`
	LastPost    ForumPost `json:"LastPost"`
	PostCount   int       `json:"NumPosts"`
	TopicCount  int       `json:"NumTopics"`
	Sort        int       `json:"Sort"`
}

type rawForum struct {
	Description string    `json:"Description"`
	ForumID     string    `json:"ForumID"`
	ForumName   string    `json:"ForumName"`
	LastPost    ForumPost `json:"LastPost"`
	PostCount   string    `json:"NumPosts"`
	TopicCount  string    `json:"NumTopics"`
	Sort        string    `json:"Sort"`
}

func (f *Forum) UnmarshalJSON(data []byte) error {
	var raw rawForum
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	forumID, err := strconv.Atoi(raw.ForumID)
	if err != nil {
		return err
	}
	posts, err := strconv.Atoi(raw.PostCount)
	if err != nil {
		return err
	}
	topics, err := strconv.Atoi(raw.TopicCount)
	if err != nil {
		return err
	}
	sort, err := strconv.Atoi(raw.Sort)
	if err != nil {
		return err
	}
	*f = Forum{
		Description: raw.Description,
		ForumID:     forumID,
		ForumName:   raw.ForumName,
		LastPost:    raw.LastPost,
		PostCount:   posts,
		TopicCount:  topics,
		Sort:        sort,
	}
	return nil
}
