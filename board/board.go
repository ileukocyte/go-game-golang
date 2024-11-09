package board

import (
	"errors"
	"math/rand"
)

type Board struct {
	size         int
	board        [][]rune
	xPoints      int
	oPoints      int
	zobristTable [][][]uint64
	stateMap     map[uint64]bool
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

	zobristTable := make([][][]uint64, size)

	for i := 0; i < size; i++ {
		zobristTable[i] = make([][]uint64, size)

		for j := 0; j < size; j++ {
			zobristTable[i][j] = make([]uint64, 3)

			for k := 0; k < 3; k++ {
				zobristTable[i][j][k] = rand.Uint64()
			}
		}
	}

	return &Board{
		size:         size,
		board:        board,
		zobristTable: zobristTable,
		stateMap:     make(map[uint64]bool),
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
