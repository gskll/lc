package main

import (
	"fmt"
)

// https://leetcode.com/problems/merge-k-sorted-lists/description/

// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
//
// Merge all the linked-lists into one sorted linked-list and return it.
//
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] is sorted in ascending order.
// The sum of lists[i].length will not exceed 10^4.

// Solution O(n*k), O(1) memory where k is number of lists and n is total number of nodes
type ListNode struct {
	Val  int
	Next *ListNode
}

func merge2Lists(list1, list2 *ListNode) *ListNode {
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

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	sorted := lists[0]

	for _, list := range lists[1:] {
		sorted = merge2Lists(sorted, list)
	}

	return sorted
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
	in := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}
	lists := []*ListNode{}
	for _, l := range in {
		list := createList(l)
		lists = append(lists, list)
		printList(list)
	}
	res1 := mergeKLists(lists)
	fmt.Print("Want: ")
	printList(createList([]int{1, 1, 2, 3, 4, 4, 5, 6}))
	fmt.Print("Got: ")
	printList(res1)

	// Test case 2
	in = [][]int{}
	lists = []*ListNode{}
	for _, l := range in {
		list := createList(l)
		lists = append(lists, list)
		printList(list)
	}
	res1 = mergeKLists(lists)
	fmt.Print("Want: ")
	printList(createList([]int{}))
	fmt.Print("Got: ")
	printList(res1)

	// Test case 3
	in = [][]int{{}}
	lists = []*ListNode{}
	for _, l := range in {
		list := createList(l)
		lists = append(lists, list)
		printList(list)
	}
	res1 = mergeKLists(lists)
	fmt.Print("Want: ")
	printList(createList([]int{}))
	fmt.Print("Got: ")
	printList(res1)

	// Test case 4
	in = [][]int{{}, {1}}
	lists = []*ListNode{}
	for _, l := range in {
		list := createList(l)
		lists = append(lists, list)
		printList(list)
	}
	res1 = mergeKLists(lists)
	fmt.Print("Want: ")
	printList(createList([]int{1}))
	fmt.Print("Got: ")
	printList(res1)
}
