package board

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func (b *Board) Display(enableIndices bool) {
	digitCount := func(i int) int {
		if i == 0 {
			return 1
		}

		return int(math.Floor(math.Log10(float64(i)))) + 1
	}

	if enableIndices {
		padding := strings.Repeat(" ", digitCount(b.size)+1)

		fmt.Print(padding)

		for i := 0; i < b.size; i++ {
			fmt.Printf("%*d ", digitCount(b.size), i)
		}

		fmt.Println()
	}

	for y, row := range b.board {
		for x, cell := range row {
			if x == 0 && enableIndices {
				fmt.Printf("%*d", digitCount(b.size), y)
			}

			padding := ""

			if !enableIndices {
				if x != 0 {
					padding = strings.Repeat(" ", 1)
				}
			} else {
				padding = strings.Repeat(" ", digitCount(b.size))
			}

			fmt.Printf("%s%c", padding, cell)
		}

		fmt.Println()
	}
}

// ReadInput returns the coordinates read from the standard input and whether the actual input is "pass"
func ReadInput(turn Turn) (int, int, bool) {
	var x, y int

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Current turn: %c\n", turn)
		fmt.Print("Enter two numbers (row column) or 'pass': ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "pass" {
			return x, y, true
		}

		fields := strings.Fields(input)

		if len(fields) != 2 {
			fmt.Println("Invalid input! Try again!")

			continue
		}

		x, errX := strconv.Atoi(fields[0])
		y, errY := strconv.Atoi(fields[1])

		if errX != nil || errY != nil {
			fmt.Println("Invalid input! Try again!")

			continue
		}

		return x, y, false
	}
}
