package main

import "fmt"

// https://leetcode.com/problems/implement-trie-prefix-tree/description/

// A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.
//
// Implement the Trie class:
//
// Trie() Initializes the trie object.
// void insert(String word) Inserts the string word into the trie.
// boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
// boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.
//
// 1 <= word.length, prefix.length <= 2000
// word and prefix consist only of lowercase English letters.
// At most 3 * 10^4 calls in total will be made to insert, search, and startsWith.

type Trie struct {
	children   [26]*Trie
	isTerminal bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	curr := this
	for _, char := range word {
		if curr.children[char-'a'] == nil {
			curr.children[char-'a'] = &Trie{}
		}
		curr = curr.children[char-'a']
	}
	curr.isTerminal = true
}

func (this *Trie) Search(word string) bool {
	curr := this.traverse(word)
	return curr != nil && curr.isTerminal
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.traverse(prefix)
	return curr != nil
}

func (this *Trie) traverse(chars string) *Trie {
	curr := this
	for _, char := range chars {
		if curr == nil {
			return curr
		}
		curr = curr.children[char-'a']
	}
	return curr
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // return True
	fmt.Println(trie.Search("app"))     // return False
	fmt.Println(trie.StartsWith("app")) // return True
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // return True
}
