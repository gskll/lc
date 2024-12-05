package main

import "fmt"

// ## Merge Sort
// * Time Complexity: O(n log n) in all cases
// * Space Complexity: O(n)
// * Advantages:
//     * Stable sort (preserves order of equal elements)
//     * Guaranteed O(n log n) performance
//     * Excellent for sorting linked lists
//     * Predictable performance regardless of input
//     * Parallelizes well
// * Disadvantages:
//     * Requires O(n) extra space
//     * Not in-place
//     * Overkill for small arrays
//     * Cache performance not as good as quicksort

func mergesort(arr []int) []int {
	n := len(arr)

	if n <= 1 {
		return arr
	}

	mid := n / 2
	left := mergesort(arr[:mid])
	right := mergesort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))

	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	merged = append(merged, left[i:]...)
	// for i < len(left) {
	// 	merged = append(merged, left[i])
	// 	i++
	// }

	merged = append(merged, right[j:]...)
	// for j < len(right) {
	// 	merged = append(merged, right[j])
	// 	j++
	// }

	return merged
}

func main() {
	arr := []int{3, 6, 4, 1, 0, 8, 5, 2, 9, 7}
	fmt.Println(arr)
	sorted := mergesort(arr)
	fmt.Println(sorted)
}
