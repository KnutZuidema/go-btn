package model

type ForumPage struct {
	Info     ForumInfo     `json:"ForumInfo"`
	MaxPages int           `json:"MaxPages"`
	Threads  []ForumThread `json:"Threads"`
}
