package main

import (
	"fmt"
	"time"
)

// https://leetcode.com/problems/palindrome-partitioning/description/

// Given a string s, partition s such that every substring of the partition is a palindrome . Return all possible palindrome partitioning of s.
//
// 1 <= s.length <= 16
// s contains only lowercase English letters.

func isPalindrome(s string, l, r int) bool {
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func partition(s string) [][]string {
	var res [][]string
	var currPath []string
	var backtrack func(l, r int)
	backtrack = func(l, r int) {
		if r >= len(s) {
			if l == r {
				res = append(res, append(make([]string, 0, len(currPath)), currPath...))
			}
			return
		}

		if isPalindrome(s, l, r) {
			currPath = append(currPath, s[l:r+1])
			backtrack(r+1, r+1)
			currPath = currPath[:len(currPath)-1]
		}

		backtrack(l, r+1)
	}

	backtrack(0, 0)
	return res
}

func main() {
	tests := []struct {
		s   string
		out string
	}{
		{
			s:   "aab",
			out: `[["a","a","b"],["aa","b"]]`,
		},
		{
			s:   "a",
			out: `[["a"]]`,
		},
	}

	for _, test := range tests {
		fmt.Println()
		start := time.Now()
		res := partition(test.s)
		fmt.Printf("time: %s\n", time.Since(start))
		fmt.Printf("got: %+v\n", res)
		fmt.Printf("exp: %+v\n", test.out)
	}
}
