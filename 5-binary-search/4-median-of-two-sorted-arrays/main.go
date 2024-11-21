package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/median-of-two-sorted-arrays/description/

// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
//
// The overall run time complexity should be O(log (m+n)).
//
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -106 <= nums1[i], nums2[i] <= 106

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// median partitions the merged array in 2 equal parts
	//
	// find smallest array
	// find median of smallest array, and complete the left hand partition with elements from larger array
	// have to watch for invalid indices when the partition includes all of the smaller array
	// left-most of each partition should be <= right-most of other array partition
	// if correctly partitioned we can calculate median

	n1, n2 := len(nums1), len(nums2)
	if n1 > n2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	maxLeftX, minRightX, maxLeftY, minRightY := 0, 0, 0, 0

	l, r := 0, n1
	for l <= r {
		iX := l + (r-l)/2
		iY := (n1+n2)/2 - iX

		maxLeftX = getMaxLeft(nums1, iX-1)
		minRightX = getMinRight(nums1, iX)
		maxLeftY = getMaxLeft(nums2, iY-1)
		minRightY = getMinRight(nums2, iY)

		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			// valid partition
			break
		}

		if maxLeftY > minRightX {
			l = iX + 1
		} else {
			r = iX - 1
		}
	}

	if (n1+n2)%2 == 1 {
		median := min(minRightX, minRightY)
		return float64(median)
	}

	left := max(maxLeftX, maxLeftY)
	right := min(minRightX, minRightY)
	return float64(left+right) / 2
}

func getMinRight(nums []int, i int) int {
	if i >= len(nums) {
		return math.MaxInt
	}
	return nums[i]
}

func getMaxLeft(nums []int, i int) int {
	if i < 0 {
		return math.MinInt
	}
	return nums[i]
}

func main() {
	type test struct {
		nums1 []int
		nums2 []int
		out   float64
	}
	tests := []test{
		{nums1: []int{1, 3}, nums2: []int{2}, out: 2.0},
		{nums1: []int{1, 2}, nums2: []int{3, 4}, out: 2.5},
		{nums1: []int{1, 2, 3, 4, 5}, nums2: []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}, out: 9},
	}
	for i, t := range tests {
		res := findMedianSortedArrays(t.nums1, t.nums2)
		fmt.Printf("%d: %+v %+v --> %+v == %+v\n", i, t.nums1, t.nums2, res, t.out)
	}
}
