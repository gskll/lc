package main

// https://leetcode.com/problems/product-of-array-except-self/

// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
//
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
//
// You must write an algorithm that runs in O(n) time and without using the division operation.
//
// 2 <= nums.length <= 10^5
// -30 <= nums[i] <= 30
// The input is generated such that answer[i] is guaranteed to fit in a 32-bit integer.

func productExceptSelf(nums []int) []int {
	products := make([]int, len(nums))

	for i := range nums {
		if i == 0 {
			products[i] = 1
			continue
		}
		products[i] = nums[i-1] * products[i-1]
	}

	suf := 1
	for i := len(nums) - 1; i >= 0; i-- {
		products[i] = suf * products[i]
		suf *= nums[i]
	}

	return products
}
