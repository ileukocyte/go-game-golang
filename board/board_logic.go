package board

import "errors"

func (b *Board) OccupyCell(i, j int, turn Turn) error {
	if i >= b.size || j >= b.size {
		return errors.New("index out of range")
	}

	cell := &b.board[i][j]

	if *cell != '.' {
		return errors.New("occupied cell")
	}

	*cell = rune(turn)

	var probableSuicide = !b.hasLiberties(i, j)
	opp := GetOppTurn(turn)

	copied := b.AsSlice()

	for x, row := range b.board {
		for y, c := range row {
			if c == rune(opp) {
				if !b.hasLiberties(x, y) {
					probableSuicide = false

					copied[x][y] = '.'

					if turn == CROSS {
						b.xPoints++
					} else {
						b.oPoints++
					}
				}
			}
		}
	}

	if probableSuicide {
		*cell = '.'

		return errors.New("invalid move that results in an unprofitable suicide")
	}

	stateStr := b.AsStateStr()

	if _, contains := b.stateSet[stateStr]; contains {
		*cell = '.'

		return errors.New("duplicate state forbidden by the ko rule")
	}

	b.stateSet[stateStr] = struct{}{}
	b.board = copied

	return nil
}

func (b *Board) hasLiberties(i, j int) bool {
	cell := b.board[i][j]

	visited := make([][]bool, b.size)

	for i := range visited {
		visited[i] = make([]bool, b.size)

		for j := range visited[i] {
			visited[i][j] = false
		}
	}

	return b.libertyCheck(i, j, cell, &visited)
}

func (b *Board) libertyCheck(i, j int, sign rune, visited *[][]bool) bool {
	if i >= b.size || j >= b.size || (*visited)[i][j] {
		return false
	}

	if b.board[i][j] == '.' {
		return true
	}

	if b.board[i][j] != sign {
		return false
	}

	(*visited)[i][j] = true

	hasLiberty := (i > 0 && b.libertyCheck(i-1, j, sign, visited)) ||
		(j > 0 && b.libertyCheck(i, j-1, sign, visited)) ||
		(i < b.size-1 && b.libertyCheck(i+1, j, sign, visited)) ||
		(j < b.size-1 && b.libertyCheck(i, j+1, sign, visited))

	return hasLiberty
}
