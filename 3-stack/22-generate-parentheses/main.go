package main

import (
	"strings"
)

// https://leetcode.com/problems/generate-parentheses/description/

// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
//
// 1 <= n <= 8

func generateParenthesis(n int) []string {
	var stack []string
	var combinations []string

	var choose func(int, int)
	choose = func(numOpen, numClose int) {
		if numOpen == n && numClose == n && numOpen == numClose {
			combinations = append(combinations, strings.Join(stack, ""))
			return
		}

		if numOpen < n {
			stack = append(stack, "(")
			choose(numOpen+1, numClose)
			stack = stack[:len(stack)-1]
		}

		if numClose < numOpen {
			stack = append(stack, ")")
			choose(numOpen, numClose+1)
			stack = stack[:len(stack)-1]
		}
	}
	choose(0, 0)
	return combinations
}
