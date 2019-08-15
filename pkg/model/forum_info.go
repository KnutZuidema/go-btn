package model

import (
	"encoding/json"
	"strconv"
)

type ForumInfo struct {
	GroupName    string `json:"GroupName"`
	MinClassRead int    `json:"MinClassRead"`
	Name         string `json:"Name"`
	TopicCount   int    `json:"Topics"`
}

type rawForumInfo struct {
	GroupName    string `json:"GroupName"`
	MinClassRead string `json:"MinClassRead"`
	Name         string `json:"Name"`
	Topics       string `json:"Topics"`
}

func (i *ForumInfo) UnmarshalJSON(data []byte) error {
	var raw rawForumInfo
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	minread, err := strconv.Atoi(raw.MinClassRead)
	if err != nil {
		return err
	}
	topics, err := strconv.Atoi(raw.Topics)
	if err != nil {
		return err
	}
	*i = ForumInfo{
		GroupName:    raw.GroupName,
		MinClassRead: minread,
		Name:         raw.Name,
		TopicCount:   topics,
	}
	return nil
}
