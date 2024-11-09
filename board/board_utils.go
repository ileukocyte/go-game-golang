package board

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

func (b *Board) AsSlice() [][]rune {
	copied := make([][]rune, b.size)

	for i := range b.board {
		copied[i] = make([]rune, b.size)

		copy(copied[i], b.board[i])
	}

	return copied
}

func (b *Board) calculateHash() uint64 {
	var hash uint64

	for i, row := range b.board {
		for j, cell := range row {
			var piece int

			switch cell {
			case '.':
				piece = 0
			case rune(Cross):
				piece = 1
			case rune(Nought):
				piece = 2
			}

			hash ^= b.zobristTable[i][j][piece]
		}
	}

	return hash
}
