package model

import (
	"encoding/json"
	"strconv"
	"time"
)

type Torrent struct {
	Category       string    `json:"Category"`
	Codec          string    `json:"Codec"`
	Container      string    `json:"Container"`
	DownloadURL    string    `json:"DownloadURL"`
	GroupID        int       `json:"GroupID"`
	GroupName      string    `json:"GroupName"`
	ImdbID         int       `json:"ImdbID"`
	InfoHash       string    `json:"InfoHash"`
	Leechers       int       `json:"Leechers"`
	Origin         string    `json:"Origin"`
	ReleaseName    string    `json:"ReleaseName"`
	Resolution     string    `json:"Resolution"`
	Seeders        int       `json:"Seeders"`
	Series         string    `json:"Series"`
	Seriesbanner   string    `json:"Seriesbanner"`
	SeriesID       int       `json:"SeriesID"`
	SeriesPoster   string    `json:"SeriesPoster"`
	Size           int       `json:"Size"`
	Snatched       int       `json:"Snatched"`
	Source         string    `json:"Source"`
	Time           time.Time `json:"Time"`
	TorrentID      int       `json:"TorrentID"`
	TvdbID         int       `json:"TvdbID"`
	TvrageID       int       `json:"TvrageID"`
	YoutubeTrailer string    `json:"YoutubeTrailer"`
}

type rawTorrent struct {
	Category       string `json:"Category"`
	Codec          string `json:"Codec"`
	Container      string `json:"Container"`
	DownloadURL    string `json:"DownloadURL"`
	GroupID        string `json:"GroupID"`
	GroupName      string `json:"GroupName"`
	ImdbID         string `json:"ImdbID"`
	InfoHash       string `json:"InfoHash"`
	Leechers       string `json:"Leechers"`
	Origin         string `json:"Origin"`
	ReleaseName    string `json:"ReleaseName"`
	Resolution     string `json:"Resolution"`
	Seeders        string `json:"Seeders"`
	Series         string `json:"Series"`
	Seriesbanner   string `json:"Seriesbanner"`
	SeriesID       string `json:"SeriesID"`
	SeriesPoster   string `json:"SeriesPoster"`
	Size           string `json:"Size"`
	Snatched       string `json:"Snatched"`
	Source         string `json:"Source"`
	Time           string `json:"Time"`
	TorrentID      string `json:"TorrentID"`
	TvdbID         string `json:"TvdbID"`
	TvrageID       string `json:"TvrageID"`
	YoutubeTrailer string `json:"YoutubeTrailer"`
}

func (t *Torrent) UnmarshalJSON(data []byte) error {
	var raw rawTorrent
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	groupID, err := strconv.Atoi(raw.GroupID)
	if err != nil {
		return err
	}
	imdbID, err := strconv.Atoi(raw.ImdbID)
	if err != nil {
		return err
	}
	leechers, err := strconv.Atoi(raw.Leechers)
	if err != nil {
		return err
	}
	seeders, err := strconv.Atoi(raw.Seeders)
	if err != nil {
		return err
	}
	seriesID, err := strconv.Atoi(raw.SeriesID)
	if err != nil {
		return err
	}
	size, err := strconv.Atoi(raw.Size)
	if err != nil {
		return err
	}
	snatched, err := strconv.Atoi(raw.Snatched)
	if err != nil {
		return err
	}
	ts, err := strconv.Atoi(raw.Time)
	if err != nil {
		return err
	}
	date := time.Unix(int64(ts), 0)
	torrentID, err := strconv.Atoi(raw.TorrentID)
	if err != nil {
		return err
	}
	tvdbID, err := strconv.Atoi(raw.TvdbID)
	if err != nil {
		return err
	}
	tvrageID, err := strconv.Atoi(raw.TvrageID)
	if err != nil {
		return err
	}
	*t = Torrent{
		Category:       raw.Category,
		Codec:          raw.Codec,
		Container:      raw.Container,
		DownloadURL:    raw.DownloadURL,
		GroupID:        groupID,
		GroupName:      raw.GroupName,
		ImdbID:         imdbID,
		InfoHash:       raw.InfoHash,
		Leechers:       leechers,
		Origin:         raw.Origin,
		ReleaseName:    raw.ReleaseName,
		Resolution:     raw.Resolution,
		Seeders:        seeders,
		Series:         raw.Series,
		Seriesbanner:   raw.Seriesbanner,
		SeriesID:       seriesID,
		SeriesPoster:   raw.SeriesPoster,
		Size:           size,
		Snatched:       snatched,
		Source:         raw.Source,
		Time:           date,
		TorrentID:      torrentID,
		TvdbID:         tvdbID,
		TvrageID:       tvrageID,
		YoutubeTrailer: raw.YoutubeTrailer,
	}
	return nil
}
