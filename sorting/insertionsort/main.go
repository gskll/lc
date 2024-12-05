package main

import "fmt"

// Time Complexity: O(n²)
// Space Complexity: O(1)
// Advantages:
//
// Excellent for small arrays (< 50 elements)
// Adaptive (very fast for nearly sorted data)
// Stable sort
// In-place algorithm
// Simple implementation
//
//
// Disadvantages:
//
// Quadratic time complexity makes it inefficient for large datasets
// Not suitable for large-scale sorting

func insertionsort(arr []int) {
	for i := 1; i < len(arr); i++ {
		curr := arr[i]

		j := i - 1

		for j >= 0 && arr[j] > curr {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = curr
	}
}

func main() {
	arr := []int{3, 6, 4, 1, 0, 8, 5, 2, 9, 7}
	fmt.Println(arr)
	insertionsort(arr)
	fmt.Println(arr)
}
