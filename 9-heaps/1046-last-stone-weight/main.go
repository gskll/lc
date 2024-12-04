package main

import (
	"fmt"
)

// https://leetcode.com/problems/last-stone-weight/description/
// You are given an array of integers stones where stones[i] is the weight of the ith stone.
//
// We are playing a game with the stones. On each turn, we choose the heaviest two stones and smash them together. Suppose the heaviest two stones have weights x and y with x <= y. The result of this smash is:
//
// If x == y, both stones are destroyed, and
// If x != y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.
// At the end of the game, there is at most one stone left.
//
// Return the weight of the last remaining stone. If there are no stones left, return 0.
//
// 1 <= stones.length <= 30
// 1 <= stones[i] <= 1000

type MaxHeap []int

func (h *MaxHeap) Pop() int {
	old := *h
	n := len(old)

	if n == 0 {
		return 0
	}

	largest := old[0]
	old[0] = old[n-1]
	*h = old[:n-1]
	if len(*h) > 0 {
		h.heapifyDown(0)
	}
	return largest
}

func (h *MaxHeap) Push(val int) {
	*h = append(*h, val)
	h.heapifyUp(len(*h) - 1)
}

// works because slice value points to underlying array
// need pointer receivers if modifying underlying array
// not idiomatic go - same receiver type for a given type
func (h MaxHeap) heapifyUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2

		if h[parent] >= h[i] {
			break
		}

		h[i], h[parent] = h[parent], h[i]
		i = parent
	}
}

func (h MaxHeap) heapifyDown(i int) {
	n := len(h)
	for {
		largest := i
		left := 2*i + 1
		right := 2*i + 2

		if left < n && h[left] > h[largest] {
			largest = left
		}

		if right < n && h[right] > h[largest] {
			largest = right
		}

		if largest == i {
			break
		}

		h[i], h[largest] = h[largest], h[i]
		i = largest
	}
}

func lastStoneWeight(stones []int) int {
	h := MaxHeap(stones)

	for i := len(h)/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}

	for len(h) > 1 {
		x := h.Pop()
		y := h.Pop()

		if diff := x - y; diff > 0 {
			h.Push(diff)
		}
	}

	if len(h) == 0 {
		return 0
	}

	return h[0]
}

func main() {
	tests := []struct {
		stones []int
		out    int
	}{
		{
			stones: []int{2, 7, 4, 1, 8, 1},
			out:    1,
		},
		{
			stones: []int{1},
			out:    1,
		},
		{
			stones: []int{10, 5, 4, 10, 3, 1, 7, 8},
			out:    0,
		},
	}

	for _, t := range tests {
		res := lastStoneWeight(t.stones)
		fmt.Println("exp:", t.out, "got:", res)
	}
}
