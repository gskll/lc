package main

import "fmt"

// https://leetcode.com/problems/largest-rectangle-in-histogram/description/

// Given an array of integers heights representing the histogram's bar height where the width of each bar is 1, return the area of the largest rectangle in the histogram.
//
// 1 <= heights.length <= 10^5
// 0 <= heights[i] <= 10^4

// for each bar we want to consider how far left/right it can go
// a bar can only not extend left/right if it reaches a bar that is smaller
// can use a stack, pop when we reach a smaller element
// track the left index with the heights
// the right index is the index of the smaller element reached
// or the end of the graph

type bar struct {
	height int
	index  int
}

type stack []bar

func (s *stack) pop() bar {
	if len(*s) == 0 {
		return bar{index: -1}
	}
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

func (s *stack) peek() bar {
	if len(*s) == 0 {
		return bar{index: -1}
	}
	return (*s)[len(*s)-1]
}
func (s *stack) push(x bar)    { *s = append(*s, x) }
func (s *stack) isEmpty() bool { return len(*s) == 0 }

func largestRectangleArea(heights []int) int {
	stack := make(stack, 0, len(heights))
	maxArea := 0

	for i, height := range heights {
		start := i
		for !stack.isEmpty() && height < stack.peek().height {
			bar := stack.pop()
			width := i - bar.index
			area := bar.height * width
			maxArea = max(maxArea, area)
			start = bar.index
		}

		stack.push(bar{height: height, index: start})
	}

	for !stack.isEmpty() {
		bar := stack.pop()
		width := len(heights) - bar.index
		maxArea = max(maxArea, bar.height*width)
	}

	return maxArea
}

func main() {
	type test struct {
		in  []int
		out int
	}
	tests := []test{
		{in: []int{2, 1, 5, 6, 2, 3}, out: 10},
		{in: []int{2, 4}, out: 4},
		{in: []int{5, 4, 1, 2}, out: 8},
	}
	for i, t := range tests {
		res := largestRectangleArea(t.in)
		fmt.Printf("%d: %+v --> %d: %v\n", i, t.in, res, t.out == res)
	}
}
