package models

import "time"

type Account struct {
	Auth

	Username     string
	ID           uint64
	RegisteredAt time.Time
	LastLogin    time.Time
}

type Auth struct {
	GoogleID      string
	AccessToken   string
	Email         string
	EmailVerified bool
}
