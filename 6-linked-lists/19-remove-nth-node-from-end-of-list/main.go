package main

import (
	"fmt"
)

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/

// Given the head of a linked list, remove the nth node from the end of the list and return its head.
//
// The number of nodes in the list is sz.
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}

	left, right := dummy, dummy
	for i := 0; i < n; i++ {
		right = right.Next
	}

	for right.Next != nil {
		left = left.Next
		right = right.Next
	}

	left.Next = left.Next.Next

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

func printList(x string, head *ListNode) {
	current := head
	fmt.Printf("%s: ", x)
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1
	list1 := createList([]int{1, 2, 3, 4, 5})
	printList("list1", list1)
	l1 := removeNthFromEnd(list1, 2)
	printList("list 1 res", l1)
	printList("list 1 exp", createList([]int{1, 2, 3, 5}))

	// Test case 2
	list2 := createList([]int{1})
	printList("list2", list2)
	l2 := removeNthFromEnd(list2, 1)
	printList("list2 res", l2)
	printList("list2 exp", createList([]int{}))

	// Test case 3
	list3 := createList([]int{1, 2})
	printList("list3", list3)
	l3 := removeNthFromEnd(list3, 1)
	printList("list3 res", l3)
	printList("list3 exp", createList([]int{1}))

	// Test case 4
	list4 := createList([]int{1, 2})
	printList("list4", list4)
	l4 := removeNthFromEnd(list4, 2)
	printList("list4 res", l4)
	printList("list4 exp", createList([]int{2}))
}
