package model

import (
	"time"
)

type User struct {
	Id              string    `json:"id"`
	FacebookId      string    `json:"facebook_id"`
	GoogleId        string    `json:"google_id"`
	InstagramId     string    `json:"instagram_id"`
	SpotifyId       string    `json:"spotify_id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Email           string    `json:"email"`
	PermissionLevel int64     `json:"permission_level"`
	CreatedAt       time.Time `json:"created_at"`
}
