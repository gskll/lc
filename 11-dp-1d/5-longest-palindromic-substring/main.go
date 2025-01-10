package main

import (
	"fmt"
	"time"
)

// https://leetcode.com/problems/longest-palindromic-substring/description/

// Given a string s, return the longest palindromic substring in s.
//
// 1 <= s.length <= 1000
// s consist of only digits and English letters.

/*
	- looking for longest so optimization problem
	- palindrome is where string reads same forwards as in reverse so order matters and can't change
	- if there are multiple palindromic substrings of equal length we want to return the first one
	- need to track the longest found so far
	- brute force - check all substrings and check if they're palindromes O(n^2) * O(n)
		- can we cache the palindrome check

	- dynamic programming: can check if substrings of all lengths are palindromes
	- base cases are for len=1, len=2 to handle odd and even

	- can then go over each substring of each length and check whether the characters are the same, and the inner characters are palindromes
	- in which case this substring is a palindrome, and we can update accordingly
*/

// optimal refactored - expand around the center but handle duplicates in one loop
func longestPalindrome(s string) string {
	start, maxLen := 0, 0

	for i := range s {
		l, r := i-1, i+1
		for l >= 0 && s[i] == s[l] {
			l--
		}
		for r < len(s) && s[i] == s[r] {
			r++
		}

		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}

		if r-l-1 > maxLen {
			start = l + 1
			maxLen = r - l - 1
		}
	}

	return s[start : start+maxLen]
}

// optimal solution - expand around the center of each letter checking for palindrome
// O(1) space
// have to handle odd/even length palindromes e.g. aba and abba
// func longestPalindrome(s string) string {
// 	start, maxLen := 0, 0
//
// 	for i := range s {
// 		// odd length
// 		l, r := i, i
// 		for l >= 0 && r < len(s) && s[l] == s[r] {
// 			if r-l+1 > maxLen {
// 				start = l
// 				maxLen = r - l + 1
// 			}
// 			l--
// 			r++
// 		}
//
// 		// even
// 		l, r = i, i+1
// 		for l >= 0 && r < len(s) && s[l] == s[r] {
// 			if r-l+1 > maxLen {
// 				start = l
// 				maxLen = r - l + 1
// 			}
// 			l--
// 			r++
// 		}
// 	}
//
// 	return s[start : start+maxLen]
// }

// O(n^2) / O(n^2)
// func longestPalindrome(s string) string {
// 	n := len(s)
// 	start, maxLen := 0, 0
// 	dp := make([][]bool, n)
// 	for i := range dp {
// 		dp[i] = make([]bool, n)
// 	}
//
// 	for i := n - 1; i >= 0; i-- {
// 		for j := i; j < n; j++ {
// 			if s[i] == s[j] && (j-i <= 2 || dp[i-1][j+1]) {
// 				dp[i][j] = true
//
// 				if j-i+1 >= maxLen {
// 					start = i
// 					maxLen = j - i + 1
// 				}
// 			}
// 		}
// 	}
// 	return s[start : start+maxLen]
// }

// O(n^2)/ O(n^2)
// func longestPalindrome(s string) string {
// 	start, maxLen := 0, 1
// 	n := len(s)
//
// 	dp := make([][]bool, n)
// 	for i := range dp {
// 		dp[i] = make([]bool, n)
// 		dp[i][i] = true
//
// 		if i < n-1 && s[i] == s[i+1] {
// 			dp[i][i+1] = true
// 			start = i
// 			maxLen = 2
// 		}
// 	}
//
// 	for length := 3; length <= n; length++ {
// 		for i := 0; i <= n-length; i++ {
// 			j := i + length - 1
// 			if s[i] == s[j] && dp[i+1][j-1] {
// 				dp[i][j] = true
// 				if length > maxLen {
// 					maxLen = length
// 					start = i
// 				}
// 			}
// 		}
// 	}
//
// 	return s[start : start+maxLen]
// }

// brute force O(n^3) / O(1)
// func isPalindrome(s string) bool {
// 	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
// 		if s[i] != s[j] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func longestPalindrome(s string) string {
// 	var longest string
// 	for i := 0; i < len(s); i++ {
// 		for j := i + 1; j < len(s); j++ {
// 			if isPalindrome(s[i:j+1]) && j-1+1 > len(longest) {
// 				longest = s[i : j+1]
// 			}
// 		}
// 	}
//
// 	return longest
// }

func main() {
	tests := []struct {
		in  string
		out string
	}{
		{
			in:  "babad",
			out: "bab",
		},
		{
			in:  "cbbd",
			out: "bb",
		},
	}

	for _, t := range tests {
		start := time.Now()
		res := longestPalindrome(t.in)
		fmt.Printf("%v == %v | %v\n", t.out, res, time.Since(start))
	}
}
