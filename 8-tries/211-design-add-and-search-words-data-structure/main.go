package main

import (
	"fmt"
)

// https://leetcode.com/problems/design-add-and-search-words-data-structure/description/

// Design a data structure that supports adding new words and finding if a string matches any previously added string.
//
// Implement the WordDictionary class:
//
// WordDictionary() Initializes the object.
// void addWord(word) Adds word to the data structure, it can be matched later.
// bool search(word) Returns true if there is any string in the data structure that matches word or false otherwise. word may contain dots '.' where dots can be matched with any letter.
//
// 1 <= word.length <= 25
// word in addWord consists of lowercase English letters.
// word in search consist of '.' or lowercase English letters.
// There will be at most 2 dots in word for search queries.
// At most 104 calls will be made to addWord and search.

type WordDictionary struct {
	children   [26]*WordDictionary
	isTerminal bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this

	for _, char := range word {
		if curr.children[char-'a'] == nil {
			curr.children[char-'a'] = &WordDictionary{}
		}
		curr = curr.children[char-'a']
	}
	curr.isTerminal = true
}

// recursive:  O(26^d * n) where d is number of dots, n is word length / O(n) space
func (this *WordDictionary) Search(word string) bool {
	var dfs func(pos int, node *WordDictionary) bool
	dfs = func(pos int, node *WordDictionary) bool {
		// v1
		if pos == len(word) {
			return node.isTerminal
		}

		char := word[pos]

		if char == '.' {
			for _, child := range node.children {
				if child != nil && dfs(pos+1, child) {
					return true
				}
			}
			return false
		}

		if node.children[char-'a'] == nil {
			return false
		}
		return dfs(pos+1, node.children[char-'a'])

		// v2
		// curr := node
		// for i := pos; i < len(word); i++ {
		// 	char := word[i]
		// 	if char == '.' {
		// 		for _, child := range curr.children {
		// 			if child != nil && dfs(i+1, child) {
		// 				return true
		// 			}
		// 		}
		// 		return false
		// 	}
		//
		// 	if curr.children[char-'a'] == nil {
		// 		return false
		// 	}
		// 	curr = curr.children[char-'a']
		// }
		// return curr.isTerminal
	}

	return dfs(0, this)
}

type queueNode struct {
	node *WordDictionary
	pos  int
}

// iterative - worse because queue overhead
// memory:  O(26^d * n) where d is number of dots, n is word length
// space: O(26^d * n)
func (this *WordDictionary) SearchIt(word string) bool {
	queue := []queueNode{{this, 0}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.pos == len(word) {
			if curr.node.isTerminal {
				return true
			}
			continue
		}

		char := word[curr.pos]
		if char == '.' {
			for _, child := range curr.node.children {
				if child != nil {
					queue = append(queue, queueNode{child, curr.pos + 1})
				}
			}
			continue
		}

		if curr.node.children[char-'a'] != nil {
			queue = append(queue, queueNode{curr.node.children[char-'a'], curr.pos + 1})
		}
	}

	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

func main() {
	wordDictionary := Constructor()
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")
	fmt.Println(wordDictionary.Search("pad"), "== false") // return False
	fmt.Println(wordDictionary.Search("bad"), "== true")  // return True
	fmt.Println(wordDictionary.Search(".ad"), "== true")  // return True
	fmt.Println(wordDictionary.Search("b.."), "== true")  // return True
	// fmt.Println(wordDictionary.SearchIt("pad"), "== false") // return False
	// fmt.Println(wordDictionary.SearchIt("bad"), "== true")  // return True
	// fmt.Println(wordDictionary.SearchIt(".ad"), "== true")  // return True
	// fmt.Println(wordDictionary.SearchIt("b.."), "== true")  // return True
}
