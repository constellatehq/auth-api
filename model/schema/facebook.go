package schema

import "time"

type FacebookTokenValidationResponse struct {
	AppId               string   `json:"app_id"`
	Application         string   `json:"application"`
	DataAccessExpiresAt int64    `json:"data_access_expires_at"`
	ExpiresAt           int64    `json:"expires_at"`
	IsValid             bool     `json:"is_valid"`
	IssusedAt           int64    `json:"issued_at"`
	Scopes              []string `json:"scopes"`
	Type                string   `json:"type"`
	UserId              string   `json:"user_id"`
}

type FacebookUserInfoResponse struct {
	Id        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Birthday  time.Time `json:"birthday"`
	Gender    string    `json:"gender"`
}
