package main

import (
	"fmt"
)

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var s1Count [26]int
	var s2Count [26]int

	for i := range s1 {
		s1Count[s1[i]-'a'] += 1
		s2Count[s2[i]-'a'] += 1
	}

	matches := 0
	for i := 0; i < 26; i++ {
		if s1Count[i] == s2Count[i] {
			matches++
		}
	}

	l, r := 0, len(s1)
	for r < len(s2) {
		if matches == 26 {
			return true
		}

		index := s2[r] - 'a'
		s2Count[index] += 1
		if s2Count[index] == s1Count[index] {
			matches++
		} else if s2Count[index] == s1Count[index]+1 {
			matches--
		}

		index = s2[l] - 'a'
		s2Count[index] -= 1
		if s2Count[index] == s1Count[index] {
			matches++
		} else if s2Count[index] == s1Count[index]-1 {
			matches--
		}
		r++
		l++
	}

	return matches == 26
}

func main() {
	type test struct {
		s1  string
		s2  string
		out bool
	}
	tests := []test{
		{s1: "ab", s2: "eidbaooo", out: true},
		{s1: "ab", s2: "eidboaoo", out: false},
		{s1: "adc", s2: "dcda", out: true},
	}
	for i, t := range tests {
		res := checkInclusion(t.s1, t.s2)
		fmt.Printf("%d: %+v --> %+v == %+v\n", i, t.s1, res, t.out)
	}
}
