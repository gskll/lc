package main

import "fmt"

// Key characteristics:
//
// Not in-place (needs extra space)
// Stable (when implemented carefully)
// Good for integers in a small range
// O(n+k) time complexity where k is range of input
// O(k) extra space for counting array
//
// Best used when:
//
// Data range is not significantly larger than number of objects
// Integers or data that can be mapped to integers
// Need a stable sort
// Example: sorting ages of people, grades, etc.
//
// Bad when:
//
// Data range is very large (like sorting phone numbers)
// Working with floating point numbers
// Memory is a constraint and range is large

// Time: O(n + k) where k is range (max-min+1)
// Space: O(k) for count array
func countsortBasic(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}

	// Find range
	largest := arr[0]
	smallest := arr[0]

	for _, n := range arr {
		if n > largest {
			largest = n
		}
		if n < smallest {
			smallest = n
		}
	}

	count := make([]int, largest-smallest+1)

	// Count occurrences
	for _, n := range arr {
		count[n-smallest]++
	}

	// Reconstruct basic array
	index := 0
	for i := 0; i < len(count); i++ {
		for j := 0; j < count[i]; j++ {
			arr[index] = i + smallest
			index++
		}
	}

	return arr
}

// Time: O(n + k)
// Space: O(n + k) for output array and count array
func countsortStable(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}

	// find range
	largest, smallest := arr[0], arr[0]
	for _, n := range arr {
		if n > largest {
			largest = n
		}

		if n < smallest {
			smallest = n
		}
	}

	// build count
	count := make([]int, largest-smallest+1)
	for _, n := range arr {
		count[n-smallest]++
	}

	// build cumulative count
	// the cumulative count tells us the LAST position where each number should go
	// for each index i, the cumulative count is how many elements are <= i
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	// build output array
	// iterate right to left - crucial because when we find equal elements, the last one in the original array goes to the last available position
	output := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		output[count[arr[i]-smallest]-1] = arr[i]
		count[arr[i]-smallest]--
	}

	return output
}

func main() {
	fmt.Println("basic")
	arr := []int{3, 6, 4, 1, 0, 8, 5, 2, 9, 7}
	fmt.Println(arr)
	sorted := countsortBasic(arr)
	fmt.Println(sorted)

	fmt.Println()
	fmt.Println("stable")
	arr2 := []int{4, 2, 2, 8, 3, 3, 1}
	fmt.Println(arr2)
	sorted2 := countsortStable(arr2)
	fmt.Println(sorted2)
}
