package main

import "fmt"

type MaxHeap struct {
	items []int
}

func BuildMaxHeap(items []int) *MaxHeap {
	h := &MaxHeap{items: items}

	for i := len(items)/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}

	return h
}

func (h *MaxHeap) Insert(value int) {
	h.items = append(h.items, value)
	h.heapifyUp(len(h.items) - 1)
}

func (h *MaxHeap) Pop() (int, bool) {
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

func (h *MaxHeap) heapifyDown(i int) {
	// swap if items[i] < largest child
	largest := i
	leftChild := h.leftChild(i)
	rightChild := h.rightChild(i)

	if leftChild < len(h.items) && h.items[leftChild] > h.items[largest] {
		largest = leftChild
	}

	if rightChild < len(h.items) && h.items[rightChild] > h.items[largest] {
		largest = rightChild
	}

	if largest != i {
		h.items[largest], h.items[i] = h.items[i], h.items[largest]
		h.heapifyDown(largest)
	}
}

func (h *MaxHeap) heapifyUp(i int) {
	// if items[i] > items[parent] swap
	if i == 0 {
		return
	}

	parent := h.parent(i)

	if h.items[i] > h.items[parent] {
		h.items[i], h.items[parent] = h.items[parent], h.items[i]
		h.heapifyUp(parent)
	}
}

func (h *MaxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *MaxHeap) leftChild(i int) int {
	return 2*i + 1
}

func (h *MaxHeap) rightChild(i int) int {
	return 2*i + 2
}

func main() {
	// Example usage
	arr := []int{4, 10, 3, 5, 1, 8, 2, 7, 6, 9}
	fmt.Printf("Original array: %v\n", arr)

	heap := BuildMaxHeap(arr)
	fmt.Printf("Heap array: %v\n", heap.items)

	// Extract elements in order to verify heap property
	fmt.Print("Extracting elements in order: ")
	for len(heap.items) > 0 {
		top, _ := heap.Pop()
		fmt.Printf("%d ", top)
	}
	fmt.Println()
}
