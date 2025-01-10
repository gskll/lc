package main

import (
	"fmt"
	"time"
)

// https://leetcode.com/problems/palindromic-substrings/description/

// Given a string s, return the number of palindromic substrings in it.
//
// A string is a palindrome when it reads the same backward as forward.
//
// A substring is a contiguous sequence of characters within the string.
//
// 1 <= s.length <= 1000
// s consists of lowercase English letters.

/*
	- counting problem - combinatorics?
	- all lowercase letters so don't need to transform
	- max string length 1000 - probably looking for O(n2) or better
	- each character is a palindrome with itself
	- need to avoid duplicates

	- brute force would be O(n3) to go over every substring and check if it's a permutation
	- can we cache the permutation check?

	- dynamic programming? check for inner substrings before longer ones
	- middle out? for each character expand and increase counter if palindrome?
		- handle odd/even separately with i or i/i+1 middle?
*/

// O(n^2) / O(n^2) - dynamic programming
// not optimal but to experiment
func countSubstrings(s string) int {
	n := len(s)
	count := 0
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	for left := n - 1; left >= 0; left-- {
		for right := left; right < n; right++ {
			if s[left] == s[right] && (right-left <= 2 || dp[left+1][right-1]) {
				dp[left][right] = true
				count++
			}
		}
	}

	return count
}

// O(n^2) / O(1)
// func countSubstrings(s string) int {
// 	count := 0
// 	for i := range s {
// 		// odd
// 		l, r := i, i
// 		for l >= 0 && r < len(s) && s[l] == s[r] {
// 			count++
// 			l--
// 			r++
// 		}
//
// 		// even
// 		l, r = i, i+1
// 		for l >= 0 && r < len(s) && s[l] == s[r] {
// 			count++
// 			l--
// 			r++
// 		}
// 	}
//
// 	return count
// }

func main() {
	tests := []struct {
		in  string
		out int
	}{
		{
			in:  "abc",
			out: 3,
		},
		{
			in:  "aaa",
			out: 6,
		},
	}

	for _, t := range tests {
		start := time.Now()
		res := countSubstrings(t.in)
		fmt.Printf("%v == %v | %v\n", t.out, res, time.Since(start))
	}
}
