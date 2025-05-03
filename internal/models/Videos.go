package model

type Video struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Storageplace string `json:"storage"`
}