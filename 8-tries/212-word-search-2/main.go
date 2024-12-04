package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// https://leetcode.com/problems/word-search-ii/description/

// Given an m x n board of characters and a list of strings words, return all words on the board.
//
// Each word must be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.
//
// m == board.length
// n == board[i].length
// 1 <= m, n <= 12
// board[i][j] is a lowercase English letter.
// 1 <= words.length <= 3 * 104
// 1 <= words[i].length <= 10
// words[i] consists of lowercase English letters.
// All the strings of words are unique.

// Time Complexity: O(M * N * 4^L) where:
// - M x N is the board size
// - L is max word length
//
// Space Complexity: O(N * L) where:
// - N is number of words
// - L is max word length

type trie struct {
	root *trieNode
}

func newTrie() *trie {
	return &trie{root: &trieNode{}}
}

type trieNode struct {
	val        string
	children   [26]*trieNode
	isTerminal bool
	count      int
}

func (t *trie) insert(word string) {
	curr := t.root
	for _, c := range word {
		if curr.children[c-'a'] == nil {
			curr.children[c-'a'] = &trieNode{val: string(c)}
		}
		curr.children[c-'a'].count++
		curr = curr.children[c-'a']
	}
	curr.isTerminal = true
}

func (t *trie) markFound(word string) {
	curr := t.root
	for _, c := range word {
		if curr.children[c-'a'] != nil {
			curr.children[c-'a'].count--
		}
		curr = curr.children[c-'a']
	}
	curr.isTerminal = false
}

func findWords(board [][]byte, words []string) []string {
	trie := newTrie()
	for _, w := range words {
		trie.insert(w)
	}

	m, n := len(board), len(board[0])
	path := make([]byte, 0, 10)

	var res []string
	var dfs func(path []byte, row, col int, node *trieNode)
	dfs = func(path []byte, row, col int, node *trieNode) {
		if len(res) == len(words) {
			return
		}

		if row < 0 || col < 0 || row >= m || col >= n || board[row][col] == '#' {
			return
		}

		char := board[row][col]
		childNode := node.children[char-'a']

		if childNode == nil || childNode.count == 0 {
			return
		}

		board[row][col] = '#'
		path = append(path, char)

		if childNode.isTerminal {
			res = append(res, string(path))
			trie.markFound(string(path))
		}

		dfs(path, row+1, col, childNode)
		dfs(path, row-1, col, childNode)
		dfs(path, row, col+1, childNode)
		dfs(path, row, col-1, childNode)

		path = path[:len(path)-1]
		board[row][col] = char
	}

	for i := range board {
		for j := range board[i] {
			if len(res) == len(words) {
				return res
			}
			path = path[:0]
			dfs(path, i, j, trie.root)
		}
	}
	return res
}

func main() {
	tests := []struct {
		board string
		words string
		out   string
	}{
		{
			board: `[["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]]`,
			words: `["oath","pea","eat","rain"]`,
			out:   `["eat","oath"]`,
		},
		{
			board: `[["a","b"],["c","d"]]`,
			words: `["abcb"]`,
			out:   `[]`,
		},
		{
			board: `[["o","a","b","n"],["o","t","a","e"],["a","h","k","r"],["a","f","l","v"]]`,
			words: `["oa","oaa"]`,
			out:   `["oa","oaa"]`,
		},
		{
			board: `[["a","a"]]`,
			words: `["aaa"]`,
			out:   `[]`,
		},
		{
			board: `[["a"]]`,
			words: `["a"]`,
			out:   `["a"]`,
		},
		{
			board: `[["a","a"],["a","a"]]`,
			words: `["aaaaa"]`,
			out:   `[]`,
		},
		{
			board: `
			[["o","a","a","n"],
			["e","t","a","e"],
			["i","h","k","r"],
			["i","f","l","v"]]`,
			words: `["oath","pea","eat","rain","hklf", "hf"]`,
			out:   `["oath","eat","hklf","hf"]`,
		},
	}
	for _, test := range tests {
		var board [][]byte
		var words []string
		var out []string
		var boardRaw [][]string
		json.Unmarshal([]byte(test.board), &boardRaw)
		board = make([][]byte, len(boardRaw))
		for i := range boardRaw {
			board[i] = make([]byte, len(boardRaw[i]))
			for j := range boardRaw[i] {
				board[i][j] = byte(boardRaw[i][j][0])
			}
		}
		json.Unmarshal([]byte(test.words), &words)
		json.Unmarshal([]byte(test.out), &out)
		// fmt.Println("board:", board)
		// fmt.Println("words:", words)
		// fmt.Println("out:", out)
		fmt.Println()
		start := time.Now()
		res := findWords(board, words)
		fmt.Printf("time: %s\n", time.Since(start))
		fmt.Printf("got: %+v\n", res)
		fmt.Printf("exp: %+v\n", out)
	}
}
