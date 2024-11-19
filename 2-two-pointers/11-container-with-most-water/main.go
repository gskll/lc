package main

import (
	"fmt"
)

// https://leetcode.com/problems/container-with-most-water/
//You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
//
//Find two lines that together with the x-axis form a container, such that the container contains the most water.
//
//Return the maximum amount of water a container can store.
//
//Notice that you may not slant the container.

func maxArea(heights []int) int {
	area := 0
	l, r := 0, len(heights)-1

	for l < r {
		tmpArea := (r - l) * min(heights[l], heights[r])
		area = max(area, tmpArea)

		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}

	return area
}

func main() {
	type test struct {
		height []int
		out    int
	}
	tests := []test{
		{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, out: 49},
		{height: []int{1, 1}, out: 1},
	}
	for i, t := range tests {
		res := maxArea(t.height)
		fmt.Printf("%d: %+vv --> %+v: %+v\n", i, t.height, res, t.out)
	}
}
