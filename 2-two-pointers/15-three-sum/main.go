package main

import (
	"fmt"
	"slices"
)

// https://leetcode.com/problems/3sum/

func threeSum(nums []int) [][]int {
	var solutions [][]int
	slices.Sort(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		tgt := -nums[i]
		l, r := i+1, len(nums)-1

		for l < r {
			if nums[l]+nums[r] == tgt {
				solutions = append(solutions, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if nums[l]+nums[r] < tgt {
				l++
			} else if nums[l]+nums[r] > tgt {
				r--
			}
		}
	}

	return solutions
}

func main() {
	type test struct {
		numbers []int
		out     [][]int
	}
	tests := []test{
		{numbers: []int{-1, 0, 1, 2, -1, -4}, out: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{numbers: []int{0, 1, 1}, out: [][]int{}},
		{numbers: []int{0, 0, 0}, out: [][]int{{0, 0, 0}}},
		{numbers: []int{0, 0, 0, 0}, out: [][]int{{0, 0, 0}}},
	}
	for i, t := range tests {
		res := threeSum(t.numbers)
		fmt.Printf("%d: %+vv --> %+v: %+v\n", i, t.numbers, res, t.out)
	}
}
