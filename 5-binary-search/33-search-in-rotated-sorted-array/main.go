package main

import "fmt"

// https://leetcode.com/problems/search-in-rotated-sorted-array/

// There is an integer array nums sorted in ascending order (with distinct values).
//
// Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].
//
// Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.
//
// You must write an algorithm with O(log n) runtime complexity.
//
// 1 <= nums.length <= 5000
// -10^4 <= nums[i] <= 10^4
// All values of nums are unique.
// nums is an ascending array that is possibly rotated.
// -10^4 <= target <= 10^4

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1

	for l <= r {
		i := l + (r-l)/2
		n := nums[i]

		fmt.Println(l, r, i, n)
		if n == target {
			return i
		}

		// array split in two sorted portions
		if n > nums[r] { // in first portion
			if target > nums[r] && target <= n { // target > nums[r]: target is in first portion / target <= n: target is in first portion and to the left of mid
				r = i - 1
			} else {
				l = i + 1
			}
		} else { // in second portion
			if target <= nums[r] && target > n { // target <= nums[r]: target is in second portion / target > n: target is to the right of n
				l = i + 1
			} else {
				r = i - 1
			}
		}

	}

	return -1
}

func main() {
	type test struct {
		nums []int
		t    int
		out  int
	}
	tests := []test{
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, t: 0, out: 4},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, t: 3, out: -1},
		{nums: []int{1}, t: 0, out: -1},
		{nums: []int{3, 5, 1}, t: 3, out: 0},
		{nums: []int{3, 1}, t: 1, out: 1},
	}
	for i, t := range tests {
		res := search(t.nums, t.t)
		fmt.Printf("%d: %+v --> %+v == %+v\n", i, t.nums, res, t.out)
	}
}
