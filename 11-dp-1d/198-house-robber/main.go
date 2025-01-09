package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// https://leetcode.com/problems/house-robber/description/

// You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.
//
// Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 400

/*
	- find the maximum sum of the array without using any adjacent elements
	- min jump of 2
	- len(nums) is small
	- nums[i] can't be negative, can never decrease the sum - if we can move we want to

	- given that we can take a min jump of 2, the cumulative sum of either path will be included in -2, or -3
	- decision: which index to go to next from i?
		- plus 2: sum is cumulative sum at i plus value at i+2
		- plus 3: sum is cumulative sum at i plus value at i+3

	- track cumulative sum
*/

// O(n) / O(1)
func rob(nums []int) int {
	one, two, three := 0, 0, 0
	for _, n := range nums {
		one, three, two = n+max(two, three), two, one
	}

	return max(one, two)
}

// O(n) / O(n)
// func rob(nums []int) int {
// 	cumulative := make([]int, len(nums)+3)
//
// 	for i := len(nums) - 1; i >= 0; i-- {
// 		cumulative[i] = nums[i] + max(cumulative[i+2], cumulative[i+3])
// 	}
//
// 	return max(cumulative[0], cumulative[1])
// }

func main() {
	tests := []struct {
		in  string
		out int
	}{
		{
			in:  "[1,2,3,1]",
			out: 4,
		},
		{
			in:  "[2,7,9,3,1]",
			out: 12,
		},
		{
			in:  "[2,20,0,3,10]",
			out: 30,
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
