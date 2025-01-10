package main

import (
	"fmt"
	"time"
)

// https://leetcode.com/problems/decode-ways/description/

// You have intercepted a secret message encoded as a string of numbers. The message is decoded via the following mapping:
//
// "1" -> 'A'
//
// "2" -> 'B'
//
// ...
//
// "25" -> 'Y'
//
// "26" -> 'Z'
//
// However, while decoding the message, you realize that there are many different ways you can decode the message because some codes are contained in other codes ("2" and "5" vs "25").
//
// For example, "11106" can be decoded into:
//
// "AAJF" with the grouping (1, 1, 10, 6)
// "KJF" with the grouping (11, 10, 6)
// The grouping (1, 11, 06) is invalid because "06" is not a valid code (only "6" is valid).
// Note: there may be strings that are impossible to decode.
//
// Given a string s containing only digits, return the number of ways to decode it. If the entire string cannot be decoded in any valid way, return 0.
//
// The test cases are generated so that the answer fits in a 32-bit integer.
//
// 1 <= s.length <= 100
// s contains only digits and may contain leading zero(s).

/*
	- counting problem - combinatorics?
	- avoid counting duplicates
	- valid codes are 1-26 -> code - 'A' + 1
	- 0 by itself is not valid, nor is 0 prefix - if first digit is 0 then invalid return 0
	- answer fits in int32 so golang int is fine
	- len(s) is small - could be complicated
	- each digit could be it's own grouping, or in a grouping with next digit if lower than 26
		- iterate forwards will avoid having to look back

	- decision tree? skip or pick or pick with next

	- recurrence relation? no side effects?
	- base case - string doesn't start with 0
	- base case - reach end

	- F(i) = F(i+1) // if s[i] is valid, != '0'
				+ F(i+2) // if s[i:i+2] is valid, [10-26]
	- F(n) = 1
	- F(i) = 0 // if s[i] == '0'
*/

// bottom up - O(n) / O(1)
// space saving as constant dependencies for dp access
func numDecodings(s string) int {
	curr, prev1, prev2 := 0, 1, 0
	for i := len(s) - 1; i >= 0; i-- {
		curr = 0
		if s[i] != '0' {
			curr = prev1
		}

		if i+1 < len(s) && (s[i] == '1' || s[i] == '2' && s[i+1] <= '6') {
			curr += prev2
		}

		prev2 = prev1
		prev1 = curr
	}
	return curr
}

// bottom up - O(n)/O(n)
// func numDecodings(s string) int {
// 	n := len(s)
// 	dp := make([]int, n+1)
// 	for i := range dp {
// 		dp[i] = -1
// 	}
//
// 	// i: [0,n]
// 	// i requires i+1, i+2
// 	// big before small
// 	// n->0
//
// 	dp[n] = 1
// 	for i := n - 1; i >= 0; i-- {
// 		if s[i] == '0' {
// 			dp[i] = 0
// 			continue
// 		}
//
// 		dp[i] = dp[i+1]
//
// 		if i < n-1 && (s[i] == '1' || s[i] == '2' && s[i+1] <= '6') {
// 			dp[i] += dp[i+2]
// 		}
// 	}
//
// 	return dp[0]
// }

// O(n) / O(n) - topdown
// func numDecodings(s string) int {
// 	n := len(s)
// 	dp := make([]int, n)
// 	for i := range dp {
// 		dp[i] = -1
// 	}
//
// 	var count func(i int) int
// 	count = func(i int) int {
// 		if i == n {
// 			return 1
// 		}
// 		if dp[i] != -1 {
// 			return dp[i]
// 		}
// 		if s[i] == '0' {
// 			return 0
// 		}
//
// 		dp[i] = count(i + 1)
//
// 		if i < n-1 && (s[i] == '1' || s[i] == '2' && s[i+1] <= '6') {
// 			dp[i] += count(i + 2)
// 		}
//
// 		return dp[i]
// 	}
//
// 	return count(0)
// }

func main() {
	tests := []struct {
		in  string
		out int
	}{
		{
			in:  "12",
			out: 2,
		},
		{
			in:  "226",
			out: 3,
		},
		{
			in:  "06",
			out: 0,
		},
	}

	for _, t := range tests {
		start := time.Now()
		res := numDecodings(t.in)
		fmt.Printf("%v == %v | %v\n", t.out, res, time.Since(start))
	}
}
