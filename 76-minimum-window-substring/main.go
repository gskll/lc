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

	// freqT := [128]int{}
	// freqS := [128]int{}
	freqT := make(map[rune]int)
	freqS := make(map[rune]int)

	for _, ch := range t {
		freqT[ch]++
		// freqT[ch]++
	}

	fmt.Println(freqT)

	window := ""
	l, r := 0, 0
	matches := 0

	for r < len(s) {
		newChar := rune(s[r])
		fmt.Println(newChar)
		freqS[rune(newChar)]++
		// freqS[newChar]++

		if freqS[newChar] == freqT[newChar] {
			matches++
		} else if freqS[newChar]-1 == freqT[newChar] {
			matches--
		}

		fmt.Println(matches)

		for l < r && matches == len(t) {
			fmt.Println(matches)
			if r-l+1 < len(window) {
				fmt.Println(s[l : r+1])
				window = s[l : r+1]
			}
			oldChar := rune(s[l])
			if freqS[oldChar] > 0 {
				freqT[oldChar]--
			}

			if freqS[oldChar] == freqT[oldChar] {
				matches++
			} else if freqS[oldChar]+1 == freqT[oldChar] {
				matches--
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
		// {s1: "a", s2: "a", out: "a"},
		// {s1: "a", s2: "aa", out: ""},
	}
	for i, t := range tests {
		res := minWindow(t.s1, t.s2)
		fmt.Printf("%d: %+v --> %+v == %+v\n", i, t.s1, res, t.out)
	}
}
