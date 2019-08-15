package model

import (
	"encoding/json"
	"time"
)

type ChangelogEntry struct {
	Changes string    `json:"Changes"`
	Time    time.Time `json:"Time"`
	User    string    `json:"User"`
}

func (c *ChangelogEntry) UnmarshalJSON(data []byte) error {
	var raw rawChangelogEntry
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02 15:04:05", raw.Time)
	if err != nil {
		return err
	}
	*c = ChangelogEntry{
		Changes: raw.Changes,
		Time:    t,
		User:    raw.User,
	}
	return nil
}

type rawChangelogEntry struct {
	Changes string `json:"Changes"`
	Time    string `json:"Time"`
	User    string `json:"User"`
}
