package models

import (
	"time"

	"golang.org/x/oauth2"
)

type Account struct {
	Username     string    `json:"username"`
	ID           int64     `json:"id"`
	RegisteredAt time.Time `json:"registered_at"`
	LastLogin    time.Time `json:"last_login"`

	Token oauth2.Token
}
