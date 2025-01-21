package main

import (
	"encoding/json"
	"fmt"
)

// https://leetcode.com/problems/max-area-of-island/description/

// You are given an m x n binary matrix grid. An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.
//
// The area of an island is the number of cells with a value 1 in the island.
//
// Return the maximum area of an island in grid. If there is no island, return 0.
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 50
// grid[i][j] is either 0 or 1.

/*
	- grid dimensions small
	- iterate through grid, when we find 1 explore through dfs, set to 0, add to sum
*/

// iterative O(m*n) / O(m*n)
func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	maxArea := 0
	dirs := [4][2]int{{-1, 0}, {+1, 0}, {0, -1}, {0, +1}}
	stack := [][2]int{}

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				grid[i][j] = 0
				stack = append(stack, [2]int{i, j})

				area := 0

				for len(stack) > 0 {
					pos := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					area++

					for _, d := range dirs {
						ni, nj := pos[0]+d[0], pos[1]+d[1]
						if ni >= 0 && nj >= 0 && ni < m && nj < n && grid[ni][nj] == 1 {
							grid[ni][nj] = 0
							stack = append(stack, [2]int{ni, nj})
						}
					}
				}

				if area > maxArea {
					maxArea = area
				}

				if maxArea == m*n {
					return maxArea
				}
			}
		}
	}

	return maxArea
}

// recursive O(m*n) / O(m*n)
func maxAreaOfIslandR(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	maxArea := 0
	dirs := [4][2]int{{-1, 0}, {+1, 0}, {0, -1}, {0, +1}}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 || i == m || j == n || grid[i][j] == 0 {
			return 0
		}

		areaFromCurr := 1
		grid[i][j] = 0

		for _, d := range dirs {
			areaFromCurr += dfs(i+d[0], j+d[1])
		}

		return areaFromCurr
	}

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				maxArea = max(maxArea, dfs(i, j))
			}
		}
	}

	return maxArea
}

func main() {
	tests := []struct {
		in  string
		out int
	}{
		{
			in:  `[[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]`,
			out: 6,
		},
		{
			in:  `[[0,0,0,0,0,0,0,0]]`,
			out: 0,
		},
	}

	for _, t := range tests {
		var grid [][]int
		json.Unmarshal([]byte(t.in), &grid)

		res := maxAreaOfIsland(grid)
		fmt.Printf("res: %v == out: %v\n", res, t.out)
	}
}
