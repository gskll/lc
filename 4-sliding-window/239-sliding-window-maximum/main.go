package main

import (
	"fmt"
)

// https://leetcode.com/problems/sliding-window-maximum/description/

// You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position.

// Return the max sliding window.

// 1 <= nums.length <= 10^5
// -10^4 <= nums[i] <= 10^4
// 1 <= k <= nums.length

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0, len(nums)-k+1)
	Q := make([]int, len(nums))

	start, end := 0, 0
	for i := range nums {
		for start < end && nums[Q[end-1]] < nums[i] {
			end--
		}
		Q[end] = i
		end++

		if i < k-1 {
			continue
		}

		res = append(res, nums[Q[start]])

		if k == i-Q[start]+1 {
			start++
		}

	}

	return res
}

func main() {
	type test struct {
		nums []int
		k    int
		out  []int
	}
	tests := []test{
		{nums: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3, out: []int{3, 3, 5, 5, 6, 7}},
		{nums: []int{1}, k: 1, out: []int{1}},
		{nums: []int{1, 3, 1, 2, 0, 5}, k: 3, out: []int{3, 3, 2, 5}},
	}
	for i, t := range tests {
		res := maxSlidingWindow(t.nums, t.k)
		fmt.Printf("%d: %+v %d --> %+v == %+v\n", i, t.nums, t.k, res, t.out)
	}
}
