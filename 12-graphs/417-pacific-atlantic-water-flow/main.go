package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// https://leetcode.com/problems/pacific-atlantic-water-flow/description/

// There is an m x n rectangular island that borders both the Pacific Ocean and Atlantic Ocean. The Pacific Ocean touches the island's left and top edges, and the Atlantic Ocean touches the island's right and bottom edges.
//
// The island is partitioned into a grid of square cells. You are given an m x n integer matrix heights where heights[r][c] represents the height above sea level of the cell at coordinate (r, c).
//
// The island receives a lot of rain, and the rain water can flow to neighboring cells directly north, south, east, and west if the neighboring cell's height is less than or equal to the current cell's height. Water can flow from any cell adjacent to an ocean into the ocean.
//
// Return a 2D list of grid coordinates result where result[i] = [ri, ci] denotes that rain water can flow from cell (ri, ci) to both the Pacific and Atlantic oceans.
//
// m == heights.length
// n == heights[r].length
// 1 <= m, n <= 200
// 0 <= heights[r][c] <= 10^5

/*
	 - heights are large,
	- grid size medium
	- we want to find all the cells that have a path to both the top-left and bottom-right sides by going vertically or horizontally, where a path is defined by having all cells that are less than or equal the height of the square

	- dfs from each grid?
	- dfs from each cell along top/left? less to try out but how do we know when to switch downhill?
	- is there a bfs solution? could try from the max in each row?

	- some kind of dynamic programming? if we find a valid path and an adjacent cell has a larger value then it's also a valid path

	- we can early return from exploring from a cell as soon as we find a valid path

	- brute force backtracking solution


	- base case will be (i < 0 || j < 0) && (i >= m || j >=m)
*/

var dirs = [4][2]int{{-1, 0}, {+1, 0}, {0, -1}, {0, +1}}

func pacificAtlantic(heights [][]int) [][]int {
}

func main() {
	tests := []struct {
		in  string
		out string
	}{
		{
			in:  `[[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]`,
			out: `[[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]`,
		},
		{
			in:  `[[1]]`,
			out: `[[0,0]]`,
		},
	}

	for _, t := range tests {
		var grid [][]int
		json.Unmarshal([]byte(t.in), &grid)

		res := pacificAtlantic(grid)
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
