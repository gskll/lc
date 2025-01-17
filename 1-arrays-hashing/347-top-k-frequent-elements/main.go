package main

// https://leetcode.com/problems/top-k-frequent-elements/

// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
//
// 1 <= nums.length <= 10^5
// -10^4 <= nums[i] <= 10^4
// k is in the range [1, the number of unique elements in the array].
// It is guaranteed that the answer is unique.

func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int)
	for _, n := range nums {
		counter[n]++
	}

	frequencies := make([][]int, len(nums)+1)
	for n, freq := range counter {
		frequencies[freq] = append(frequencies[freq], n)
	}

	var res []int
	for i := len(frequencies) - 1; i >= 0 && len(res) < k; i-- {
		res = append(res, frequencies[i]...)
	}
	return res
}
