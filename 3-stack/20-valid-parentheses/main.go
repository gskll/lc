package main

// https://leetcode.com/problems/valid-parentheses/

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// An input string is valid if:
//
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.
//
// 1 <= s.length <= 10^4
// s consists of parentheses only '()[]{}'.

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 == 1 {
		return false
	}

	brackets := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	matches := make([]byte, 0, len(s)/2)
	for i := 0; i < len(s); i++ {
		char := s[i]
		if opening, isClosing := brackets[char]; isClosing {
			if len(matches) == 0 || matches[len(matches)-1] != opening {
				return false
			}
			matches = matches[:len(matches)-1]
		} else {
			matches = append(matches, char)
		}
	}

	return len(matches) == 0
}
