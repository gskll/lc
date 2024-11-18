package main

import "fmt"

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		for numbers[l]+numbers[r] < target {
			l++
		}
		for numbers[l]+numbers[r] > target {
			r--
		}
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		}
	}
	return nil
}

func main() {
	type test struct {
		numbers []int
		target  int
		out     []int
	}
	tests := []test{
		{numbers: []int{2, 7, 11, 15}, target: 9, out: []int{1, 2}},
		{numbers: []int{2, 3, 4}, target: 6, out: []int{1, 3}},
		{numbers: []int{-1, 0}, target: -1, out: []int{1, 2}},
	}
	for i, t := range tests {
		res := twoSum(t.numbers, t.target)
		fmt.Printf("%d: %+v %v --> %+v: %+v\n", i, t.numbers, t.target, res, t.out)
	}
}
