package model

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type NewsItem struct {
	Author string    `json:"Author"`
	Body   string    `json:"Body"`
	Time   time.Time `json:"Time"`
	Title  string    `json:"Title"`
	ID     int       `json:"ID"`
}

func (i *NewsItem) UnmarshalJSON(data []byte) error {
	var raws rawNews
	if err := json.Unmarshal(data, &raws); err != nil {
		return err
	}
	if len(raws) != 1 {
		return fmt.Errorf("not a single news item")
	}
	for idStr, raw := range raws {
		t := time.Unix(int64(raw.Time), 0)
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return err
		}
		*i = NewsItem{
			Author: raw.Author,
			Body:   raw.Body,
			Time:   t,
			Title:  raw.Title,
			ID:     id,
		}
	}
	return nil
}
