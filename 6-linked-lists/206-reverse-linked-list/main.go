package main

import (
	"fmt"
)

// https://leetcode.com/problems/reverse-linked-list/description/

// Given the head of a singly linked list, reverse the list, and return the reversed list.
//
// The number of nodes in the list is the range [0, 5000].
// -5000 <= Node.val <= 5000

type ListNode struct {
	Val  int
	Next *ListNode
}

// recursive
func reverseList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}

// iterative
// func reverseList(head *ListNode) *ListNode {
// 	curr := head
// 	var prev *ListNode
// 	for curr != nil {
// 		curr.Next, prev, curr = prev, curr, curr.Next
// 	}
// 	return prev
// }

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
	// Test case 1: [1,2,3,4,5]
	list1 := createList([]int{1, 2, 3, 4, 5})
	fmt.Print("Original list 1: ")
	printList(list1)
	reversed1 := reverseList(list1)
	fmt.Print("Reversed list 1: ")
	printList(reversed1)

	// Test case 2: [1,2]
	list2 := createList([]int{1, 2})
	fmt.Print("Original list 2: ")
	printList(list2)
	reversed2 := reverseList(list2)
	fmt.Print("Reversed list 2: ")
	printList(reversed2)

	// Test case 3: empty list
	list3 := createList([]int{})
	fmt.Print("Original list 3: ")
	printList(list3)
	reversed3 := reverseList(list3)
	fmt.Print("Reversed list 3: ")
	printList(reversed3)

	// Test case 4: single element
	list4 := createList([]int{1})
	fmt.Print("Original list 4: ")
	printList(list4)
	reversed4 := reverseList(list4)
	fmt.Print("Reversed list 4: ")
	printList(reversed4)
}
