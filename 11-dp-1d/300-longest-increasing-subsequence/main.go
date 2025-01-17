package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// https://leetcode.com/problems/longest-increasing-subsequence/description/

// Given an integer array nums, return the length of the longest strictly increasing subsequence
// A subsequence is an array that can be derived from another array by deleting some or no elements without changing the order of the remaining elements.
//
//
// 1 <= nums.length <= 2500
// -10^4 <= nums[i] <= 10^4

/*
	- nums length is medium, num is medium
	- if array is len 1, LIS is array[0]
	- if array is len 2, LIS is 2 if array[0] < array[1], else 1


	nums: "[0,1,0,3,2,3]",
	- [] 0
	- [3] 1
	- [2,3] 2
	- [_,2,3] 2
	- [3,2,3] 2
	- [0,_,2,3] 3
	- [1,_,_,2,3] 3
	- [0,1,_,_,2,3] 4

	- F(0) = 1
	- F(i) = 1 + max(F(j)) for 0 <=j < i && nums[i] > nums[j]

	- [0] -> 1
	- [0,1] -> 2
	- [0,1,0] -> 1
	- [0,1,0,3] -> 3
		- F(3) = 1 + max(F(j)) for j=0, 1,2 if nums[j] < 3
		- need to see if there's a sequence that 3 can extend
		- [ 0, 1, 0]
		- [ 1,2,1]
		- 3 can extend S2 - 2
		- 5 can extend S1 - 3
		- 3 can extend S0 - 2
		- S3 = 3

	- for each index, we can either include or skip
	- gives us two variables to track in dp as opposed to one
	- F(i, prevNum) = max(1+F(i+1, nums[i]) // include nums[i] if nums[i] > prevNum
							F(i+1, prevNum)) // exclude nums[i]
*/

/*
	- Extra- O(n*logn) solution with binary search
	- Key Ideas:
		- Maintain array dp[] where dp[i] = smallest number that can end a subsequence of length i+1 dp[] is always kept sorted
		- For each number, use binary search to find its insertion position in dp[]
		- If insert at end -> found longer subsequence
		- If insert in middle -> found better ending for existing length
		-
	- Why it works:
		-
		- Smaller endings are always better - they give more room for future numbers
		- Length of dp[] equals length of LIS, even though dp values may not form actual subsequence
		- Binary search gives O(log n) insertion, total O(n log n)
*/

func lengthOfLIS(nums []int) int {
	dp := make([]int, 0, len(nums))

	for _, n := range nums {
		dpI := binarySearch(dp, n)
		if dpI == len(dp) {
			dp = append(dp, n)
		} else {
			dp[dpI] = n
		}
	}

	return len(dp)
}

func binarySearch(dp []int, target int) int {
	l, r := 0, len(dp)
	for l < r {
		mid := l + (r-l)/2

		if dp[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}

// O(n^2) / O(n)
func lengthOfLISBottomUp(nums []int) int {
	n := len(nums)

	dp := make([]int, n)
	dp[0] = 1

	maxLIS := 0
	for i := 0; i < n; i++ {
		lisI1 := 0
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				lisI1 = max(lisI1, dp[j])
			}
		}

		dp[i] = lisI1 + 1
		maxLIS = max(maxLIS, dp[i])
	}

	return maxLIS
}

// O(n^2)/ O(n)
func lengthOfLISTopDown(nums []int) int {
	n := len(nums)

	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[n] = 0

	var count func(i int) (lis int)
	count = func(i int) (lis int) {
		if dp[i] > 0 {
			return dp[i]
		}
		lisI1 := 0
		for j := i + 1; j < n; j++ {
			if nums[i] < nums[j] {
				lisI1 = max(lisI1, count(j))
			}
		}

		dp[i] = lisI1 + 1
		return dp[i]
	}

	maxLIS := 0
	for i := 0; i < n; i++ {
		maxLIS = max(maxLIS, count(i))
	}
	return maxLIS
}

func main() {
	tests := []struct {
		nums string
		out  int
	}{
		{
			nums: "[10,9,2,5,3,7,101,18]",
			out:  4,
		},
		{
			nums: "[0,1,0,3,2,3]",
			out:  4,
		},
		{
			nums: "[7,7,7,7,7,7,7]",
			out:  1,
		},
	}

	for _, t := range tests {
		var nums []int
		json.Unmarshal([]byte(t.nums), &nums)
		start := time.Now()
		res := lengthOfLIS(nums)
		fmt.Printf("%v == %v | %v\n", t.out, res, time.Since(start))
	}
}
