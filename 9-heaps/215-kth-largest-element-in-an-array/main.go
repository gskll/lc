package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

// https://leetcode.com/problems/kth-largest-element-in-an-array/description/

// Given an integer array nums and an integer k, return the kth largest element in the array.
//
// Note that it is the kth largest element in the sorted order, not the kth distinct element.
//
// Can you solve it without sorting?
//
// 1 <= k <= nums.length <= 10^5
// -10^4 <= nums[i] <= 10^4

// Two solutions
// 1. Use max heap of size k -> O(nlogk)/O(k)
// 2. Use quickselect -> O(n) / O(1)

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	return quickselect(nums, 0, n-1, n-k)
}

func quickselect(nums []int, lo, hi, k int) int {
	for lo < hi {
		pi := partition(nums, lo, hi)
		if k <= pi {
			hi = pi
		} else {
			lo = pi + 1
		}
	}
	return nums[lo]
}

func partition(nums []int, lo, hi int) int {
	pi := lo + rand.Intn(hi-lo+1)
	nums[pi], nums[lo] = nums[lo], nums[pi]

	pivot := nums[lo]
	i, j := lo-1, hi+1

	for {
		for {
			i++
			if nums[i] >= pivot {
				break
			}
		}

		for {
			j--
			if nums[j] <= pivot {
				break
			}
		}

		if i >= j {
			return j
		}

		nums[i], nums[j] = nums[j], nums[i]
	}
}

func minHeapifyDown(heap []int, i int) {
	n := len(heap)
	for i < n {
		smallest := i
		left := 2*i + 1
		right := left + 1

		if left < n && heap[left] < heap[smallest] {
			smallest = left
		}

		if right < n && heap[right] < heap[smallest] {
			smallest = right
		}

		if smallest == i {
			break
		}

		heap[smallest], heap[i] = heap[i], heap[smallest]
		i = smallest
	}
}

func findKthLargestHeap(nums []int, k int) int {
	minHeap := nums[:k]
	for i := len(minHeap) - 1; i >= 0; i-- {
		minHeapifyDown(minHeap, i)
	}

	for _, n := range nums[k:] {
		if n > minHeap[0] {
			minHeap[0] = n
			minHeapifyDown(minHeap, 0)
		}
	}

	return minHeap[0]
}

func main() {
	tests := []struct {
		nums string
		k    int
		out  int
	}{
		{
			nums: "[3,2,1,5,6,4]",
			k:    2,
			out:  5,
		},
		{
			nums: "[3,2,3,1,2,4,5,5,6]",
			k:    4,
			out:  4,
		},
	}

	for _, t := range tests {
		var nums []int
		json.Unmarshal([]byte(t.nums), &nums)
		res := findKthLargest(nums, t.k)
		fmt.Println("exp:", t.out, "got:", res)
	}
}
