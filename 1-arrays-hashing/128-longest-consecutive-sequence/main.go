package main

// https://leetcode.com/problems/longest-consecutive-sequence/

// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
//
// You must write an algorithm that runs in O(n) time.
//
// 0 <= nums.length <= 10^5
// -10^9 <= nums[i] <= 10^9

func longestConsecutive(nums []int) int {
	numMap := make(map[int]bool, len(nums))

	for _, x := range nums {
		numMap[x] = true
	}

	var longest int
	for _, x := range nums {
		if numMap[x-1] {
			continue
		}

		curr := 1
		for numMap[x+1] {
			curr++
			x++
		}

		longest = max(curr, longest)
	}

	return longest
}
