package main

import (
	"fmt"
)

// https://leetcode.com/problems/reorder-list/description/

// You are given the head of a singly linked-list. The list can be represented as:
//
// L0 → L1 → … → Ln - 1 → Ln
// Reorder the list to be on the following form:
//
// L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
// You may not modify the values in the list's nodes. Only nodes themselves may be changed.
//
// The number of nodes in the list is in the range [1, 5 * 10^4].
// 1 <= Node.val <= 1000

type ListNode struct {
	Val  int
	Next *ListNode
}

// Splits a list in two halves
// For list of length n, returns first ceil(n/2) nodes and remainder
func splitList(head *ListNode) (*ListNode, *ListNode) {
	if head == nil || head.Next == nil {
		return head, nil
	}

	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	secondHalf := slow.Next
	slow.Next = nil

	return head, secondHalf
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		curr.Next, prev, curr = prev, curr, curr.Next
	}

	return prev
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	first, second := splitList(head)

	second = reverseList(second)

	for first != nil && second != nil {
		firstNext := first.Next
		secondNext := second.Next

		first.Next = second
		second.Next = firstNext

		first = firstNext
		second = secondNext
	}
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

func main() {
	// Test case 1
	list1 := createList([]int{1, 2, 3, 4})
	printList(list1)
	reorderList(list1)
	fmt.Print("Res list 1: ")
	printList(list1)
	printList(createList([]int{1, 4, 2, 3}))

	// Test case 2
	list2 := createList([]int{1, 2, 3, 4, 5})
	reorderList(list2)
	fmt.Print("Res list 2: ")
	printList(list2)
	printList(createList([]int{1, 5, 2, 4, 3}))
}
