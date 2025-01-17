package main

import "strconv"

// https://leetcode.com/problems/evaluate-reverse-polish-notation/

// You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.
//
// Evaluate the expression. Return an integer that represents the value of the expression.
//
// Note that:
//
// The valid operators are '+', '-', '*', and '/'.
// Each operand may be an integer or another expression.
// The division between two integers always truncates toward zero.
// There will not be any division by zero.
// The input represents a valid arithmetic expression in a reverse polish notation.
// The answer and all the intermediate calculations can be represented in a 32-bit integer.
//
// 1 <= tokens.length <= 10^4
// tokens[i] is either an operator: "+", "-", "*", or "/", or an integer in the range [-200, 200].

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for _, token := range tokens {
		if i, err := strconv.Atoi(token); err == nil {
			stack = append(stack, i)
			continue
		}

		x, y := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]

		switch token {
		case "+":
			stack = append(stack, x+y)
		case "*":
			stack = append(stack, x*y)
		case "-":
			stack = append(stack, x-y)
		case "/":
			stack = append(stack, x/y)
		}
	}
	return stack[0]
}
