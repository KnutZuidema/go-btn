package model

import (
	"encoding/json"
	"strconv"
)

type Subscription struct {
	ID       int              `json:"Id"`
	Forum    string           `json:"Forum"`
	ForumID  int              `json:"ForumID"`
	LastPost SubscriptionPost `json:"LastPost"`
	NumPosts int              `json:"NumPosts"`
	NumViews int              `json:"NumViews"`
	Title    string           `json:"Title"`
	Unread   bool             `json:"Unread"`
}

type Subscriptions []Subscription

func (s *Subscriptions) UnmarshalJSON(data []byte) error {
	var raws rawSubscriptions
	if err := json.Unmarshal(data, &raws); err != nil {
		return err
	}
	var subs []Subscription
	for idStr, raw := range raws {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return err
		}
		forumID, err := strconv.Atoi(raw.ForumID)
		if err != nil {
			return err
		}
		posts, err := strconv.Atoi(raw.NumPosts)
		if err != nil {
			return err
		}
		views, err := strconv.Atoi(raw.NumViews)
		if err != nil {
			return err
		}
		unread := raw.Unread == "1"
		subs = append(subs, Subscription{
			ID:       id,
			Forum:    raw.Forum,
			ForumID:  forumID,
			LastPost: raw.LastPost,
			NumPosts: posts,
			NumViews: views,
			Title:    raw.Title,
			Unread:   unread,
		})
	}
	*s = subs
	return nil
}

type rawSubscriptions map[string]struct {
	Forum    string           `json:"Forum"`
	ForumID  string           `json:"ForumID"`
	LastPost SubscriptionPost `json:"LastPost"`
	NumPosts string           `json:"NumPosts"`
	NumViews string           `json:"NumViews"`
	Title    string           `json:"Title"`
	Unread   string           `json:"Unread"`
}
