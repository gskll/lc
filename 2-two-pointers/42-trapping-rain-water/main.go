package main

import (
	"fmt"
)

// https://leetcode.com/problems/trapping-rain-water
// Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

// for each height, we can calculate how much water it can trap by taking the smallest of the max left and right bars, minus the height
// min(max(left), max(right)) - height
// can be done in O(n) complexity O(n) space with prefix/suffix sums
// or O(n) complexity and O(1) space with two pointers

// two pointers
// using the fact that we only care about the min left/right value
// shift the pointer with the smallest height, and calculate trapped water at that point
// recalculate max boundary for new pointer
func trap(height []int) int {
	sum := 0
	l, r := 0, len(height)-1
	maxLeft, maxRight := height[l], height[r]

	for l < r {
		if maxLeft < maxRight {
			l++
			maxLeft = max(maxLeft, height[l])
			sum += maxLeft - height[l]
		} else {
			r--
			maxRight = max(maxRight, height[r])
			sum += maxRight - height[r]
		}
	}

	return sum
}

// prefix sums
// func trap(height []int) int {
// 	sum := 0
//
// 	maxLeft := make([]int, len(height))
// 	for i := 0; i < len(height); i++ {
// 		if i == 0 {
// 			maxLeft[i] = 0
// 			continue
// 		}
// 		maxLeft[i] = max(height[i-1], maxLeft[i-1])
// 	}
//
// 	maxRight := make([]int, len(height))
// 	for i := len(height) - 1; i >= 0; i-- {
// 		if i == len(height)-1 {
// 			maxRight[i] = 0
// 			continue
// 		}
// 		maxRight[i] = max(height[i+1], maxRight[i+1])
// 	}
//
// 	minRightLeft := make([]int, len(height))
// 	for i := 0; i < len(minRightLeft); i++ {
// 		minRightLeft[i] = min(maxLeft[i], maxRight[i])
// 	}
//
// 	for i := range height {
// 		trapped := minRightLeft[i] - height[i]
// 		if trapped < 0 {
// 			trapped = 0
// 		}
// 		sum += trapped
// 	}
//
// 	return sum
// }

func main() {
	type test struct {
		height []int
		out    int
	}
	tests := []test{
		{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, out: 6},
		{height: []int{4, 2, 0, 3, 2, 5}, out: 9},
	}
	for i, t := range tests {
		res := trap(t.height)
		fmt.Printf("%d: %+v --> %+v == %+v\n", i, t.height, res, t.out)
	}
}
