package main

import "fmt"

// ## Heap Sort
// * Time Complexity: O(n log n) in all cases
// * Space Complexity: O(1)
// * Advantages:
//     * In-place sorting
//     * Guaranteed O(n log n) performance
//     * Excellent for finding k largest/smallest elements
//     * No extra space needed unlike merge sort
// * Disadvantages:
//     * Unstable sort
//     * Slower in practice than quicksort
//     * Poor cache performance due to jumping around in memory
//     * Complex implementation compared to other algorithms

// for k largest elements we can stop after k iteration of extracting
// for k smallest elements we use a min heap
// but in practice usually better to maintain a minheap of size k for k largest elements
func heapsort(arr []int) {
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n, i int) {
	for {
		largest := i
		left := i*2 + 1
		right := left + 1

		if left < n && arr[largest] < arr[left] {
			largest = left
		}

		if right < n && arr[largest] < arr[right] {
			largest = right
		}

		if largest == i {
			break
		}

		arr[i], arr[largest] = arr[largest], arr[i]
		i = largest
	}
}

func main() {
	arr := []int{3, 6, 4, 1, 0, 8, 5, 2, 9, 7}
	fmt.Println(arr)
	heapsort(arr)
	fmt.Println(arr)
}
