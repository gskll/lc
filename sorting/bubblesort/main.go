package main

import "fmt"

// Time Complexity: O(n²)
// Space Complexity: O(1)
// Advantages:
//
// Very simple to implement
// Stable sort
// Adaptive (fast for nearly sorted arrays)
// In-place algorithm
//
//
// Disadvantages:
//
// Very inefficient for large datasets
// Rarely used in practice
// Primarily used for educational purposes

func bubblesort(arr []int) {
	n := len(arr)
	swapped := false // flag to optimize for already sorted arrays

	for i := 0; i < n-1; i++ {
		swapped = false

		// after each pass the last i elements are in order
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		// if no swaps occurred in this pass - then array is sorted
		if !swapped {
			break
		}
	}
}

func main() {
	arr := []int{3, 6, 4, 1, 0, 8, 5, 2, 9, 7}
	fmt.Println(arr)
	bubblesort(arr)
	fmt.Println(arr)
}
