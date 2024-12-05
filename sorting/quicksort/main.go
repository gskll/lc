package main

import (
	"fmt"
	"math/rand"
)

// ## Quick Sort
// * Time Complexity: Average O(n log n), Worst O(n²)
// * Space Complexity: O(log n) average case for recursion stack
// * Advantages:
//     * Best practical performance in average case
//     * In-place sorting (low space overhead)
//     * Cache-friendly due to good locality of reference
//     * Quickselect variant useful for finding kth smallest element
// * Disadvantages:
//     * Unstable sort (doesn't preserve relative order of equal elements)
//     * Poor worst-case performance O(n²)
//     * Performance depends heavily on pivot selection
// - can be O(n^2) for several reasons:
// - sorted arrays, all elements are duplicates, pattern based arrays (repeating sequences), intentionally malicious input if someone knows your pivot

// - key is how to choose the pivot - lots of ways
// - lomuto - basic - guarantees pivot is in correct place after partition
// - hoare partitioning - generally good performance - 3x less swaps than lomuto
// - 3 way partitioning - if lots of duplicates
//
// go Slice.Sort uses pattern defeating quicksort
// - general idea is
// - check if array is nearly sorted - if so use insertion sort
// - check if lots of duplicates - if so use 3-way partitioning quicksort
// - check if partitions are failing (track recursion depth)- if so use heap sort
// - if none of these then it's optimized quick sort using block partition

// look for
// - itemFromLeft : first item from left that is larger than pivot
// - itemFromRight: first item from right that is smaller than pivot
// - swap them
// - stop when they cross
// - pivotPos is itemFromLeft

// generally for hoare and 3-way we choose the first element as the pivot, for lomuto the last
// this leads to O(n^2) if sorted arrays or all same elements
// we can reduce risk of bad partitioning by adding random pivot selection
// - the chance of consistently choosing a bad pivot becomes small
// - expected time complexity becomes O(nlogn) for any input
// - makes the algorithm resistant to malicious outputs

func partitionHoare(arr []int, low, high int) int {
	pivot := arr[low]
	left := low - 1
	right := high + 1

	for {
		for {
			left++
			if arr[left] >= pivot {
				break
			}
		}

		for {
			right--
			if arr[right] <= pivot {
				break
			}
		}

		if left >= right {
			return right
		}

		arr[left], arr[right] = arr[right], arr[left]
	}
}

func quicksortHoare(arr []int, low, high int) {
	if low < high {
		pivotPos := partitionHoare(arr, low, high)
		quicksort(arr, low, pivotPos)
		quicksort(arr, pivotPos+1, high)
	}
}

// good with lots of duplicates - we have a 3rd partition with data that's the same as the pivot - which we don't need to then include in the next iteration
func partition3Way(arr []int, low, high int) (int, int) {
	pivot := arr[low]
	lt := low
	i := low + 1
	gt := high

	for i <= gt {
		if arr[i] < pivot {
			arr[lt], arr[i] = arr[i], arr[lt]
			lt++
			i++
		} else if arr[i] > pivot {
			arr[gt], arr[i] = arr[i], arr[gt]
			gt--
		} else {
			i++
		}
	}
	return lt, gt
}

func quicksort3Way(arr []int, low, high int) {
	if low < high {
		lt, gt := partition3Way(arr, low, high)
		quicksort3Way(arr, low, lt-1)
		quicksort3Way(arr, gt+1, high)
	}
}

func partitionLomuto(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	pivotPos := i + 1
	arr[pivotPos], arr[high] = arr[high], arr[pivotPos]
	return pivotPos
}

func quicksortLomuto(arr []int, low, high int) {
	if low < high {
		pivotPos := partitionLomuto(arr, low, high)
		quicksort(arr, low, pivotPos-1)
		quicksort(arr, pivotPos+1, high)
	}
}

func partitionHoareRandom(arr []int, low, high int) int {
	pivotIdx := low + rand.Intn(high-low+1)
	arr[low], arr[pivotIdx] = arr[pivotIdx], arr[low]
	// return partitionHoare(arr, low, high)

	pivot := arr[low]
	i, j := low, high

	for {
		for arr[i] < pivot {
			i++
		}

		for arr[j] > pivot {
			j--
		}

		if i >= j {
			return j
		}

		arr[i], arr[j] = arr[j], arr[i]
	}
}

func quicksortHoareRandom(arr []int, low, high int) {
	if low < high {
		pivotPos := partitionHoareRandom(arr, low, high)
		quicksort(arr, low, pivotPos)
		quicksort(arr, pivotPos+1, high)
	}
}

// basic lomuto
func quicksort(arr []int, low, high int) {
	if low < high {
		// 1. choose pivot
		pivot := arr[high]

		// 2. partition
		i := low - 1
		for j := low; j < high; j++ {
			if arr[j] <= pivot {
				i++
				if i != j {
					arr[i], arr[j] = arr[j], arr[i]
				}
			}
		}
		pivotPos := i + 1
		arr[pivotPos], arr[high] = arr[high], arr[pivotPos]

		// 3. recursion
		quicksort(arr, low, pivotPos-1)
		quicksort(arr, pivotPos+1, high)
	}
}

func main() {
	fmt.Println("basic")
	arr := []int{3, 6, 4, 1, 0, 8, 5, 2, 9, 7}
	fmt.Println(arr)
	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)

	fmt.Println()
	fmt.Println("lomuto")
	arr2 := []int{3, 6, 4, 1, 0, 8, 5, 2, 9, 7}
	fmt.Println(arr2)
	quicksortLomuto(arr2, 0, len(arr2)-1)
	fmt.Println(arr2)

	fmt.Println()
	fmt.Println("hoare")
	arr3 := []int{3, 6, 4, 1, 0, 8, 5, 2, 9, 7}
	fmt.Println(arr3)
	quicksortHoare(arr3, 0, len(arr3)-1)
	fmt.Println(arr3)

	fmt.Println()
	fmt.Println("hoare random")
	arr4 := []int{3, 6, 4, 1, 0, 8, 5, 2, 9, 7}
	fmt.Println(arr4)
	quicksortHoareRandom(arr4, 0, len(arr4)-1)
	fmt.Println(arr4)

	fmt.Println()
	fmt.Println("3-way")
	arr5 := []int{2, 1, 1, 4, 3, 4, 3, 2, 2, 4, 4, 3}
	fmt.Println(arr5)
	quicksort3Way(arr5, 0, len(arr5)-1)
	fmt.Println(arr5)
}
