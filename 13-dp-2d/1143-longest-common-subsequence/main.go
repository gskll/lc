package main

import (
	"fmt"
	"time"
)

// https://leetcode.com/problems/longest-common-subsequence/description/
//
// Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.
//
// A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.
//
//     For example, "ace" is a subsequence of "abcde".
//
// A common subsequence of two strings is a subsequence that is common to both strings.
//
//
//
// Example 1:
//
// Input: text1 = "abcde", text2 = "ace"
// Output: 3
// Explanation: The longest common subsequence is "ace" and its length is 3.
//
// Example 2:
//
// Input: text1 = "abc", text2 = "abc"
// Output: 3
// Explanation: The longest common subsequence is "abc" and its length is 3.
//
// Example 3:
//
// Input: text1 = "abc", text2 = "def"
// Output: 0
// Explanation: There is no such common subsequence, so the result is 0.
//
//
//
// Constraints:
//
//     1 <= text1.length, text2.length <= 1000
//     text1 and text2 consist of only lowercase English characters.
//
//

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)

	memo := make([][]int, m+1)

	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == m || j == n {
			return 0
		}

		if memo[i][j] > 0 {
			return memo[i][j]
		}

		if text1[i] == text2[j] {
			memo[i][j] = 1 + dfs(i+1, j+1)
		} else {
			memo[i][j] = max(dfs(i+1, j), dfs(i, j+1))
		}

		return memo[i][j]
	}

	return dfs(0, 0)
}

func longestCommonSubsequenceRecursive(text1 string, text2 string) int {
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == len(text1) || j == len(text2) {
			return 0
		}

		if text1[i] == text2[j] {
			return 1 + dfs(i+1, j+1)
		}

		r1 := dfs(i+1, j)
		r2 := dfs(i, j+1)
		return max(r1, r2)
	}

	return dfs(0, 0)
}

func main() {
	tests := []struct {
		text1 string
		text2 string
		out   int
	}{
		{
			text1: "abcde",
			text2: "ace",
			out:   3,
		},
		{
			text1: "abc",
			text2: "abc",
			out:   3,
		},
		{
			text1: "abc",
			text2: "def",
			out:   0,
		},
	}

	for _, t := range tests {
		start := time.Now()
		res := longestCommonSubsequence(t.text1, t.text2)
		fmt.Printf("%v == %v | %v\n", t.out, res, time.Since(start))
	}
}
