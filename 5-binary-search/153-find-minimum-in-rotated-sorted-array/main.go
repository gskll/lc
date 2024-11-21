package main

import "fmt"

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

// Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:
//
// [4,5,6,7,0,1,2] if it was rotated 4 times.
// [0,1,2,4,5,6,7] if it was rotated 7 times.
// Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].
//
// Given the sorted rotated array nums of unique elements, return the minimum element of this array.
//
// You must write an algorithm that runs in O(log n) time.
//
// n == nums.length
// 1 <= n <= 5000
// -5000 <= nums[i] <= 5000
// All the integers of nums are unique.
// nums is sorted and rotated between 1 and n times.

func findMin(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		return nums[0]
	}

	if nums[0] < nums[n-1] {
		return nums[0]
	}

	l, r := 0, len(nums)-1
	for l < r {
		i := l + (r-l)/2

		if nums[i] > nums[r] {
			l = i + 1
		} else {
			r = i
		}
	}
	return nums[l]
}

func main() {
	type test struct {
		nums []int
		out  int
	}
	tests := []test{
		{nums: []int{3, 4, 5, 1, 2}, out: 1},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, out: 0},
		{nums: []int{11, 13, 15, 17}, out: 11},
	}
	for i, t := range tests {
		res := findMin(t.nums)
		fmt.Printf("%d: %+v --> %+v == %+v\n", i, t.nums, res, t.out)
	}
}
