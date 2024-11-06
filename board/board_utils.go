package board

import "strings"

type Turn rune

const (
	Cross  Turn = 'X'
	Nought Turn = 'O'
)

func GetOppTurn(cur Turn) (Turn, bool) {
	switch cur {
	case Cross:
		return Nought, true
	case Nought:
		return Cross, true
	default:
		return 0, false
	}
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
