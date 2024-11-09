package board

import "errors"

func (b *Board) OccupyCell(i, j int, turn Turn) (int, error) {
	if i >= b.size || j >= b.size {
		return 0, errors.New("index out of range")
	}

	opp, valid := GetOppTurn(turn)

	if !valid {
		return 0, errors.New("invalid turn")
	}

	cell := &b.board[i][j]

	if *cell != '.' {
		return 0, errors.New("occupied cell")
	}

	*cell = rune(turn)

	var captured int
	var probableSuicide = !b.hasLiberties(i, j)

	copied := b.AsSlice()

	for x, row := range b.board {
		for y, c := range row {
			if c == rune(opp) {
				if !b.hasLiberties(x, y) {
					probableSuicide = false

					copied[x][y] = '.'
					captured++
				}
			}
		}
	}

	if probableSuicide {
		*cell = '.'

		return 0, errors.New("invalid move that results in an unprofitable suicide")
	}

	hash := b.calculateHash()

	if b.stateMap[hash] {
		*cell = '.'

		return 0, errors.New("duplicate state forbidden by the ko rule")
	}

	if turn == Cross {
		b.xPoints += captured
	} else {
		b.oPoints += captured
	}

	b.stateMap[hash] = true
	b.board = copied

	return captured, nil
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
