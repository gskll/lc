package main

import (
	"fmt"
)

// https://leetcode.com/problems/add-two-numbers/description/

// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
// The number of nodes in each linked list is in the range [1, 100].
// 0 <= Node.val <= 9
// It is guaranteed that the list represents a number that does not have leading zeros.

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	currSum := dummy

	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		currSum.Next = &ListNode{Val: sum % 10}
		carry = sum / 10

		currSum = currSum.Next
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
	list1 := createList([]int{2, 4, 3})
	list2 := createList([]int{5, 6, 4})
	printList("list1", list1)
	printList("list2", list2)
	sum1 := addTwoNumbers(list1, list2)
	printList("list 1 res", sum1)
	printList("list 1 exp", createList([]int{7, 0, 8}))

	// Test case 2
	list1 = createList([]int{0})
	list2 = createList([]int{0})
	printList("list1", list1)
	printList("list2", list2)
	sum1 = addTwoNumbers(list1, list2)
	printList("list 1 res", sum1)
	printList("list 1 exp", createList([]int{0}))

	// Test case 3
	list1 = createList([]int{9, 9, 9, 9, 9, 9, 9})
	list2 = createList([]int{9, 9, 9, 9})
	printList("list1", list1)
	printList("list2", list2)
	sum1 = addTwoNumbers(list1, list2)
	printList("list 1 res", sum1)
	printList("list 1 exp", createList([]int{8, 9, 9, 9, 0, 0, 0, 1}))
}
