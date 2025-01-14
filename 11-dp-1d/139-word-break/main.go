package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// https://leetcode.com/problems/word-break/description/

// Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.
//
// Note that the same word in the dictionary may be reused multiple times in the segmentation.
//
// Constraints:
//
// 1 <= s.length <= 300
// 1 <= wordDict.length <= 1000
// 1 <= wordDict[i].length <= 20
// s and wordDict[i] consist of only lowercase English letters.
// All the strings of wordDict are unique.

/*
	- s.length is medium
	- wordDict length is medium
	- word length is small
	- all lowercase letters

	- same word can be reused multiple times

	- put the wordDict into a trie, then for each segmentation of s we can iterate through the trie and check that everything is an endpoint
	- somehow cache the segmentations?

	- leetcode, ["leet", "code"]
	- build trie with "leet", "code"
	- iterate through string, if we reach a valid trie word, start with next segment
		- if we reach end of string and all segments are valid trie words, then true -> base case, when startP == n
		- if we don't get a valid word, we can back track on previous segment, see if it can be exteneded

	- two pointers plus stack for segments
		- valid segment goes on stack, if next segment isn't valid we pop off stack and keep exploring
	- backtracking...

	- if current segment is a valid word, we need to explore rest of string
	- F(i) = OR(j=i to n) {
		isWord(s[i:j]) && dp[j]
	}
	- F(n) = true

	- purely recursive: 2^n
	- wordBreak(s) = OR(i=1 to n) {
		wordDict[0:i] AND wordBreak(s[i:])
	}
*/

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func (t *Trie) Insert(word string) {
	curr := t
	for _, c := range word {
		if curr.children[c-'a'] == nil {
			curr.children[c-'a'] = &Trie{}
		}
		curr = curr.children[c-'a']
	}
	curr.isEnd = true
}

// O(n*L) where n is string lenght, L is length of longest word in dictionary
// for each position i, we check each possible length from i to n, but this is bounded by L
func wordBreakTopDown(s string, wordDict []string) bool {
	n := len(s)
	root := &Trie{}
	for _, word := range wordDict {
		root.Insert(word)
	}

	dp := make([]int, len(s)+1)
	for i := range dp {
		dp[i] = -1
	}

	var explore func(start int) bool
	explore = func(start int) bool {
		if start == n {
			return true
		}

		if dp[start] != -1 {
			return dp[start] == 1
		}

		curr := root
		for i := start; i < n; i++ {
			curr = curr.children[s[i]-'a']
			if curr == nil {
				break
			}
			if curr.isEnd && explore(i+1) {
				dp[start] = 1
				return true
			}
		}

		dp[start] = 0
		return false
	}

	return explore(0)
}

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	root := &Trie{}
	for _, word := range wordDict {
		root.Insert(word)
	}

	dp := make([]bool, n+1)
	dp[n] = true

	for start := n - 1; start >= 0; start-- {
		curr := root
		for i := start; i < n; i++ {
			curr = curr.children[s[i]-'a']
			if curr == nil {
				break
			}
			if curr.isEnd && dp[i+1] {
				dp[start] = true
				break
			}
		}
	}

	return dp[0]
}

func main() {
	tests := []struct {
		s        string
		wordDict string
		out      bool
	}{
		{
			s:        "leetcode",
			wordDict: `["leet","code"]`,
			out:      true,
		},
		{
			s:        "applepenapple",
			wordDict: `["apple","pen"]`,
			out:      true,
		},
		{
			s:        "catsandog",
			wordDict: `["cats","dog","sand","and","cat"]`,
			out:      false,
		},
		{
			s:        "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
			wordDict: `["a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"]`,
			out:      false,
		},
		{
			s:        "aaaaaaa",
			wordDict: `["aaaa","aaa"]`,
			out:      true,
		},
		{
			s:        "bb",
			wordDict: `["a","b","bbb","bbbb"]`,
			out:      true,
		},
		{
			s:        "abcd",
			wordDict: `["a","abc","b","cd"]`,
			out:      true,
		},
	}

	for _, t := range tests {
		var wordDict []string
		json.Unmarshal([]byte(t.wordDict), &wordDict)
		start := time.Now()
		res := wordBreak(t.s, wordDict)
		fmt.Printf("%v == %v | %v\n", t.out, res, time.Since(start))
	}
}
