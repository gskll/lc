package main

import (
	"encoding/json"
	"fmt"
	"slices"
)

// https://leetcode.com/problems/subsets-ii/description/

// Given an integer array nums that may contain duplicates, return all possible subsets (the power set).
//
// The solution set must not contain duplicate subsets. Return the solution in any order.
//
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10

func subsetsWithDup(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0, 2*n)

	slices.Sort(nums)

	var curr []int
	var backtrack func(currIndex int)
	backtrack = func(currIndex int) {
		fmt.Println(currIndex, curr)
		if currIndex == n {
			res = append(res, append(make([]int, 0, len(curr)), curr...))
			return
		}

		// pick
		curr = append(curr, nums[currIndex])
		backtrack(currIndex + 1)
		curr = curr[:len(curr)-1]

		// skip duplicates
		for currIndex+1 < n && nums[currIndex] == nums[currIndex+1] {
			currIndex++
		}
		backtrack(currIndex + 1)
	}

	backtrack(0)
	return res
}

func main() {
	tests := []struct {
		nums string
		out  string
	}{
		{
			nums: "[1,1,2]",
			out:  "[[],[1],[1,2],[1,1,2],[2],[1,1]]",
		},
		{
			nums: "[1,2,2]",
			out:  "[[],[1],[1,2],[1,2,2],[2],[2,2]]",
		},
		{
			nums: "[0]",
			out:  "[[],[0]]",
		},
	}

	for _, t := range tests {
		var nums []int
		json.Unmarshal([]byte(t.nums), &nums)
		res := subsetsWithDup(nums)
		fmt.Println("got:", res)
		fmt.Println("exp:", t.out)
	}
}
