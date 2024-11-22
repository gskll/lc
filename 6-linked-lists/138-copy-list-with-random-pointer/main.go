package main

import (
	"fmt"
)

// https://leetcode.com/problems/copy-list-with-random-pointer/description/

// A linked list of length n is given such that each node contains an additional random pointer, which could point to any node in the list, or null.
//
// Construct a deep copy of the list. The deep copy should consist of exactly n brand new nodes, where each new node has its value set to the value of its corresponding original node. Both the next and random pointer of the new nodes should point to new nodes in the copied list such that the pointers in the original list and copied list represent the same list state. None of the pointers in the new list should point to nodes in the original list.
//
// For example, if there are two nodes X and Y in the original list, where X.random --> Y, then for the corresponding two nodes x and y in the copied list, x.random --> y.
//
// Return the head of the copied linked list.
//
// The linked list is represented in the input/output as a list of n nodes. Each node is represented as a pair of [val, random_index] where:
//
// val: an integer representing Node.val
// random_index: the index of the node (range from 0 to n-1) that the random pointer points to, or null if it does not point to any node.
// Your code will only be given the head of the original linked list.
//
// 0 <= n <= 1000
// -104 <= Node.val <= 104
// Node.random is null or is pointing to some node in the linked list.

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	curr := head

	for curr != nil {
		cpy := &Node{Val: curr.Val, Next: curr.Next}
		curr.Next = cpy
		curr = cpy.Next
	}

	curr = head
	for curr != nil {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		}
		curr = curr.Next.Next
	}

	curr = head
	cpyHead := head.Next
	cpyCurr := head.Next
	for curr != nil {
		curr.Next = curr.Next.Next
		if cpyCurr.Next != nil {
			cpyCurr.Next = cpyCurr.Next.Next
		}
		curr = curr.Next
		cpyCurr = cpyCurr.Next
	}

	return cpyHead
}

// using single pass hash map to store copied nodes
// func copyRandomList(head *Node) *Node {
// 	nodes := make(map[*Node]*Node)
//
// 	curr := head
//
// 	for curr != nil {
// 		cpy, ok := nodes[curr]
// 		if !ok {
// 			cpy = &Node{Val: curr.Val}
// 			nodes[curr] = cpy
// 		}
//
// 		if curr.Next != nil {
// 			nextCpy, ok := nodes[curr.Next]
// 			if !ok {
// 				nextCpy = &Node{Val: curr.Next.Val}
// 				nodes[curr.Next] = nextCpy
// 			}
// 			cpy.Next = nextCpy
// 		}
//
// 		if curr.Random != nil {
// 			randCpy, ok := nodes[curr.Random]
// 			if !ok {
// 				randCpy = &Node{Val: curr.Random.Val}
// 				nodes[curr.Random] = randCpy
// 			}
// 			cpy.Random = randCpy
// 		}
//
// 		curr = curr.Next
// 	}
//
// 	return nodes[head]
// }

func createRandomList(input [][]int) *Node {
	if len(input) == 0 {
		return nil
	}

	nodes := make([]*Node, len(input))
	for i, pair := range input {
		nodes[i] = &Node{Val: pair[0]}
	}

	for i, pair := range input {
		if i < len(input)-1 {
			nodes[i].Next = nodes[i+1]
		}

		if pair[1] != -1 { // assuming -1 represents null
			nodes[i].Random = nodes[pair[1]]
		}
	}

	return nodes[0]
}

func printRandomList(head *Node) {
	curr := head
	for curr != nil {
		randomIndex := -1
		if curr.Random != nil {
			temp := head
			i := 0
			for temp != curr.Random {
				temp = temp.Next
				i++
			}
			randomIndex = i
		}
		fmt.Printf("[%d,%d] -> ", curr.Val, randomIndex)
		curr = curr.Next
	}
	fmt.Println("nil")
}

func main() {
	// Test case 1
	list1 := createRandomList([][]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}})
	printRandomList(list1)
	l1 := copyRandomList(list1)
	printRandomList(l1)

	// Test case 2
	list2 := createRandomList([][]int{{1, 1}, {2, 1}})
	printRandomList(list2)
	l2 := copyRandomList(list2)
	printRandomList(l2)

	// Test case 3
	list3 := createRandomList([][]int{{3, -1}, {3, 0}, {3, -1}})
	printRandomList(list3)
	l3 := copyRandomList(list3)
	printRandomList(l3)
}
