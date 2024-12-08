package main

import (
	"encoding/json"
	"fmt"
	"slices"
)

// https://leetcode.com/problems/combination-sum/description/

// Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.
//
// The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the
// frequency
//  of at least one of the chosen numbers is different.
//
// The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations for the given input.
//
// 1 <= candidates.length <= 30
// 2 <= candidates[i] <= 40
// All elements of candidates are distinct.
// 1 <= target <= 40

// better with sorting as we can return early
// both solutions are O(n^(t/m)) time, O(t/m) space
// n is len candidates, t is target, m is minimum number
// t/m represents recursion stack
func combinationSumOptimal(candidates []int, target int) [][]int {
	var res [][]int

	slices.Sort(candidates)

	var curr []int
	var backtrack func(start, remaining int)
	backtrack = func(start, remaining int) {
		if remaining == 0 {
			res = append(res, append(make([]int, 0, len(curr)), curr...))
			return
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] > remaining {
				break
			}

			curr = append(curr, candidates[i])
			backtrack(i, remaining-candidates[i])
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0, target)

	return res
}

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int

	var curr []int

	var dfs func(currIndex, currTarget int)
	dfs = func(currIndex, currTarget int) {
		if currTarget < 0 {
			return
		}

		if currTarget == 0 {
			res = append(res, append(make([]int, 0, len(curr)), curr...))
			return
		}

		for i := currIndex; i < len(candidates); i++ {
			c := candidates[i]
			curr = append(curr, c)
			dfs(i, currTarget-c)
			curr = curr[:len(curr)-1]
		}
	}

	dfs(0, target)

	return res
}

func main() {
	tests := []struct {
		candidates string
		target     int
		out        string
	}{
		{
			candidates: "[2,3,6,7]",
			target:     7,
			out:        "[[2,2,3],[7]]",
		},
		{
			candidates: "[2,3,5]",
			target:     8,
			out:        "[[2,2,2,2],[2,3,3],[3,5]]",
		},
		{
			candidates: "[2]",
			target:     1,
			out:        "[]",
		},
	}

	for _, t := range tests {
		var candidates []int
		json.Unmarshal([]byte(t.candidates), &candidates)
		res := combinationSumOptimal(candidates, t.target)
		fmt.Println("got:", res)
		fmt.Println("exp:", t.out)
	}
}
