package db

import (
	"time"

	"github.com/ileukocyte/go-game-golang/db/models"
)

func (env *Env) NewGame(playerX, playerO int64, boardSize int) (int64, error) {
	var id int64

	sqlStatement := `INSERT INTO games (player_x, player_o, size, board, x_turn, last_passed, started_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	emptyBoard := make([][]string, boardSize)

	for i := range emptyBoard {
		emptyBoard[i] = make([]string, boardSize)

		for j := range emptyBoard[i] {
			emptyBoard[i][j] = "."
		}
	}

	err := env.db.QueryRow(
		sqlStatement,
		playerX,
		playerO,
		boardSize,
		emptyBoard,
		true,
		false,
		time.Now(),
	).Scan(&id)

	return id, err
}

func (env *Env) GetGame(id int64) (*models.Game, error) {
	sqlStatement := `SELECT * FROM games WHERE id = $1`

	var g models.Game

	err := env.db.QueryRow(sqlStatement, id).Scan(
		&g.ID,
		&g.PlayerX,
		&g.PlayerO,
		&g.BoardSize,
		&g.Board,
		&g.XTurn,
		&g.LastPassed,
		&g.StartedAt,
		&g.EndedAt,
	)

	return &g, err
}
