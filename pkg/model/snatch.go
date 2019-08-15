package model

import (
	"encoding/json"
	"strconv"
	"time"
)

type Snatch struct {
	ID          int           `json:"ID"`
	Downloaded  int           `json:"Downloaded"`
	Uploaded    int           `json:"Uploaded"`
	IsSeeding   bool          `json:"IsSeeding"`
	Ratio       float64       `json:"Ratio"`
	SeedTime    time.Duration `json:"SeedTime"`
	SnatchTime  time.Time     `json:"SnatchTime"`
	TorrentID   int           `json:"TorrentId"`
	TorrentInfo TorrentInfo   `json:"TorrentInfo"`
}

func (s *Snatch) UnmarshalJSON(data []byte) error {
	var raw rawSnatch
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	down, err := strconv.Atoi(raw.Downloaded)
	if err != nil {
		return err
	}
	up, err := strconv.Atoi(raw.Uploaded)
	if err != nil {
		return err
	}
	isSeeding := raw.IsSeeding == "1"
	ratio, err := strconv.ParseFloat(raw.Ratio, 64)
	if err != nil {
		return err
	}
	seedTime, err := strconv.Atoi(raw.SeedTime)
	if err != nil {
		return err
	}
	torrentID, err := strconv.Atoi(raw.TorrentID)
	if err != nil {
		return err
	}
	snatchTime, err := time.Parse("2006-01-02 15:04:05", raw.SnatchTime)
	*s = Snatch{
		Downloaded:  down,
		Uploaded:    up,
		IsSeeding:   isSeeding,
		Ratio:       ratio,
		SeedTime:    time.Duration(seedTime) * time.Second,
		SnatchTime:  snatchTime,
		TorrentID:   torrentID,
		TorrentInfo: raw.TorrentInfo,
	}
	return nil
}

type rawSnatch struct {
	Downloaded  string      `json:"Downloaded"`
	Uploaded    string      `json:"Uploaded"`
	IsSeeding   string      `json:"IsSeeding"`
	Ratio       string      `json:"Ratio"`
	SeedTime    string      `json:"SeedTime"`
	SnatchTime  string      `json:"SnatchTime"`
	TorrentID   string      `json:"TorrentId"`
	TorrentInfo TorrentInfo `json:"TorrentInfo"`
}
