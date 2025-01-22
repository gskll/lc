package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// https://neetcode.io/problems/islands-and-treasure

// You are given a m×n 2D grid initialized with these three possible values:
//
// -1 - A water cell that can not be traversed.
// 0 - A treasure chest.
// INF - A land cell that can be traversed. We use the integer 2^31 - 1 = 2147483647 to represent INF.
//
// Fill each land cell with the distance to its nearest treasure chest. If a land cell cannot reach a treasure chest than the value should remain INF.
//
// Assume the grid can only be traversed up, down, left, or right.
//
// Modify the grid in-place
//
// Constraints:
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 100
// grid[i][j] is one of {-1, 0, 2147483647}

/*
	- m,n medium
	- inf is maxint32

	- for each land cell represented by inf, we want to calculate the distance to the nearest treasure chest and replace the cell value with the distance
	- we want to minimise the distance to the nearest treasure chest

	- approaches?
		- 1. can we parse all the treasure chests and perform math on each land cell ? some kind of ordering for the chests based on pos
		- 2. start at each cell and dfs to all treasure chests and take min - brute force - O(m*n*4^mn)
		- 3. for each treasure chest we dfs all cells nearby and replace the value if it's smaller

	- find treasure cell
	- go in all 4 directions while there's connecting land and the distance from the current cell is smaller than the distance from the previously calculated distance

	- bfs solution instead of dfs
	- find all treasure chests and add to the queue
	- we can then expand from each chest at the same time, marking
*/

const land = (1 << 31) - 1

var dirs = [4][2]int{{-1, 0}, {+1, 0}, {0, -1}, {0, +1}}

// multi-source bfs from each chest O(m*n) / O(m*n)
func islandsAndTreasure(grid [][]int) {
	m, n := len(grid), len(grid[0])
	queue := make([][2]int, 0, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		i, j := cur[0], cur[1]

		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			if ni < 0 || nj < 0 || ni >= m || nj >= n || grid[ni][nj] != land {
				continue
			}

			grid[ni][nj] = grid[i][j] + 1
			queue = append(queue, [2]int{ni, nj})
		}

	}
}

// dfs O(m*n*m*n) / O(m*n)
// for every chest we visit potentially cell, and every cell is potentially a chest. m*n cells
func islandsAndTreasureDFS(grid [][]int) {
	m, n := len(grid), len(grid[0])

	var calculateDistances func(i, j, dist int)
	calculateDistances = func(i, j, dist int) {
		if i < 0 || j < 0 || i >= m || j >= n { // out of bounds
			return
		}

		if grid[i][j] == -1 { // water cell
			return
		}

		if dist != 0 && grid[i][j] <= dist { // found smaller already
			return
		}

		grid[i][j] = dist

		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			calculateDistances(ni, nj, dist+1)
		}
	}

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 0 {
				calculateDistances(i, j, 0)
			}
		}
	}
}

func main() {
	tests := []struct {
		in  string
		out string
	}{
		{
			in: `[
  [2147483647,-1,0,2147483647],
  [2147483647,2147483647,2147483647,-1],
  [2147483647,-1,2147483647,-1],
  [0,-1,2147483647,2147483647]
]`,
			out: `[
  [3,-1,0,1],
  [2,2,1,-1],
  [1,-1,2,-1],
  [0,-1,3,4]
]`,
		},
		{
			in: `[
  [0,-1],
  [2147483647,2147483647]
]`,
			out: `[
  [0,-1],
  [1,2]
]`,
		},
	}

	for _, t := range tests {
		var grid [][]int
		json.Unmarshal([]byte(t.in), &grid)

		islandsAndTreasure(grid)
		fmt.Printf("out: %v\nres: %v\n", t.out, printGrid(grid))
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
