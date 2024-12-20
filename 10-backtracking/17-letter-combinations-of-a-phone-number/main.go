package main

import (
	"fmt"
	"time"
)

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/

// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.
//
// A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
//
// 0 <= digits.length <= 4
// digits[i] is a digit in the range ['2', '9'].

var digitMapping = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	res := make([]string, 0, 3^len(digits))
	var curr []byte
	var backtrack func(i int)
	backtrack = func(i int) {
		if i >= len(digits) {
			res = append(res, string(curr))
			return
		}

		for _, c := range digitMapping[digits[i]] {
			curr = append(curr, c)
			backtrack(i + 1)
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0)
	return res
}

func main() {
	tests := []struct {
		digits string
		out    string
	}{
		{
			digits: "23",
			out:    `["ad","ae","af","bd","be","bf","cd","ce","cf"]`,
		},
		{
			digits: "",
			out:    "",
		},
	}

	for _, test := range tests {
		fmt.Println()
		start := time.Now()
		res := letterCombinations(test.digits)
		fmt.Printf("time: %s\n", time.Since(start))
		fmt.Printf("got: %+v\n", res)
		fmt.Printf("exp: %+v\n", test.out)
	}
}
