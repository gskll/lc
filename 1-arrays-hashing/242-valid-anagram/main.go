package main

// https://leetcode.com/problems/valid-anagram/description/

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
//
// 1 <= s.length, t.length <= 5 * 10^4
// s and t consist of lowercase English letters.

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var count [26]int
	for i := range s {
		count[t[i]-'a']++
		count[s[i]-'a']--
	}

	for i := range count {
		if count[i] != 0 {
			return false
		}
	}
	return true
}
