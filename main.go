package main

import (
	"flag"
	"fmt"
	ggboard "go-game-golang/board"
)

// TODO: implement scoreboard
func main() {
	sizePtr := flag.Int("size", 9, "Size of a side of the board")

	flag.Parse()

	board, err := ggboard.NewBoard(*sizePtr)

	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	board.Display(true)

	turn := ggboard.CROSS
	lastPassed := false

	for {
		x, y, passed := ggboard.ReadInput(turn)

		if !passed {
			err := board.OccupyCell(x, y, turn)

			if err != nil {
				fmt.Printf("Invalid turn: %s! Try again!\n", err.Error())

				continue
			}

			turn = ggboard.GetOppTurn(turn)
			lastPassed = false

			board.Display(true)

			continue
		}

		if lastPassed {
			board.Display(true)

			return
		}

		turn = ggboard.GetOppTurn(turn)
		lastPassed = passed

		board.Display(true)
	}
}
