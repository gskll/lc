package main

import "fmt"

type MinHeap struct {
	items []int
}

func BuildMinHeap(items []int) *MinHeap {
	h := &MinHeap{items: items}

	for i := len(items)/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}

	return h
}

func (h *MinHeap) Insert(value int) {
	h.items = append(h.items, value)
	h.heapifyUp(len(h.items) - 1)
}

func (h *MinHeap) Pop() (int, bool) {
	if len(h.items) == 0 {
		return 0, false
	}

	top := h.items[0]
	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]

	if len(h.items) > 0 {
		h.heapifyDown(0)
	}

	return top, true
}

func (h *MinHeap) heapifyUp(i int) {
	// bubble up if items[i] < items[parent]
	if i == 0 {
		return
	}

	parent := parent(i)
	if h.items[i] < h.items[parent] {
		h.items[i], h.items[parent] = h.items[parent], h.items[i]
		h.heapifyUp(parent)
	}
}

func (h *MinHeap) heapifyDown(i int) {
	// bubble down if items[i] > smallest child
	smallest := i
	rightChild := rightChild(i)
	leftChild := leftChild(i)

	if rightChild < len(h.items) && h.items[smallest] > h.items[rightChild] {
		smallest = rightChild
	}

	if leftChild < len(h.items) && h.items[smallest] > h.items[leftChild] {
		smallest = leftChild
	}

	if smallest != i {
		h.items[i], h.items[smallest] = h.items[smallest], h.items[i]
		h.heapifyDown(smallest)
	}
}

func parent(i int) int {
	return i/2 - 1
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

func main() {
	// Example usage
	arr := []int{4, 10, 3, 5, 1, 8, 2, 7, 6, 9}
	fmt.Printf("Original array: %v\n", arr)

	heap := BuildMinHeap(arr)
	fmt.Printf("Heap array: %v\n", heap.items)

	// Extract elements in order to verify heap property
	fmt.Print("Extracting elements in order: ")
	for len(heap.items) > 0 {
		top, _ := heap.Pop()
		fmt.Printf("%d ", top)
	}
	fmt.Println()
}
