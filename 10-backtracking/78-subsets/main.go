package main

import (
	"encoding/json"
	"fmt"
)

// https://leetcode.com/problems/subsets/description/

// Given an integer array nums of unique elements, return all possible
// subsets
//  (the power set).
//
// The solution set must not contain duplicate subsets. Return the solution in any order.
//
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
// All the numbers of nums are unique.

// no recursion overhead
// O(n * 2^n) / O(n*2^n) to store all subsets
func subsetsBitManipulation(nums []int) [][]int {
	n := len(nums)
	total := 1 << n
	res := make([][]int, 0, total)

	for i := 0; i < total; i++ {
		var subset []int

		for j := 0; j < n; j++ {
			if (i & (1 << j)) != 0 {
				fmt.Println(j)
				subset = append(subset, nums[j])
			}
		}
		res = append(res, subset)
	}
	return res
}

// O(n * 2^n) / O(n*2^n) to store all subsets
func subsets(nums []int) [][]int {
	res := [][]int{}

	var subset []int
	var dfs func(i int)
	dfs = func(i int) {
		if i >= len(nums) {
			// tmp := make([]int, len(subset))
			// copy(tmp, subset)
			// res = append(res, tmp)
			res = append(res, append(make([]int, 0, len(subset)), subset...))
			return
		}

		// pick
		subset = append(subset, nums[i])
		dfs(i + 1)

		// skip
		subset = subset[:len(subset)-1]
		dfs(i + 1)
	}

	dfs(0)
	return res
}

// func subsets(nums []int) [][]int {
// 	res := [][]int{{}}
//
// 	var dfs func(curr []int, i int)
// 	dfs = func(curr []int, i int) {
// 		res = append(res, append(curr, nums[i]))
//
// 		if i < len(nums)-1 {
// 			dfs(curr, i+1)
// 			dfs(append(curr, nums[i]), i+1)
// 		}
// 	}
//
// 	dfs([]int{}, 0)
// 	return res
// }

func main() {
	tests := []struct {
		nums string
		out  string
	}{
		{
			nums: "[1,2,3]",
			out:  "[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]",
		},
		{
			nums: "[0]",
			out:  "[[],[0]]",
		},
	}

	for _, t := range tests {
		var nums []int
		json.Unmarshal([]byte(t.nums), &nums)
		res := subsetsBitManipulation(nums)
		fmt.Println("got:", res)
		fmt.Println("exp:", t.out)
	}
}
