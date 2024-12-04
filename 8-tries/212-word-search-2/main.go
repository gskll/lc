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

type path struct {
	chars   []byte
	visited map[[2]int]bool
}

func (p *path) visit(row, col int) {
	p.visited[[2]int{row, col}] = true
}

func (p *path) unvisit(row, col int) {
	p.visited[[2]int{row, col}] = false
}

func (p *path) isVisited(row, col int) bool {
	return p.visited[[2]int{row, col}]
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

	var res []string
	var dfs func(path path, row, col int, node *trieNode)
	dfs = func(path path, row, col int, node *trieNode) {
		if len(res) == len(words) || node == nil || node.count == 0 {
			return
		}
		// fmt.Println(row, col, path.visited, string(path.chars), fmt.Sprintf("<%p - %s - %t", node, node.val, node.isTerminal))

		if node.isTerminal {
			// fmt.Println("path found:", string(path))
			res = append(res, string(path.chars))
			trie.markFound(string(path.chars))
		}

		path.visit(row, col)

		directions := [4][2]int{{+1, 0}, {0, +1}, {-1, 0}, {0, -1}}

		for _, direction := range directions {
			newRow, newCol := row+direction[0], col+direction[1]

			if path.isVisited(newRow, newCol) || newRow < 0 || newCol < 0 || newRow >= m || newCol >= n {
				continue
			}

			cell := board[newRow][newCol]
			childNode := node.children[cell-'a']

			if childNode == nil || childNode.count == 0 {
				continue
			}

			path.chars = append(path.chars, cell)
			dfs(path, newRow, newCol, childNode)
			path.chars = path.chars[:len(path.chars)-1]
		}

		path.unvisit(row, col)
	}

	for i := range board {
		for j := range board[i] {
			if len(res) == len(words) {
				return res
			}
			startChar := board[i][j]
			path := path{
				chars:   []byte{startChar},
				visited: map[[2]int]bool{},
			}
			dfs(path, i, j, trie.root.children[startChar-'a'])
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
