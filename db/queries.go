package db

import (
	"time"

	"github.com/ileukocyte/go-game-golang/db/models"
)

func (e *Env) NewGame(playerX, playerO int64, boardSize int) (int64, error) {
	sqlStatement := `INSERT INTO games (player_x, player_o, size, board, x_turn, last_passed, started_at)
		VALUES ($1, $2, $3, $4, $5)`

	emptyBoard := make([][]string, boardSize)

	for i := range emptyBoard {
		emptyBoard[i] = make([]string, boardSize)

		for j := range emptyBoard[i] {
			emptyBoard[i][j] = "."
		}
	}

	res, err := e.db.Exec(sqlStatement, playerX, playerO, boardSize, emptyBoard, true, false, time.Now())

	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()

	return id, err
}

func (e *Env) GetGame(id int64) (*models.Game, error) {
	sqlStatement := `SELECT * FROM games WHERE id = $1`

	var g models.Game

	err := e.db.QueryRow(sqlStatement, id).Scan(
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
