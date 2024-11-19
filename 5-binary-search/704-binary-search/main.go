package main

import (
	"fmt"
)

// https://leetcode.com/problems/binary-search/description/

// Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.
//
// You must write an algorithm with O(log n) runtime complexity.
//
// 1 <= nums.length <= 10^4
// -10^4 < nums[i], target < 10^4
// All the integers in nums are unique.
// nums is sorted in ascending order.

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		i := (l + r) / 2

		if nums[i] == target {
			return i
		}

		if nums[i] >= target {
			r = i - 1
		} else {
			l = i + 1
		}

	}

	return -1
}

func main() {
	type test struct {
		nums   []int
		target int
		out    int
	}
	tests := []test{
		{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9, out: 4},
		{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2, out: -1},
	}
	for i, t := range tests {
		res := search(t.nums, t.target)
		fmt.Printf("%d: %+v %d --> %+v == %+v\n", i, t.nums, t.target, res, t.out)
	}
}
