package main

import "fmt"

// ## Selection Sort
// * Time Complexity: O(n²)
// * Space Complexity: O(1)
// * Advantages:
//     * Simple implementation
//     * Minimal memory usage
//     * Makes minimum number of swaps (O(n))
// * Disadvantages:
//     * Always O(n²) even if array is sorted
//     * Unstable sort
//     * Not practical for large datasets

func selectionsort(arr []int) {
	n := len(arr)

	for i := 0; i < n; i++ {
		minIdx := i

		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	arr := []int{3, 6, 4, 1, 0, 8, 5, 2, 9, 7}
	fmt.Println(arr)
	selectionsort(arr)
	fmt.Println(arr)
}
