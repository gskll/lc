package main

import (
	"encoding/json"
	"fmt"
)

// https://leetcode.com/problems/permutations/description/

// Given an array nums of distinct integers, return all the possible permutations . You can return the answer in any order.
//
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// All the integers of nums are unique.

// Heap's algorithm
// O(n! * n) (*n for copying leaf) / O(n!*n) for result storage
// Pure Heap's algorithm we process each permutation when it's found: O(n!) / O(n)
func permuteHeap(nums []int) [][]int {
	n := len(nums)
	totalPerms := 1
	for i := 2; i <= n; i++ {
		totalPerms *= i
	}

	res := make([][]int, 0, totalPerms)
	var generate func(k int, nums []int)
	generate = func(k int, nums []int) {
		if k == 1 {
			res = append(res, append(make([]int, 0, n), nums...))
			return
		}

		generate(k-1, nums)

		for i := 0; i < k-1; i++ {
			if k%2 == 0 {
				nums[i], nums[k-1] = nums[k-1], nums[i]
			} else {
				nums[0], nums[k-1] = nums[k-1], nums[0]
			}
			generate(k-1, nums)
		}
	}

	generate(n, nums)
	return res
}

// O(n! * n) (n for copying leaf)/ O(n! * n) memory for storing all permutations and recursion stack
func permute(nums []int) [][]int {
	n := len(nums)
	totalPerms := 1
	for i := 2; i <= n; i++ {
		totalPerms *= i
	}

	res := make([][]int, 0, totalPerms)

	used := make([]bool, n)
	curr := make([]int, 0, n)

	var backtrack func()
	backtrack = func() {
		if len(curr) == len(nums) {
			res = append(res, append(make([]int, 0, n), curr...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if !used[i] {
				curr = append(curr, nums[i])
				used[i] = true
				backtrack()
				curr = curr[:len(curr)-1]
				used[i] = false
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
			nums: "[0,1]",
			out:  "[[0,1],[1,0]]",
		},
		{
			nums: "[1]",
			out:  "[[1]]",
		},
	}

	for _, t := range tests {
		var nums []int
		json.Unmarshal([]byte(t.nums), &nums)
		res := permuteHeap(nums)
		fmt.Println("got:", res)
		fmt.Println("exp:", t.out)
	}
}
