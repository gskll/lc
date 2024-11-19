package main

import (
	"fmt"
)

// https://leetcode.com/problems/minimum-window-substring/description/
// Given two strings s and t of lengths m and n respectively, return the minimum window substring of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".

// The testcases will be generated such that the answer is unique.

func minWindow(s, t string) string {
	if len(s) < len(t) {
		return ""
	}

	freqT := [128]int{}
	freqS := [128]int{}
	numUnique := 0
	for _, ch := range t {
		if freqT[ch] == 0 {
			numUnique++
		}
		freqT[ch]++
	}

	window := ""
	l, r := 0, 0
	matches := 0

	for r < len(s) {
		newChar := s[r]

		if freqT[newChar] > 0 {
			freqS[newChar]++
			if freqS[newChar] == freqT[newChar] {
				matches++
			}
		}

		for matches == numUnique {
			if r-l+1 < len(window) || len(window) == 0 {
				window = s[l : r+1]
			}
			oldChar := rune(s[l])

			if freqT[oldChar] > 0 {
				freqS[oldChar]--
				if freqS[oldChar]+1 == freqT[oldChar] {
					matches--
				}
			}
			l++
		}

		r++
	}

	return window
}

func main() {
	type test struct {
		s1  string
		s2  string
		out string
	}
	tests := []test{
		{s1: "ADOBECODEBANC", s2: "ABC", out: "BANC"},
		{s1: "a", s2: "a", out: "a"},
		{s1: "a", s2: "aa", out: ""},
		{s1: "aa", s2: "aa", out: "aa"},
	}
	for i, t := range tests {
		res := minWindow(t.s1, t.s2)
		fmt.Printf("%d: %+v --> %+v == %+v\n", i, t.s1, res, t.out)
	}
}
