package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// https://leetcode.com/problems/word-search/description/

// Given an m x n grid of characters board and a string word, return true if word exists in the grid.
//
// The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.
//
// m == board.length
// n = board[i].length
// 1 <= m, n <= 6
// 1 <= word.length <= 15
// board and word consists of only lowercase and uppercase English letters.

// O(m*n*l^4) m =rows, n=cols, l=word length
// O(1)
func exist(board [][]byte, word string) bool {
	// check if all characters in word exist in board
	var count [128]int
	for _, b := range []byte(word) {
		count[b-'A']++
	}

	rows := len(board)
	cols := len(board[0])
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			count[board[r][c]-'A']--
		}
	}

	for _, c := range count {
		if c > 0 {
			return false
		}
	}

	n := len(word)

	var dfs func(row, col, i int) bool
	dfs = func(row, col, i int) bool {
		if i == n {
			return true
		}

		if row < 0 || col < 0 || row >= rows || col >= cols {
			return false
		}

		char := board[row][col]
		if char == '#' || char != word[i] {
			return false
		}

		board[row][col] = '#'

		res := dfs(row+1, col, i+1) || dfs(row-1, col, i+1) || dfs(row, col+1, i+1) || dfs(row, col-1, i+1)

		board[row][col] = char
		return res
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == word[0] {
				if res := dfs(r, c, 0); res {
					return res
				}
			}
		}
	}

	return false
}

func main() {
	tests := []struct {
		board string
		word  string
		out   bool
	}{
		{
			board: `[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]`,
			word:  "SEE",
			out:   true,
		},
		{
			board: `[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]`,
			word:  "ABCB",
			out:   false,
		},
	}

	for _, test := range tests {
		var board [][]byte
		var boardRaw [][]string
		json.Unmarshal([]byte(test.board), &boardRaw)
		board = make([][]byte, len(boardRaw))
		for i := range boardRaw {
			board[i] = make([]byte, len(boardRaw[i]))
			for j := range boardRaw[i] {
				board[i][j] = byte(boardRaw[i][j][0])
			}
		}
		fmt.Println()
		start := time.Now()
		res := exist(board, test.word)
		fmt.Printf("time: %s\n", time.Since(start))
		fmt.Printf("got: %+v\n", res)
		fmt.Printf("exp: %+v\n", test.out)
	}
}
