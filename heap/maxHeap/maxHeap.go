package main

import "fmt"

type HeapItem struct {
	value int
	index int
}

type MaxHeap struct {
	items []*HeapItem
}

func BuildMaxHeap(ints []int) *MaxHeap {
	items := make([]*HeapItem, len(ints))
	for i := range ints {
		items[i] = &HeapItem{value: ints[i], index: i}
	}
	h := &MaxHeap{items: items}

	for i := len(items)/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}

	return h
}

func (h *MaxHeap) Insert(value int) *HeapItem {
	item := &HeapItem{value: value, index: len(h.items)}
	h.items = append(h.items, item)
	h.heapifyUp(item.index)
	return item
}

func (h *MaxHeap) Pop() (*HeapItem, bool) {
	if len(h.items) == 0 {
		return nil, false
	}

	top := h.items[0]
	top.index = -1

	lastIdx := len(h.items) - 1

	if lastIdx > 0 {
		h.items[0] = h.items[lastIdx]
		h.items[0].index = 0
		h.items = h.items[:lastIdx]
		h.heapifyDown(0)
	} else {
		h.items = h.items[:0]
	}

	return top, true
}

func (h *MaxHeap) Update(item *HeapItem, newValue int) {
	if item.index == -1 {
		panic("updating removed item")
	}

	prevValue := item.value
	item.value = newValue

	if prevValue < newValue {
		h.heapifyUp(item.index)
	} else {
		h.heapifyDown(item.index)
	}
}

func (h *MaxHeap) heapifyDown(i int) {
	n := len(h.items)

	for {
		largest := i
		leftChild := h.leftChild(i)
		rightChild := h.rightChild(i)

		if leftChild < n && h.items[largest].value < h.items[leftChild].value {
			largest = leftChild
		}

		if rightChild < n && h.items[largest].value < h.items[rightChild].value {
			largest = rightChild
		}

		if largest == i {
			break
		}

		h.items[i], h.items[largest] = h.items[largest], h.items[i]
		h.items[i].index = i
		h.items[largest].index = largest
		i = largest
	}
}

func (h *MaxHeap) heapifyUp(i int) {
	for i > 0 {
		parent := h.parent(i)

		if h.items[parent].value >= h.items[i].value {
			break
		}

		h.items[i], h.items[parent] = h.items[parent], h.items[i]
		h.items[i].index = i
		h.items[parent].index = parent
		i = parent
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

// Helper function to verify max heap property
func (h *MaxHeap) verifyHeapProperty() bool {
	for i := 0; i < len(h.items); i++ {
		left := h.leftChild(i)
		right := h.rightChild(i)

		// Verify index tracking
		if h.items[i].index != i {
			fmt.Printf("Index mismatch at position %d: item.index=%d\n",
				i, h.items[i].index)
			return false
		}

		// Check left child
		if left < len(h.items) {
			if h.items[i].value < h.items[left].value {
				fmt.Printf("Heap property violated at index %d (value=%d) with left child (value=%d)\n",
					i, h.items[i].value, h.items[left].value)
				return false
			}
		}

		// Check right child
		if right < len(h.items) {
			if h.items[i].value < h.items[right].value {
				fmt.Printf("Heap property violated at index %d (value=%d) with right child (value=%d)\n",
					i, h.items[i].value, h.items[right].value)
				return false
			}
		}
	}
	return true
}

// Helper to print heap state
func (h *MaxHeap) printState(label string) {
	fmt.Printf("\n=== %s ===\n", label)
	fmt.Print("Heap: [")
	for i, item := range h.items {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%d(i:%d)", item.value, item.index)
	}
	fmt.Println("]")

	if !h.verifyHeapProperty() {
		fmt.Println("WARNING: Heap property violated!")
	}
}

func main() {
	// Test 1: Building heap from array
	fmt.Println("\n=== Test 1: Building Heap ===")
	arr := []int{4, 10, 3, 5, 1, 8, 2, 7, 6, 9}
	fmt.Printf("Original array: %v\n", arr)
	heap := BuildMaxHeap(arr)
	heap.printState("After BuildMaxHeap")

	// Test 2: Inserting new elements
	fmt.Println("\n=== Test 2: Insertion ===")
	items := make(map[int]*HeapItem) // Keep track of items for updates
	newValues := []int{15, 12, 4}
	for _, val := range newValues {
		fmt.Printf("\nInserting %d", val)
		items[val] = heap.Insert(val)
		heap.printState("After insertion")
	}

	// Test 3: Updating values
	fmt.Println("\n=== Test 3: Updates ===")
	// Test increasing value
	item := items[4]
	fmt.Printf("\nUpdating value %d to 20\n", item.value)
	heap.Update(item, 20)
	heap.printState("After increasing value")

	// Test decreasing value
	item = items[15]
	fmt.Printf("\nUpdating value %d to 1\n", item.value)
	heap.Update(item, 1)
	heap.printState("After decreasing value")

	// Test 4: Popping elements
	fmt.Println("\n=== Test 4: Popping Elements ===")
	fmt.Print("Values in order: ")
	values := make([]int, 0)
	for heap.items != nil && len(heap.items) > 0 {
		item, _ := heap.Pop()
		values = append(values, item.value)
		fmt.Printf("%d ", item.value)

		// Verify heap property after each pop
		if len(heap.items) > 0 {
			heap.printState("After pop")
		}
	}
	fmt.Println()

	// Verify values are in descending order
	for i := 1; i < len(values); i++ {
		if values[i-1] < values[i] {
			fmt.Printf("ERROR: Values not in descending order: %d before %d\n",
				values[i-1], values[i])
		}
	}

	// Test 5: Edge cases
	fmt.Println("\n=== Test 5: Edge Cases ===")

	// Empty heap pop
	fmt.Println("\nTesting pop from empty heap:")
	if _, ok := heap.Pop(); ok {
		fmt.Println("ERROR: Pop from empty heap returned true")
	} else {
		fmt.Println("Successfully handled empty heap pop")
	}

	// Single element
	fmt.Println("\nTesting single element:")
	heap.Insert(5)
	heap.printState("Single element")

	// Pop single element
	heap.Pop()
	if len(heap.items) != 0 {
		fmt.Println("ERROR: Heap not empty after popping single element")
	} else {
		fmt.Println("Successfully handled single element pop")
	}

	// Test 6: Update error handling
	fmt.Println("\n=== Test 6: Update Error Handling ===")
	item = &HeapItem{value: 100, index: -1} // Simulated removed item
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Successfully caught panic on updating removed item")
			}
		}()
		heap.Update(item, 200)
		fmt.Println("ERROR: Failed to panic on updating removed item")
	}()
}
