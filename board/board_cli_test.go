package board

import "testing"

func BenchmarkBoard_String(b *testing.B) {
	board, _ := NewBoard(9)

	for i := 0; i < b.N; i++ {
		_ = board.String()
	}
}

func BenchmarkBoard_StringAlt(b *testing.B) {
	board, _ := NewBoard(9)

	for i := 0; i < b.N; i++ {
		_ = board.StringAlt()
	}
}
