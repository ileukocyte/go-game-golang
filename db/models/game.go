package models

import (
	"database/sql"
	"time"
)

type Game struct {
	ID         int64               `json:"id"`
	PlayerX    int64               `json:"player_x"`
	PlayerO    int64               `json:"player_o"`
	BoardSize  int                 `json:"board_size"`
	Board      [][]string          `json:"board"`
	XTurn      bool                `json:"x_turn"`
	LastPassed bool                `json:"last_passed"`
	StartedAt  time.Time           `json:"started_at"`
	EndedAt    sql.Null[time.Time] `json:"ended_at"`
}
