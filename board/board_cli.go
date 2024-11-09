package board

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func (b *Board) String() string {
	var sb strings.Builder

	sizeDigitCount := int(math.Floor(math.Log10(float64(b.size)))) + 1

	// column indices padding
	sb.WriteString(fmt.Sprintf("%*s", sizeDigitCount+1, ""))

	// column indices
	for i := 0; i < b.size; i++ {
		sb.WriteString(fmt.Sprintf("%*d ", sizeDigitCount, i))
	}

	sb.WriteString("\n")

	// rows
	for i, row := range b.board {
		if i > 0 {
			sb.WriteString("\n")
		}

		sb.WriteString(fmt.Sprintf("%*d", sizeDigitCount, i))

		for _, cell := range row {
			sb.WriteString(fmt.Sprintf("%*c", sizeDigitCount+1, cell))
		}
	}

	return sb.String()
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
