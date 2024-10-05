package piscine

import (
	"github.com/01-edu/z01"
)

func IsValid(board [][]byte, row, col int, char byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == char || board[i][col] == char || board[3*(row/3)+i/3][3*(col/3)+i%3] == char {
			return false
		}
	}
	return true
}

func SolveSudoku(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				for char := byte('1'); char <= '9'; char++ {
					if IsValid(board, row, col, char) {
						board[row][col] = char
						if SolveSudoku(board) {
							return true
						}
						board[row][col] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func PrintBoard(board [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			z01.PrintRune(rune(board[i][j]))
			if j < 8 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
