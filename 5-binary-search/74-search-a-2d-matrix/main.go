package main

import (
	"fmt"
)

// https://leetcode.com/problems/search-a-2d-matrix

// You are given an m x n integer matrix matrix with the following two properties:
//
// Each row is sorted in non-decreasing order.
// The first integer of each row is greater than the last integer of the previous row.
// Given an integer target, return true if target is in matrix or false otherwise.
//
// You must write a solution in O(log(m * n)) time complexity.
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 100
// -104 <= matrix[i][j], target <= 10^4

func searchMatrix(matrix [][]int, target int) bool {
	targetRow := -1
	l, r := 0, len(matrix)-1

	for l <= r {
		i := (l + r) / 2
		if matrix[i][0] <= target && matrix[i][len(matrix[i])-1] >= target {
			targetRow = i
			break
		}

		if matrix[i][0] > target {
			r = i - 1
		} else {
			l = i + 1
		}
	}

	if targetRow == -1 {
		return false
	}

	row := matrix[targetRow]
	l, r = 0, len(row)
	for l <= r {
		i := (l + r) / 2
		if row[i] == target {
			return true
		}
		if row[i] < target {
			l = i + 1
		} else {
			r = i - 1
		}
	}
	return false
}

func main() {
	type test struct {
		nums   [][]int
		target int
		out    bool
	}
	tests := []test{
		{nums: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, target: 3, out: true},
		{nums: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, target: 13, out: false},
	}
	for i, t := range tests {
		res := searchMatrix(t.nums, t.target)
		fmt.Printf("%d: %+v %d --> %+v == %+v\n", i, t.nums, t.target, res, t.out)
	}
}
