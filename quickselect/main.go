package main

import (
	"fmt"
	"math/rand"
)

// Like quicksort but we only recurse over one half
// Time complexity: O(n) on average. We handle n+n/2+n/4...=2n
// Memory complexity: O(1) (excluding initial array copy if we want to preserve input)

func kthLargest(arr []int, k int) int {
	temp := make([]int, len(arr))
	copy(temp, arr)

	// pi := quickselectHoare(temp, 0, len(temp)-1, len(temp)-k)
	pi := quickselectLomuto(temp, 0, len(temp)-1, len(temp)-k)
	return temp[pi]
}

func kthLargestElements(arr []int, k int) []int {
	temp := make([]int, len(arr))
	copy(temp, arr)

	// pi := quickselectHoare(temp, 0, len(temp)-1, len(temp)-k)
	pi := quickselectLomuto(temp, 0, len(temp)-1, len(temp)-k)
	return temp[pi:]
}

func kthSmallest(arr []int, k int) int {
	temp := make([]int, len(arr))
	copy(temp, arr)

	// pi := quickselectHoare(temp, 0, len(temp)-1, k-1)
	pi := quickselectLomuto(temp, 0, len(temp)-1, k-1)
	return temp[pi]
}

func kthSmallestElements(arr []int, k int) []int {
	temp := make([]int, len(arr))
	copy(temp, arr)

	// pi := quickselectHoare(temp, 0, len(temp)-1, k-1)
	pi := quickselectLomuto(temp, 0, len(temp)-1, k-1)
	return temp[:pi+1]
}

func quickselectLomuto(arr []int, low, high, k int) int {
	for low < high {
		pivotIdx := partitionLomuto(arr, low, high)
		if k == pivotIdx {
			return pivotIdx
		} else if k < pivotIdx {
			high = pivotIdx - 1
		} else {
			low = pivotIdx + 1
		}

	}
	return low
}

func partitionLomuto(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	pi := i + 1
	arr[pi], arr[high] = arr[high], arr[pi]
	return pi
}

func quickselectHoare(arr []int, low, high, k int) int {
	for low < high {
		pivotIdx := partitionHoare(arr, low, high)
		// Hoare's partition doesn't guarantee pivot is in right place, only that everything to left is smaller than pivot
		if k <= pivotIdx {
			high = pivotIdx
		} else {
			low = pivotIdx + 1
		}
	}
	return low
}

func partitionHoare(arr []int, low, high int) int {
	pivotIdx := low + rand.Intn(high-low+1)
	arr[low], arr[pivotIdx] = arr[pivotIdx], arr[low]

	pivot := arr[low]
	i, j := low-1, high+1

	for {
		for {
			i++
			if arr[i] >= pivot {
				break
			}
		}

		for {
			j--
			if arr[j] <= pivot {
				break
			}
		}

		if i >= j {
			return j
		}

		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	nums := []int{3, 6, 4, 1, 0, 8, 5, 2, 9, 7}
	fmt.Println(kthLargest(nums, 3), "==", 7)
	fmt.Println(kthLargestElements(nums, 3), "==", "[7,8,9]")
	fmt.Println(kthSmallest(nums, 3), "==", 2)
	fmt.Println(kthSmallestElements(nums, 3), "==", "[0,1,2]")
}
