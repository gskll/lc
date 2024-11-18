package main

import (
	"fmt"
)

// https://leetcode.com/problems/longest-repeating-character-replacement/description/
// You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most k times.
//
// Return the length of the longest substring containing the same letter you can get after performing the above operations.

func characterReplacement(s string, k int) int {
	if len(s) <= k {
		return len(s)
	}

	// count := make(map[byte]int)
	// more efficient count - allocated on the stack at compile time instead of the heap
	count := [26]int{}
	maxF, length := 0, 0

	for l, r := 0, 0; r < len(s); r++ {
		count[s[r]-'A']++
		maxF = max(maxF, count[s[r]-'A'])

		for (r-l+1)-maxF > k {
			count[s[l]-'A']--
			l++
		}

		length = max(length, r-l+1)
	}

	return length
}

func main() {
	type test struct {
		s   string
		k   int
		out int
	}
	tests := []test{
		{s: "ABAB", k: 2, out: 4},
		{s: "AABABBA", k: 1, out: 4},
	}
	for i, t := range tests {
		res := characterReplacement(t.s, t.k)
		fmt.Printf("%d: %+v --> %+v == %+v\n", i, t.s, res, t.out)
	}
}
