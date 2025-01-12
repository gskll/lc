package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// https://leetcode.com/problems/maximum-product-subarray/description/

// Given an integer array nums, find a subarray that has the largest product, and return the product.
//
// The test cases are generated so that the answer will fit in a 32-bit integer.
//
// 1 <= nums.length <= 2 * 10^4
// -10 <= nums[i] <= 10
// The product of any subarray of nums is guaranteed to fit in a 32-bit integer.

/*
	- optimization problem - dp?
	- len(nums) medium
	- 32bit int so int type is fine

	- brute force
		- check every subarray and calculate the max

	- negative numbers and 0
	- including a 0 will set subarray product to 0
	- odd number of negative numbers will set to negative
	- we want to maximise the subarray with 1) no 0's and 2) an even number of negative numbers including no negative numbers
	- could iterate through array and create indices of 0s and negative numbers
	- no because 2nd test, only negatives and 0

	- sliding window approach?
		- increase to the right
		- reset on 0
		- if negative we need to somehow track if odd/even negatives in the subarray
		- local/global max sum

	- decision tree?
	- take/skip can't skip because contiguous
	- [2,3,-2,4]

	- prefix sums/products
	- [2,6,-12,-48]
	- [-48,-24,-8,4]
	- process forwards/backwards, max will be max value in both arrays? no doesn't account for negatives properly
	- !! works if we reset to 1 on 0

	- calculate remaining negatives
	- iterate over array, reset on 0
	- if we hit a negative, we store the current product
		- keep going until we hit another negative in which case the product becomes positive again

	- F(n) = max(F(n-1)*n, n-1*n, n)
	- F(0) = 0

	===================

	kadane's algorithm
		- insight that to for each position we can track the maxProduct so far AND the minProduct so far, and use that to calculate the maxProduct for i to account for negatives
		- F(n) = max(maxP[i]) for i from 0 to n-1
		- maxP[i] = max(nums[i], nums[i]*maxP[i-1], nums[i]*minP[i-1])
		- minP[i] = min(nums[i], nums[i]*maxP[i-1], nums[i]*minP[i-1])
		- maxP[0] = nums[0]
		- minP[0] = nums[0]

*/

// prefix/suffix products optimized - O(n) / O(1)
func maxProduct(nums []int) int {
	n := len(nums)
	prefix, suffix := 1, 1
	curMax := nums[0]

	for i := range nums {
		if prefix == 0 {
			prefix = 1
		}
		if suffix == 0 {
			suffix = 1
		}
		prefix *= nums[i]
		suffix *= nums[n-1-i]

		curMax = max(curMax, prefix, suffix)
	}
	return curMax
}

// prefix/suffix products - O(n)/O(1)
func maxProduct1(nums []int) int {
	n := len(nums)

	currMax := nums[0]
	prev := 1
	for _, n := range nums {
		if prev == 0 {
			prev = 1
		}
		curr := prev * n
		currMax = max(currMax, curr)
		prev = curr
	}

	prev = 1
	for i := n - 1; i >= 0; i-- {
		if prev == 0 {
			prev = 1
		}
		curr := prev * nums[i]
		currMax = max(currMax, curr)
		prev = curr
	}

	return currMax
}

// kadane's algorithm: O(n)/ O(1)
func maxProduct2(nums []int) int {
	res := nums[0]
	maxP, minP := 1, 1

	for _, n := range nums {
		currMax := max(n, n*maxP, n*minP)
		minP = min(n, n*maxP, n*minP)
		maxP = currMax
		res = max(maxP, res)
	}

	return res
}

func main() {
	tests := []struct {
		nums string
		out  int
	}{
		{
			nums: "[2,3,-2,4]",
			out:  6,
		},
		{
			nums: "[-2,0,-1]",
			out:  0,
		},
		{
			nums: "[-3,0,1,-2]",
			out:  1,
		},
		{
			nums: "[-1,-2,-3,0]",
			out:  6,
		},
	}

	for _, t := range tests {
		var nums []int
		json.Unmarshal([]byte(t.nums), &nums)
		start := time.Now()
		res := maxProduct(nums)
		fmt.Printf("%v == %v | %v\n", t.out, res, time.Since(start))
	}
}
