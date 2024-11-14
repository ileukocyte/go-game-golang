package main

import (
	"flag"
	"fmt"
	"github.com/ileukocyte/go-game-golang/internal/game"
	"os"
)

func main() {
	sizePtr := flag.Int("size", 9, "Size of a side of the game")

	flag.Parse()

	b, err := game.NewBoard(*sizePtr)

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error:", err)

		return
	}

	fmt.Println(b)

	turn := game.Cross
	lastPassed := false

	for {
		x, y, passed := game.ReadInput(turn)

		if !passed {
			captured, err := b.OccupyCell(x, y, turn)

			if err != nil {
				fmt.Printf("Invalid turn: %s! Try again!\n", err.Error())

				continue
			}

			fmt.Println(b)

			if captured > 0 {
				fmt.Printf("%c has captured %d stones!\n", turn, captured)
			}

			turn, _ = game.GetOppTurn(turn)
			lastPassed = false

			continue
		}

		if lastPassed {
			xPoints, oPoints := b.XPoints(), b.OPoints()
			xTerritory, oTerritory := b.CountTerritories()

			xPoints += xTerritory
			oPoints += oTerritory

			fmt.Println(b)

			fmt.Printf("Scoreline: %d-%d\n", xPoints, oPoints)

			return
		}

		turn, _ = game.GetOppTurn(turn)
		lastPassed = passed

		fmt.Println(b)
	}
}
