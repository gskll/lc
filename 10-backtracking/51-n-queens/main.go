package main

import (
	"fmt"
	"strings"
	"time"
)

// https://leetcode.com/problems/n-queens/description/

// The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.
//
// Given an integer n, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.
//
// Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space, respectively.
//
// 1 <= n <= 9

// O(n!) / O(n)
func solveNQueens(n int) [][]string {
	var res [][]string

	// queens[row] stores the column position of the queen in that row
	queens := make([]int, n)
	cols := make([]bool, n)

	// Diagonal attack tracking:
	// For a board position (r,c):
	// - Positive diagonal (↗): All positions where r-c is constant
	//   Index = (r-c)+(n-1) to handle negative differences
	// - Negative diagonal (↘): All positions where r+c is constant
	//   Index = r+c directly maps each diagonal
	//
	// Example for n=4:
	// pos-diag indices:  neg-diag indices:
	// 3 2 1 0            0 1 2 3
	// 4 3 2 1            1 2 3 4
	// 5 4 3 2            2 3 4 5
	// 6 5 4 3            3 4 5 6
	posDiags := make([]bool, 2*n-1) // ↗ diagonals (r-c+n-1)
	negDiags := make([]bool, 2*n-1) // ↘ diagonals (r+c)

	// Efficient board creation using pre-allocated builder
	var createBoard func() []string
	createBoard = func() []string {
		board := make([]string, n)

		for row := 0; row < n; row++ {
			var rowStr strings.Builder
			rowStr.Grow(n)

			for col := 0; col < n; col++ {
				if queens[row] == col {
					rowStr.WriteByte('Q')
				} else {
					rowStr.WriteByte('.')
				}
			}
			board[row] = rowStr.String()
		}

		return board
	}

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			res = append(res, createBoard())
			return
		}

		for col := 0; col < n; col++ {
			posDiag := row - col + n - 1 // ↗ diagonal
			negDiag := row + col         // ↘ diagonal

			if cols[col] || posDiags[posDiag] || negDiags[negDiag] {
				continue
			}

			// Place queen and mark attack lines
			queens[row] = col // will be overwritten in backtracking
			cols[col] = true
			posDiags[posDiag] = true
			negDiags[negDiag] = true

			backtrack(row + 1)

			// Backtrack: remove queen and unmark attack lines
			cols[col] = false
			posDiags[posDiag] = false
			negDiags[negDiag] = false
		}
	}

	backtrack(0)

	return res
}

// solution 1
// type board [][]int
//
// func newBoard(n int) board {
// 	b := make([][]int, 0, n)
// 	for i := 0; i < n; i++ {
// 		b = append(b, make([]int, n))
// 	}
// 	return b
// }
//
// func (b board) place(row, col int) {
// 	n := len(b)
// 	// mark queen
// 	b[row][col] = -1
//
// 	// mark horizontal/vertical
// 	for i := 0; i < n; i++ {
// 		if i != row {
// 			b[i][col]++
// 		}
//
// 		if i != col {
// 			b[row][i]++
// 		}
// 	}
//
// 	// mark diagonals
// 	for i := 1; row-i >= 0 && col-i >= 0; i++ {
// 		b[row-i][col-i]++ // up-left
// 	}
// 	for i := 1; row-i >= 0 && col+i < n; i++ {
// 		b[row-i][col+i]++ // up-right
// 	}
// 	for i := 1; row+i < n && col-i >= 0; i++ {
// 		b[row+i][col-i]++ // down-left
// 	}
// 	for i := 1; row+i < n && col+i < n; i++ {
// 		b[row+i][col+i]++ // down-right
// 	}
// }
//
// func (b board) remove(row, col int) {
// 	n := len(b)
//
// 	// mark queen
// 	b[row][col] = 0
//
// 	// mark horizontal/vertical
// 	for i := 0; i < n; i++ {
// 		if i != row {
// 			b[i][col]--
// 		}
//
// 		if i != col {
// 			b[row][i]--
// 		}
// 	}
//
// 	// mark diagonals
// 	for i := 1; row-i >= 0 && col-i >= 0; i++ {
// 		b[row-i][col-i]-- // up-left
// 	}
// 	for i := 1; row-i >= 0 && col+i < n; i++ {
// 		b[row-i][col+i]-- // up-right
// 	}
// 	for i := 1; row+i < n && col-i >= 0; i++ {
// 		b[row+i][col-i]-- // down-left
// 	}
// 	for i := 1; row+i < n && col+i < n; i++ {
// 		b[row+i][col+i]-- // down-right
// 	}
// }
//
// func solveNQueens(n int) [][]string {
// 	if n == 1 {
// 		return [][]string{{"Q"}}
// 	}
// 	if n == 2 || n == 3 {
// 		return nil
// 	}
//
// 	board := newBoard(n)
// 	var res [][]string
// 	var backtrack func(row int)
// 	backtrack = func(row int) {
// 		if row == n {
// 			validBoard := make([]string, 0, n)
// 			for r := 0; r < n; r++ {
// 				validRow := make([]string, 0, n)
// 				for c := 0; c < n; c++ {
// 					if board[r][c] == -1 {
// 						validRow = append(validRow, "Q")
// 					} else {
// 						validRow = append(validRow, ".")
// 					}
// 				}
// 				validBoard = append(validBoard, strings.Join(validRow, ""))
// 			}
// 			res = append(res, validBoard)
// 			return
// 		}
//
// 		for col := 0; col < n; col++ {
// 			if board[row][col] != 0 {
// 				continue
// 			}
// 			board.place(row, col)
// 			backtrack(row + 1)
// 			board.remove(row, col)
// 		}
// 	}
//
// 	backtrack(0)
//
// 	return res
// }

func main() {
	tests := []struct {
		n   int
		out string
	}{
		{
			n:   4,
			out: `[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]`,
		},
		{
			n:   1,
			out: `[["Q"]]`,
		},
	}

	for _, test := range tests {
		fmt.Println()
		start := time.Now()
		res := solveNQueens(test.n)
		fmt.Printf("time: %s\n", time.Since(start))
		fmt.Printf("got: %+v\n", res)
		fmt.Printf("exp: %+v\n", test.out)
	}
}
