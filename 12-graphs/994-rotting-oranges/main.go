package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// https://leetcode.com/problems/rotting-oranges/description/

// You are given an m x n grid where each cell can have one of three values:
//
// 0 representing an empty cell,
// 1 representing a fresh orange, or
// 2 representing a rotten orange.
// Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.
//
// Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 10
// grid[i][j] is 0, 1, or 2.

/*
	- grid size is small
	- count rotting/fresh oranges
	- multi-source bfs from each rotting orange then make sure none left - count longest path
	- or maybe dfs from rotting oranges to longest path?
*/

var dirs = [4][2]int{{-1, 0}, {+1, 0}, {0, -1}, {0, +1}}

const (
	empty  = 0
	fresh  = 1
	rotten = 2
)

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	queue := make([][2]int, 0, m*n)
	numFresh := 0

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			switch grid[r][c] {
			case rotten:
				queue = append(queue, [2]int{r, c})
			case fresh:
				numFresh++
			}
		}
	}

	if numFresh == 0 {
		return 0
	}

	minutes := 0
	for numFresh > 0 && len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			queue = queue[1:]
			row, col := cur[0], cur[1]

			for _, dir := range dirs {
				r, c := row+dir[0], col+dir[1]
				if r < 0 || c < 0 || r >= m || c >= n || grid[r][c] != fresh {
					continue
				}

				grid[r][c] = rotten
				numFresh--
				queue = append(queue, [2]int{r, c})
			}
		}
		minutes++
	}

	if numFresh == 0 {
		return minutes
	}

	return -1
}

func main() {
	tests := []struct {
		in  string
		out int
	}{
		{
			in:  `[[2,1,1],[1,1,0],[0,1,1]]`,
			out: 4,
		},
		{
			in:  `[[2,1,1],[0,1,1],[1,0,1]]`,
			out: -1,
		},
		{
			in:  `[[0,2]]`,
			out: 0,
		},
		{
			in:  `[[0]]`,
			out: 0,
		},
		{
			in:  `[[1,2]]`,
			out: 1,
		},
	}

	for _, t := range tests {
		var grid [][]int
		json.Unmarshal([]byte(t.in), &grid)

		res := orangesRotting(grid)
		fmt.Printf("out: %v == res: %v\n", t.out, res)
	}
}

func printGrid(grid [][]int) string {
	var s strings.Builder
	s.WriteString("[\n")

	for i, row := range grid {
		rowBytes, _ := json.Marshal(row)
		s.WriteString("  ")
		s.Write(rowBytes)
		if i < len(grid)-1 {
			s.WriteString(",")
		}
		s.WriteString("\n")
	}

	s.WriteString("]")
	return s.String()
}
