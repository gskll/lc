package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// https://leetcode.com/problems/house-robber-ii/description/

// You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. All houses at this place are arranged in a circle. That means the first house is the neighbor of the last one. Meanwhile, adjacent houses have a security system connected, and it will automatically contact the police if two adjacent houses were broken into on the same night.
//
// Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 1000

/*
	- find the max sum of non-adjacent elements in a circular array
	- the first and last elements are considered adjacent
	- no negative elements, sum can't decrease by visiting an element

	- if we start from the first element we can include up until the second-to last
	- if we start from the second element we can include up until the last

	- we run the solution to rob houses i twice on both slices
	- what about edge cases?
		if len(nums) <= 3?
*/

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var helper func(nums []int) int
	helper = func(nums []int) int {
		a, b, c := 0, 0, 0
		for i := 0; i < len(nums); i++ {
			a, b, c = nums[i]+max(b, c), a, b
		}
		return max(a, b)
	}

	return max(helper(nums[:len(nums)-1]), helper(nums[1:]))
}

func main() {
	tests := []struct {
		in  string
		out int
	}{
		{
			in:  "[2,3,2]",
			out: 3,
		},
		{
			in:  "[1,2,3,1]",
			out: 4,
		},
		{
			in:  "[1,2,3]",
			out: 3,
		},
		{
			in:  "[1]",
			out: 1,
		},
		{
			in:  "[1,2]",
			out: 2,
		},
	}

	for _, t := range tests {
		var in []int
		json.Unmarshal([]byte(t.in), &in)
		start := time.Now()
		res := rob(in)
		fmt.Printf("%v == %v | %v\n", t.out, res, time.Since(start))
	}
}
