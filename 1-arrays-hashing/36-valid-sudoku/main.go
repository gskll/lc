package main

// https://leetcode.com/problems/valid-sudoku/description/

// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
//
// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
// Note:
//
// A Sudoku board (partially filled) could be valid but is not necessarily solvable.
// Only the filled cells need to be validated according to the mentioned rules
//
// board.length == 9
// board[i].length == 9
// board[i][j] is a digit 1-9 or '.'.

func isValidSudoku(board [][]byte) bool {
	var rows, cols, squares [9][9]bool

	for r := range board {
		for c := range board[r] {
			if board[r][c] == '.' {
				continue
			}
			val := board[r][c] - '1'
			squareIndex := (r/3)*3 + c/3
			if rows[r][val] || cols[c][val] || squares[squareIndex][val] {
				return false
			}
			rows[r][val], cols[c][val], squares[squareIndex][val] = true, true, true
		}
	}

	return true
}
