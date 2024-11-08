package main

import (
	"flag"
	"fmt"
	ggboard "github.com/ileukocyte/go-game-golang/board"
	"os"
)

import (
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Hello, Go Game Web!")

	if err != nil {
		return
	}
}

func mainhttp() {
	http.HandleFunc("/", homeHandler)

	fmt.Println("Server started at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		return
	}
}

func main() {
	sizePtr := flag.Int("size", 9, "Size of a side of the board")

	flag.Parse()

	board, err := ggboard.NewBoard(*sizePtr)

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error:", err)

		return
	}

	fmt.Println(board)

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

			fmt.Println(board)

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

			fmt.Println(board)

			fmt.Printf("Scoreline: %d-%d\n", xPoints, oPoints)

			return
		}

		turn, _ = ggboard.GetOppTurn(turn)
		lastPassed = passed

		fmt.Println(board)
	}
}
