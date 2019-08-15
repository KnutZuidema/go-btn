package model

import (
	"encoding/json"
	"strconv"
	"time"
)

type News []NewsItem

type rawNews map[string]struct {
	Author string `json:"Author"`
	Body   string `json:"Body"`
	Time   int    `json:"Time"`
	Title  string `json:"Title"`
}

func (n *News) UnmarshalJSON(data []byte) error {
	var raw rawNews
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	var items []NewsItem
	for idStr, item := range raw {
		t := time.Unix(int64(item.Time), 0)
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return err
		}
		items = append(items, NewsItem{
			Author: item.Author,
			Body:   item.Body,
			Time:   t,
			Title:  item.Title,
			ID:     id,
		})
	}
	*n = items
	return nil
}
