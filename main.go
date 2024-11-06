package main

import (
	"flag"
	"fmt"
	ggboard "go-game-golang/board"
	"os"
)

func main() {
	sizePtr := flag.Int("size", 9, "Size of a side of the board")
	indicesPtr := flag.Bool("indices", true, "Enable board indices")

	flag.Parse()

	board, err := ggboard.NewBoard(*sizePtr)

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error:", err)

		return
	}

	board.Display(*indicesPtr)

	turn := ggboard.Cross
	lastPassed := false

	for {
		x, y, passed := ggboard.ReadInput(turn)

		if !passed {
			captured, err := board.OccupyCell(x, y, turn)

			if err != nil {
				fmt.Printf("Invalid turn: %s! Try again!\n", err.Error())

				continue
			}

			board.Display(*indicesPtr)

			if captured > 0 {
				fmt.Printf("%c has captured %d stones!\n", turn, captured)
			}

			turn, _ = ggboard.GetOppTurn(turn)
			lastPassed = false

			continue
		}

		if lastPassed {
			xPoints, oPoints := board.XPoints(), board.OPoints()
			xTerritory, oTerritory := board.CountTerritories()

			xPoints += xTerritory
			oPoints += oTerritory

			board.Display(*indicesPtr)

			fmt.Printf("Scoreline: %d-%d\n", xPoints, oPoints)

			return
		}

		turn, _ = ggboard.GetOppTurn(turn)
		lastPassed = passed

		board.Display(*indicesPtr)
	}
}
