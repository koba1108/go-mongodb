package model

import "time"

type Article struct {
	PlayerID         int       `json:"player_id"`
	PlayerName       string    `json:"player_name"`
	VideoID          int       `json:"video_id"`
	VideoName        string    `json:"video_name"`
	VideoDescription string    `json:"video_description"`
	VideoUpdatedDate time.Time `json:"video_updated_date"`
	VideoURL         string    `json:"video_url"`
}
