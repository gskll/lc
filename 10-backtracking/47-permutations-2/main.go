package main

import (
	"encoding/json"
	"fmt"
	"slices"
)

// https://leetcode.com/problems/permutations-ii/

// Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order.
//
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10

func permuteUnique2(nums []int) [][]int {
	n := len(nums)
	totalPerms := 1
	for i := 2; i <= n; i++ {
		totalPerms *= i
	}

	slices.Sort(nums)

	res := make([][]int, 0, totalPerms)

	used := make([]bool, n)
	curr := make([]int, 0, n)

	var backtrack func(level int)
	backtrack = func(level int) {
		if level == n {
			res = append(res, append(make([]int, 0, n), curr...))
			return
		}

		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			// we only use duplicates if the previous duplicate is being used
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			curr = append(curr, nums[i])
			used[i] = true
			backtrack(level + 1)
			curr = curr[:len(curr)-1]
			used[i] = false
		}
	}

	backtrack(0)
	return res
}

// O(n! * n) (n for copying leaf)/ O(n! * n) memory for storing all permutations and recursion stack
func permuteUnique(nums []int) [][]int {
	n := len(nums)
	totalPerms := 1
	for i := 2; i <= n; i++ {
		totalPerms *= i
	}

	res := make([][]int, 0, totalPerms)
	curr := make([]int, 0, n)

	numsMap := make(map[int]int)
	for _, n := range nums {
		numsMap[n]++
	}

	var backtrack func()
	backtrack = func() {
		if len(curr) == len(nums) {
			res = append(res, append(make([]int, 0, n), curr...))
			return
		}

		for n, count := range numsMap {
			if count > 0 {
				curr = append(curr, n)
				numsMap[n]--
				backtrack()
				numsMap[n]++
				curr = curr[:len(curr)-1]
			}
		}
	}

	backtrack()
	return res
}

func main() {
	tests := []struct {
		nums string
		out  string
	}{
		{
			nums: "[1,2,3]",
			out:  "[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]",
		},
		{
			nums: "[1,1,2]",
			out:  "[[1,1,2], [1,2,1], [2,1,1]]",
		},
	}

	for _, t := range tests {
		var nums []int
		json.Unmarshal([]byte(t.nums), &nums)
		res := permuteUnique2(nums)
		fmt.Println("got:", res)
		fmt.Println("exp:", t.out)
	}
}
