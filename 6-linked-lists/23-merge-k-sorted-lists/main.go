package main

import (
	"container/heap"
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

// Solution O(nlogK), O(k) memory where k is number of lists and n is total number of nodes
type ListNode struct {
	Val  int
	Next *ListNode
}

type MinHeap struct {
	items []*ListNode
}

func (h *MinHeap) Len() int {
	return len(h.items)
}

func (h *MinHeap) Less(i, j int) bool {
	return h.items[i].Val < h.items[j].Val
}

func (h *MinHeap) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *MinHeap) Push(x any) {
	h.items = append(h.items, x.(*ListNode))
}

func (h *MinHeap) Pop() any {
	v := h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]

	return v
}

func (h *MinHeap) PushValue(n *ListNode) {
	heap.Push(h, n)
}

func (h *MinHeap) PopValue() *ListNode {
	return heap.Pop(h).(*ListNode)
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &MinHeap{}
	heap.Init(h)
	// add each linked list to heap and then the next as it leaves
	dummy := &ListNode{}
	sorted := dummy
	for _, list := range lists {
		if list != nil {
			h.PushValue(list)
		}
	}

	for h.Len() > 0 {
		minNode := h.PopValue()
		if minNode.Next != nil {
			h.PushValue(minNode.Next)
			minNode.Next = nil
		}
		sorted.Next = minNode
		sorted = sorted.Next
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
}
