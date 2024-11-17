package main

import (
	"fmt"
	"strings"
)

func isAlphaNum(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b >= '0' && b <= '9'
}

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	l, r := 0, len(s)-1

	s = strings.ToLower(s)
	for l < r {
		for l < r && l < len(s)-1 && !isAlphaNum(s[l]) {
			l++
		}
		for l < r && r > 0 && !isAlphaNum(s[r]) {
			r--
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	type test struct {
		in  string
		out bool
	}
	tests := []test{
		{in: "A man, a plan, a canal: Panama", out: true},
		{in: "race a car", out: false},
		{in: " ", out: true},
	}
	for i, t := range tests {
		res := isPalindrome(t.in)
		fmt.Printf("%d: %+v --> %v: %v\n", i, t.in, res, t.out == res)
	}
}
