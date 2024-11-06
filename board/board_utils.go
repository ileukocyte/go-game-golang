package board

import "strings"

type Turn rune

const (
	CROSS  Turn = 'X'
	NOUGHT Turn = 'O'
)

func GetOppTurn(cur Turn) Turn {
	if cur == NOUGHT {
		return CROSS
	}

	return NOUGHT
}

func (b *Board) AsStateStr() string {
	builder := strings.Builder{}

	for _, row := range b.board {
		for _, cell := range row {
			builder.WriteString(string(cell))
		}
	}

	return builder.String()
}

func (b *Board) AsSlice() [][]rune {
	copied := make([][]rune, b.size)

	for i := range b.board {
		copied[i] = make([]rune, b.size)

		copy(copied[i], b.board[i])
	}

	return copied
}
