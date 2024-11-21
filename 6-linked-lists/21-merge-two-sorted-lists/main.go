package main

import (
	"fmt"
)

// https://leetcode.com/problems/merge-two-sorted-lists/description/

// You are given the heads of two sorted linked lists list1 and list2.
//
// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
//
// Return the head of the merged linked list.
//
// The number of nodes in both lists is in the range [0, 50].
// -100 <= Node.val <= 100
// Both list1 and list2 are sorted in non-decreasing order.

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	dummy := &ListNode{}
	curr := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	}

	if list2 != nil {
		curr.Next = list2
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

func main() {
	// Test case 1
	list1 := createList([]int{1, 2, 4})
	list1_2 := createList([]int{1, 3, 4})
	printList(list1)
	printList(list1_2)
	res1 := mergeTwoLists(list1, list1_2)
	fmt.Print("Res list 1: ")
	printList(res1)

	// Test case 2
	list2 := createList([]int{})
	list2_2 := createList([]int{})
	res2 := mergeTwoLists(list2, list2_2)
	fmt.Print("Res list 2: ")
	printList(res2)

	// Test case 3
	list3 := createList([]int{})
	list3_2 := createList([]int{0})
	res3 := mergeTwoLists(list3, list3_2)
	fmt.Print("Res list 3: ")
	printList(res3)
}
