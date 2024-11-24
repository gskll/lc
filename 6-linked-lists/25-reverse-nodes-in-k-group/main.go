package main

import (
	"fmt"
)

// https://leetcode.com/problems/reverse-nodes-in-k-group/description/

// Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.
//
// k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.
//
// You may not alter the values in the list's nodes, only nodes themselves may be changed.
//
// The number of nodes in the list is n.
// 1 <= k <= n <= 5000
// 0 <= Node.val <= 1000
//
// Follow-up: Can you solve the problem in O(1) extra memory space?

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		curr.Next, prev, curr = prev, curr, curr.Next
	}

	return prev
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	for head != nil {
		tail := prev
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}

		next := tail.Next
		tail.Next = nil
		newHead := reverseList(head)
		head.Next = next
		prev.Next = newHead

		prev = head
		head = next
	}

	return dummy.Next
}

func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
	}
	return head
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	fmt.Println()
}

// func main() {
// 	list := createList([]int{1, 2, 3, 4, 5})
// 	head := list
// 	tail := head.Next.Next.Next.Next
// 	printList(list)
// 	rev := reverseList(list)
// 	printList(rev)
// 	printList(head)
// 	printList(tail)
//
// 	for tail != nil {
// 		fmt.Printf("%p\n", tail)
// 		tail = tail.Next
// 	}
// 	fmt.Printf("%p\n", head)
// }

func main() {
	// Test case 1
	list := createList([]int{1, 2, 3, 4, 5})
	k := 2
	res1 := reverseKGroup(list, k)
	fmt.Print("Want: ")
	printList(createList([]int{2, 1, 4, 3, 5}))
	fmt.Print("Got: ")
	printList(res1)

	// Test case 2
	list = createList([]int{1, 2, 3, 4, 5})
	k = 3
	res1 = reverseKGroup(list, k)
	fmt.Print("Want: ")
	printList(createList([]int{3, 2, 1, 4, 5}))
	fmt.Print("Got: ")
	printList(res1)

	// Test case 3
	list = createList([]int{1, 2, 3, 4, 5})
	k = 1
	res1 = reverseKGroup(list, k)
	fmt.Print("Want: ")
	printList(createList([]int{1, 2, 3, 4, 5}))
	fmt.Print("Got: ")
	printList(res1)
}
