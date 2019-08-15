package model

import (
	"encoding/json"
	"strconv"
	"time"
)

type User struct {
	UserID          int       `json:"UserID"`
	Username        string    `json:"Username"`
	Email           string    `json:"Email"`
	Upload          int       `json:"Upload"`
	Download        int       `json:"Download"`
	Lumens          int       `json:"Lumens"`
	Bonus           float64   `json:"Bonus"`
	JoinDate        time.Time `json:"JoinDate"`
	Title           string    `json:"Title"`
	Enabled         bool      `json:"Enabled"`
	Paranoia        int       `json:"Paranoia"`
	Invites         int       `json:"Invites"`
	Class           string    `json:"Class"`
	ClassLevel      int       `json:"ClassLevel"`
	HnR             int       `json:"HnR"`
	UploadsSnatched int       `json:"UploadsSnatched"`
	Snatches        int       `json:"Snatches"`
}

type rawUser struct {
	UserID          string `json:"UserID"`
	Username        string `json:"Username"`
	Email           string `json:"Email"`
	Upload          string `json:"Upload"`
	Download        string `json:"Download"`
	Lumens          string `json:"Lumens"`
	Bonus           string `json:"Bonus"`
	JoinDate        string `json:"JoinDate"`
	Title           string `json:"Title"`
	Enabled         string `json:"Enabled"`
	Paranoia        string `json:"Paranoia"`
	Invites         string `json:"Invites"`
	Class           string `json:"Class"`
	ClassLevel      string `json:"ClassLevel"`
	HnR             string `json:"HnR"`
	UploadsSnatched string `json:"UploadsSnatched"`
	Snatches        string `json:"Snatches"`
}

func (u *User) UnmarshalJSON(data []byte) error {
	var raw rawUser
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	id, err := strconv.Atoi(raw.UserID)
	if err != nil {
		return err
	}
	upload, err := strconv.Atoi(raw.Upload)
	if err != nil {
		return err
	}
	download, err := strconv.Atoi(raw.Download)
	if err != nil {
		return err
	}
	lumens, err := strconv.Atoi(raw.Lumens)
	if err != nil {
		return err
	}
	bonus, err := strconv.ParseFloat(raw.Bonus, 64)
	if err != nil {
		return err
	}
	ts, err := strconv.Atoi(raw.JoinDate)
	if err != nil {
		return err
	}
	join := time.Unix(int64(ts), 0)
	enabled := raw.Enabled != "0"
	paranoia, err := strconv.Atoi(raw.Paranoia)
	if err != nil {
		return err
	}
	invites, err := strconv.Atoi(raw.Invites)
	if err != nil {
		return err
	}
	level, err := strconv.Atoi(raw.ClassLevel)
	if err != nil {
		return err
	}
	hnr, err := strconv.Atoi(raw.HnR)
	if err != nil {
		return err
	}
	upsnatched, err := strconv.Atoi(raw.UploadsSnatched)
	if err != nil {
		return err
	}
	snatches, err := strconv.Atoi(raw.Snatches)
	if err != nil {
		return err
	}
	*u = User{
		UserID:          id,
		Username:        raw.Username,
		Email:           raw.Email,
		Upload:          upload,
		Download:        download,
		Lumens:          lumens,
		Bonus:           bonus,
		JoinDate:        join,
		Title:           raw.Title,
		Enabled:         enabled,
		Paranoia:        paranoia,
		Invites:         invites,
		Class:           raw.Class,
		ClassLevel:      level,
		HnR:             hnr,
		UploadsSnatched: upsnatched,
		Snatches:        snatches,
	}
	return nil
}
