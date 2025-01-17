package main

import (
	"errors"
)

// https://leetcode.com/problems/daily-temperatures/description/

// Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature. If there is no future day for which this is possible, keep answer[i] == 0 instead.
//
// 1 <= temperatures.length <= 10^5
// 30 <= temperatures[i] <= 100

// return array of same length
// how many days until we get a warmer temperature?
// index delta between i and next highest value
//
// for each temperature
// check if there's something in the stack, if not, add it
// if there is - check if it's bigger, if it is then write the ∂indices to the top index
// add our currIdxent one
// at the end, if the stack is not empty, write 0s for the remaining

type stack []int

func (s *stack) pop() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("empty stack - can't pop")
	}

	last := len(*s) - 1
	x := (*s)[last]
	*s = (*s)[:last]
	return x, nil
}

func (s *stack) push(x int) {
	*s = append(*s, x)
}

func (s *stack) peek() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("empty stack - can't peak")
	}
	return (*s)[len(*s)-1], nil
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func dailyTemperatures(temperatures []int) []int {
	days := make([]int, len(temperatures))
	var s stack

	for currIdx := range temperatures {
		for !s.isEmpty() {
			prevIdx, err := s.peek()
			if err != nil || temperatures[currIdx] <= temperatures[prevIdx] {
				break
			}
			days[prevIdx] = currIdx - prevIdx
			s.pop()
		}
		s.push(currIdx)
	}

	return days
}
