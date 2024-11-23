package main

import (
	"fmt"
)

// https://leetcode.com/problems/lru-cache/

// Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.
//
// Implement the LRUCache class:
//
// LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
// int get(int key) Return the value of the key if the key exists, otherwise return -1.
// void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.
// The functions get and put must each run in O(1) average time complexity
//
// 1 <= capacity <= 3000
// 0 <= key <= 104
// 0 <= value <= 105
// At most 2 * 105 calls will be made to get and put.

type Node struct {
	key, value int
	next, prev *Node
}

type LRUCache struct {
	head, tail *Node
	cache      map[int]*Node
	capacity   int
}

func Constructor(capacity int) LRUCache {
	if capacity <= 0 {
		panic("capacity must be positive")
	}

	// use dummy head/tail nodes to simplify edge cases
	cache := LRUCache{
		head:     &Node{},
		tail:     &Node{},
		capacity: capacity,
		cache:    make(map[int]*Node, capacity),
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head

	return cache
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.cache[key]; exists {
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.cache[key]; exists {
		node.value = value
		this.moveToHead(node)
		return
	}

	node := &Node{key: key, value: value}
	this.addToHead(node)
	this.cache[key] = node

	if len(this.cache) > this.capacity {
		this.evictLRU()
	}
}

func (this *LRUCache) evictLRU() {
	if lru := this.tail.prev; lru != this.head {
		delete(this.cache, lru.key)
		this.removeNode(lru)
	}
}

func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) addToHead(node *Node) {
	node.next = this.head.next
	node.prev = this.head

	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func testLRUCache() {
	// Test cases as arrays of operations and arguments
	// operations := []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"}
	operations := []string{"LRUCache", "put", "put", "put", "put", "get", "get"}
	// arguments := [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}}
	arguments := [][]int{{2}, {2, 1}, {1, 1}, {2, 3}, {4, 1}, {1}, {2}}

	// Expected outputs for each operation
	// expected := []int{0, 0, 0, 1, 0, -1, 0, -1, 3, 4}
	expected := []int{0, 0, 0, 0, 0, -1, 3}

	// Initialize cache with first argument
	var cache LRUCache
	var result int

	fmt.Println("Running LRU Cache tests...")

	for i, op := range operations {
		switch op {
		case "LRUCache":
			cache = Constructor(arguments[i][0])
			result = 0
		case "put":
			cache.Put(arguments[i][0], arguments[i][1])
			result = 0
		case "get":
			result = cache.Get(arguments[i][0])
		}

		// Verify result
		if result != expected[i] {
			fmt.Printf("Test failed at operation %d (%s):\n", i, op)
			fmt.Printf("Expected: %d, Got: %d\n", expected[i], result)
			return
		}

		// Print operation and current state
		printOperation(op, arguments[i], result)
	}

	fmt.Println("All tests passed!")
}

// Helper function to print operations
func printOperation(op string, args []int, result int) {
	switch op {
	case "LRUCache":
		fmt.Printf("Initialize cache with capacity %d\n", args[0])
	case "put":
		fmt.Printf("Put(%d, %d)\n", args[0], args[1])
	case "get":
		fmt.Printf("Get(%d) -> %d\n", args[0], result)
	}
}

// Example usage:
func main() {
	testLRUCache()
}
