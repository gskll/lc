package main

import (
	"encoding/json"
	"fmt"
	"slices"
)

// https://leetcode.com/problems/combination-sum-ii/description/

// Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target.
//
// Each number in candidates may only be used once in the combination.
//
// Note: The solution set must not contain duplicate combinations.
//
// 1 <= candidates.length <= 100
// 1 <= candidates[i] <= 50
// 1 <= target <= 30

func combinationSum2Optimal(candidates []int, target int) [][]int {
	n := len(candidates)
	var res [][]int
	var curr []int

	slices.Sort(candidates)

	var backtrack func(start, remain int)
	backtrack = func(start, remain int) {
		if remain == 0 {
			res = append(res, append(make([]int, 0, len(curr)), curr...))
			return
		}

		for i := start; i < n; i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			if candidates[i] > remain {
				break
			}

			curr = append(curr, candidates[i])
			backtrack(i+1, remain-candidates[i])
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0, target)
	return res
}

// O(n*2^n) / O(n)
func combinationSum2(candidates []int, target int) [][]int {
	n := len(candidates)
	var res [][]int
	var curr []int

	slices.Sort(candidates)

	var backtrack func(start, remain int)
	backtrack = func(start, remain int) {
		if remain == 0 {
			res = append(res, append(make([]int, 0, len(curr)), curr...))
			return
		}

		if remain < 0 || start == n {
			return
		}

		curr = append(curr, candidates[start])
		backtrack(start+1, remain-candidates[start])
		curr = curr[:len(curr)-1]

		for start+1 < n && candidates[start+1] == candidates[start] {
			start++
		}
		if start+1 < n && candidates[start+1] > remain {
			return
		}
		backtrack(start+1, remain)
	}

	backtrack(0, target)
	return res
}

func main() {
	tests := []struct {
		candidates string
		target     int
		out        string
	}{
		{
			candidates: "[1,2,2,3]",
			target:     5,
			out:        "[[1,2,2],[2,3]]",
		},
		{
			candidates: "[10,1,2,7,6,1,5]",
			target:     8,
			out:        "[ [1,1,6], [1,2,5], [1,7], [2,6] ]",
		},
		{
			candidates: "[2,5,2,1,2]",
			target:     5,
			out:        "[ [1,2,2], [5] ]",
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
		res := combinationSum2Optimal(candidates, t.target)
		fmt.Println("got:", res)
		fmt.Println("exp:", t.out)
	}
}
