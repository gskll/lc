package main

import (
	"fmt"
)

// https://leetcode.com/problems/linked-list-cycle/description/

// Given head, the head of a linked list, determine if the linked list has a cycle in it.
//
// There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.
//
// Return true if there is a cycle in the linked list. Otherwise, return false.
//
// The number of the nodes in the list is in the range [0, 104].
// -105 <= Node.val <= 105
// pos is -1 or a valid index in the linked-list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}

func createList(values []int, tailNext int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
	}
	if tailNext >= 0 {
		tail := current
		current = head
		for i := 0; i < tailNext; i++ {
			current = current.Next
		}
		tail.Next = current
	}
	return head
}

func printList(x string, head *ListNode, n int) {
	current := head
	fmt.Printf("%s: ", x)
	for i := 0; i < n+1; i++ {
		if current == nil {
			break
		}
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	fmt.Println()
}

func main() {
	// Test case 1
	list1 := createList([]int{3, 2, 0, -4}, 1)
	printList("list1", list1, 4)
	res := hasCycle(list1)
	fmt.Printf("got: %t, expected: %t\n", res, true)

	// Test case 2
	list1 = createList([]int{1, 2}, 0)
	printList("list1", list1, 2)
	res = hasCycle(list1)
	fmt.Printf("got: %t, expected: %t\n", res, true)

	// Test case 3
	list1 = createList([]int{1}, -1)
	printList("list1", list1, 1)
	res = hasCycle(list1)
	fmt.Printf("got: %t, expected: %t\n", res, false)
}
