package board

type pair struct {
	First  int
	Second int
}

func (b *Board) CountTerritories() (int, int) {
	territoryOwner := func(region *[]pair) (Turn, bool) {
		owner, isConsistent := rune(0), true

		for _, coords := range *region {
			i, j := coords.First, coords.Second

			if i > 0 && isConsistent {
				if b.board[i-1][j] != '.' {
					if owner == 0 {
						owner = b.board[i-1][j]
					} else {
						isConsistent = owner == b.board[i-1][j]
					}
				}
			}

			if j > 0 && isConsistent {
				if b.board[i][j-1] != '.' {
					if owner == 0 {
						owner = b.board[i][j-1]
					} else {
						isConsistent = owner == b.board[i][j-1]
					}
				}
			}

			if i < b.size-1 && isConsistent {
				if b.board[i+1][j] != '.' {
					if owner == 0 {
						owner = b.board[i+1][j]
					} else {
						isConsistent = owner == b.board[i+1][j]
					}
				}
			}

			if j < b.size-1 && isConsistent {
				if b.board[i][j+1] != '.' {
					if owner == 0 {
						owner = b.board[i][j+1]
					} else {
						isConsistent = owner == b.board[i][j+1]
					}
				}
			}
		}

		return Turn(owner), isConsistent
	}

	var xTerritory, oTerritory int

	visited := make([][]bool, b.size)

	for i := range visited {
		visited[i] = make([]bool, b.size)

		for j := range visited[i] {
			visited[i][j] = false
		}
	}

	for i, row := range b.board {
		for j, cell := range row {
			if cell == '.' && !visited[i][j] {
				var region []pair

				b.fillBlankRegion(i, j, &visited, &region)

				owner, isConsistent := territoryOwner(&region)

				if isConsistent {
					switch owner {
					case Cross:
						xTerritory += len(region)
					case Nought:
						oTerritory += len(region)
					}
				}
			}
		}
	}

	return xTerritory, oTerritory
}

func (b *Board) fillBlankRegion(i, j int, visited *[][]bool, region *[]pair) {
	if i < 0 || j < 0 || i >= b.size || j >= b.size {
		return
	}

	if (*visited)[i][j] || b.board[i][j] != '.' {
		return
	}

	*region = append(*region, pair{i, j})
	(*visited)[i][j] = true

	b.fillBlankRegion(i-1, j, visited, region)
	b.fillBlankRegion(i+1, j, visited, region)
	b.fillBlankRegion(i, j-1, visited, region)
	b.fillBlankRegion(i, j+1, visited, region)
}
