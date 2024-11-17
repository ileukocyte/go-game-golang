package models

import "time"

/*type AuthType int

const (
	Google AuthType = iota
)*/

type Account struct {
	GoogleAuth

	Username     string    `json:"username"`
	ID           int64     `json:"id"`
	RegisteredAt time.Time `json:"registered_at"`
	LastLogin    time.Time `json:"last_login"`
}

type GoogleAuth struct {
	AccessToken   string `json:"access_token"`
	GoogleID      string `json:"google_id"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}
