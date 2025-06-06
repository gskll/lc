package main

// https://leetcode.com/problems/contains-duplicate/description/

// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
//
// 1 <= nums.length <= 10^5
// -10^9 <= nums[i] <= 10^9

func containsDuplicate(nums []int) bool {
	seen := make(map[int]interface{})
	for i := range nums {
		if _, exists := seen[nums[i]]; exists {
			return true
		}
		seen[nums[i]] = struct{}{}
	}
	return false
}
