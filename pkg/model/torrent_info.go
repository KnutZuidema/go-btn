package model

import (
	"encoding/json"
	"strconv"
)

type TorrentInfo struct {
	Codec      Codec      `json:"Codec"`
	Container  Container  `json:"Container"`
	GroupName  string     `json:"GroupName"`
	Resolution Resolution `json:"Resolution"`
	Series     string     `json:"Series"`
	Source     Source     `json:"Source"`
	Year       int        `json:"Year"`
}

func (i *TorrentInfo) UnmarshalJSON(data []byte) error {
	var raw rawTorrentInfo
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	year, err := strconv.Atoi(raw.Year)
	if err != nil {
		return err
	}
	*i = TorrentInfo{
		Codec:      raw.Codec,
		Container:  raw.Container,
		GroupName:  raw.GroupName,
		Resolution: raw.Resolution,
		Series:     raw.Series,
		Source:     raw.Source,
		Year:       year,
	}
	return nil
}

type rawTorrentInfo struct {
	Codec      Codec      `json:"Codec"`
	Container  Container  `json:"Container"`
	GroupName  string     `json:"GroupName"`
	Resolution Resolution `json:"Resolution"`
	Series     string     `json:"Series"`
	Source     Source     `json:"Source"`
	Year       string     `json:"Year"`
}
