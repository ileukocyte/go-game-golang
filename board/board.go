package board

import (
	"errors"
)

type Board struct {
	size     int
	board    [][]rune
	xPoints  int
	oPoints  int
	stateSet map[string]struct{}
}

func NewBoard(size int) (*Board, error) {
	if size < 1 {
		return nil, errors.New("non-positive board size")
	}

	board := make([][]rune, size)

	for i := range board {
		board[i] = make([]rune, size)

		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	return &Board{
		size:     size,
		board:    board,
		stateSet: make(map[string]struct{}),
	}, nil
}

func (b *Board) Size() int {
	return b.size
}

func (b *Board) XPoints() int {
	return b.xPoints
}

func (b *Board) OPoints() int {
	return b.oPoints
}
