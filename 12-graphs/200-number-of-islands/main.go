package main

import (
	"encoding/json"
	"fmt"
)

// https://leetcode.com/problems/number-of-islands/description/

// Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.
//
// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.
//
// Constraints:
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] is '0' or '1'.

/*
   - grid is medium
   - when we reach a 1 we want to explore and find island edges
   - then iterate to next unseen 1 and explore from there
   - we can iterate over the grid. if a cell is a 1 and up/left are water it's a new island -- doesn't work in 3rd test case

    - explore each island and mark the visited cells as .
    -
*/

func numIslands(grid [][]byte) int {
	count := 0
	m, n := len(grid), len(grid[0])

	dirs := [4][2]int{{-1, 0}, {0, -1}, {+1, 0}, {0, +1}}

	var exploreIsland func(i, j int)
	exploreIsland = func(i, j int) {
		if i < 0 || i == m || j < 0 || j == n || grid[i][j] != '1' {
			return
		}

		grid[i][j] = '0'

		for _, d := range dirs {
			exploreIsland(i+d[0], j+d[1])
		}
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				count++
				exploreIsland(i, j)
			}
		}
	}

	return count
}

func main() {
	tests := []struct {
		in  string
		out int
	}{
		{
			in: `[
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]`,
			out: 1,
		},
		{
			in: `[
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]`,
			out: 3,
		},
		{
			in: `[["1","1","1"],
                  ["0","1","0"],
                  ["1","1","1"]]`,
			out: 1,
		},
	}

	for _, t := range tests {
		var rawGrid [][]string
		json.Unmarshal([]byte(t.in), &rawGrid)

		grid := make([][]byte, len(rawGrid))
		for i := range rawGrid {
			grid[i] = make([]byte, len(rawGrid[i]))
			for j := range rawGrid[i] {
				grid[i][j] = rawGrid[i][j][0]
			}
		}

		res := numIslands(grid)
		fmt.Printf("res: %v == out: %v\n", res, t.out)
	}
}
