package main

import (
	"fmt"
)

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
// Given a string s, find the length of the longest substring without repeating characters.

func lengthOfLongestSubstring(s string) int {
	length := 0
	seen := [128]int{}

	for i := range seen {
		seen[i] = -1
	}

	for l, r := 0, 0; r < len(s); r++ {
		char := s[r]
		if i := seen[char]; i >= 0 && l <= i {
			l = i + 1
		}

		seen[char] = r
		length = max(length, r-l+1)

	}

	return length
}

func main() {
	type test struct {
		in  string
		out int
	}
	tests := []test{
		{in: "abcabcbb", out: 3},
		{in: "bbbbb", out: 1},
		{in: "pwwkew", out: 3},
		{in: "abba", out: 2},
	}
	for i, t := range tests {
		res := lengthOfLongestSubstring(t.in)
		fmt.Printf("%d: %+v --> %+v == %+v\n", i, t.in, res, t.out)
	}
}
