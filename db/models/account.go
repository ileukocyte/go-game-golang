package models

import "time"

/*type AuthType int

const (
	Google AuthType = iota
)*/

type Account struct {
	GoogleAuth

	Username     string
	ID           uint64
	RegisteredAt time.Time
	LastLogin    time.Time
}

type GoogleAuth struct {
	AccessToken   string
	GoogleID      string
	Email         string
	EmailVerified bool
}
