package main

import (
	"os"
	"piscine"

	"github.com/01-edu/z01"
)

func printError() {
	z01.PrintRune('E')
	z01.PrintRune('r')
	z01.PrintRune('r')
	z01.PrintRune('o')
	z01.PrintRune('r')
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) != 10 {
		printError()
		return
	}

	board := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		if len(os.Args[i+1]) != 9 {
			printError()
			return
		}
		board[i] = []byte(os.Args[i+1])

		for _, char := range board[i] {
			if !(char == '.' || (char >= '0' && char <= '9')) {
				printError()
				return
			}
		}
	}

	if piscine.SolveSudoku(board) {
		piscine.PrintBoard(board)
	} else {
		printError()
	}
}
