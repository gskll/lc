package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// https://leetcode.com/problems/maximum-subarray/description/

// Given an integer array nums, find the subarray with the largest sum, and return its sum.
//
// 1 <= nums.length <= 10^5
// -10^4 <= nums[i] <= 10^4

/*
	- sum will always increase with positive numbers
	- if the subarray sum goes negative, we want to reset

	- F(i) = max(nums[i], F(i-1)+n)
	- where F(i) is the max subarray sum ending at i
*/

// kadane's algorithm - O(n)/O(1)
func maxSubArray(nums []int) int {
	res := nums[0]
	curMax := 0

	for i := range nums {
		curMax = max(nums[i], nums[i]+curMax)
		res = max(res, curMax)
	}

	return res
}

func main() {
	tests := []struct {
		nums string
		out  int
	}{
		{
			nums: "[-2,1,-3,4,-1,2,1,-5,4]",
			out:  6,
		},
		{
			nums: "[1]",
			out:  1,
		},
		{
			nums: "[5,4,-1,7,8]",
			out:  23,
		},
	}

	for _, t := range tests {
		var nums []int
		json.Unmarshal([]byte(t.nums), &nums)
		start := time.Now()
		res := maxSubArray(nums)
		fmt.Printf("%v == %v | %v\n", t.out, res, time.Since(start))
	}
}
