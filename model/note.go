package model

type Note struct {
	ID     int    `json:"id"`
	Task   string `json:"task"`
	Status bool   `json:"status"`
}
