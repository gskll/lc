package main

// https://leetcode.com/problems/two-sum/description/

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// You can return the answer in any order.
//
// 2 <= nums.length <= 10^4
// -10^9 <= nums[i] <= 10^9
// -10^9 <= target <= 10^9
// Only one valid answer exists.

func twoSum(nums []int, target int) []int {
	complements := make(map[int]int)
	for i := range nums {
		if j, exists := complements[target-nums[i]]; exists {
			return []int{i, j}
		}
		complements[nums[i]] = i
	}
	return nil
}
