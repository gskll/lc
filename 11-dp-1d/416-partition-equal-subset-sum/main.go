package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// https://leetcode.com/problems/partition-equal-subset-sum/description/
//
// Given an integer array nums, return true if you can partition the array into two subsets such that the sum of the elements in both subsets is equal or false otherwise.
//
// 1 <= nums.length <= 200
// 1 <= nums[i] <= 100

/*
	- find whether nums array can be split in two such that the elements sum equally
	- nums length is small

	- sort the array first?
	- take prefix and suffix sums for each index i
	- find an index where the prefix i and suffix i match

	[1,5,11,5]
	[0,1,6,17]
	[21,16,5,0]

	[1,2,3,5]
	[0,1,3,6]
	[10,9,6,0]

	F(i) = true if prefix[i] + suffix[i] == nums[i]

	[2,2,3,5]
	[0,2,4,7]
	[10,8,5,0]

	odd vs even length?
	pick/skip?
	subset sum = total sum / 2
	if total sum is odd not possible

	take total sum
	if odd return false

	we're looking for a subset where sum = totalsum/2

	- decision pick/skip - 2D dp
	- F(i,t) = OR(
					F(i+1,t-nums[i]) // pick if nums[i] <= t
					F(i+1,t)		 // skip
				)
	- F(i, 0) = true
	- F(n, t) = t==0

	- is there a 1D dp solution?
	- nums: "[1,5,11,5]",
	- out:  true,
	- total sum = 22, target = 11
	- F(i) = nums[i] + F(i-1) == 11
	- F(0) = nums[i]
	- still depends on the target somehow
*/

// optimized: O(n) / O(target)
// dp[i] represents can we make i with the subset seen so far
func canPartition(nums []int) bool {
	sum := 0
	for _, x := range nums {
		sum += x
	}
	if sum%2 == 1 {
		return false
	}

	sum /= 2

	dp := make([]bool, sum+1)
	dp[0] = true

	for _, n := range nums {
		if n > sum {
			continue
		}

		if dp[sum-n] {
			return true
		}

		for i := sum; i >= n; i-- {
			if dp[i-n] {
				dp[i] = true
			}
		}
	}

	return false
}

// bottom up - O(n*t)/ O(n*t)
// Bounds
// i -> [0,len(nums)]
// remaining -> [0, target]
// i depends on i+1
// remaining depends on remaining-nums[i]
//
// i,remaining requires
//   - take i+1, remaining-nums[i]
//   - i depends on i+1
//   - remaining depends on remaining-nums[i]
//   - skip i+1, remaining
//   - i depends on i+1
//   - remaining is constant
//
// i depends on i+1
// remaining depends on remaining-nums[i]
func canPartitionBottomUp(nums []int) bool {
	n := len(nums)
	total := 0
	for _, x := range nums {
		total += x
	}
	if total%2 == 1 {
		return false
	}
	target := total / 2

	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}

	for i := range dp[n] {
		dp[n][i] = false
	}

	for i := n - 1; i >= 0; i-- {
		for remaining := 1; remaining <= target; remaining++ {
			dp[i][remaining] = dp[i+1][remaining]

			if nums[i] <= remaining {
				dp[i][remaining] = dp[i][remaining] || dp[i+1][remaining-nums[i]]
			}
		}
	}

	return dp[0][target]
}

// top down - O(n*target) / O(n*target)
func canPartitionTopDown(nums []int) bool {
	n := len(nums)
	total := 0
	for _, x := range nums {
		total += x
	}
	if total%2 == 1 {
		return false
	}
	target := total / 2

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
		dp[i][0] = 1
	}

	var findSubset func(i, remaining int) bool
	findSubset = func(i, remaining int) bool {
		if dp[i][remaining] != -1 {
			return dp[i][remaining] == 1
		}
		if remaining == 0 {
			return true
		}
		if i == n {
			return false
		}

		// pick
		pick := false
		if nums[i] <= remaining {
			pick = findSubset(i+1, remaining-nums[i])
		}

		// skip
		skip := findSubset(i+1, remaining)

		if pick || skip {
			dp[i][remaining] = 1
		} else {
			dp[i][remaining] = 0
		}

		return pick || skip
	}

	return findSubset(0, target)
}

// top down with map dp
// significantly slower than array dp
func canPartitionMap(nums []int) bool {
	n := len(nums)
	total := 0
	for _, x := range nums {
		total += x
	}
	if total%2 == 1 {
		return false
	}
	target := total / 2

	dp := make(map[int]bool, n+1)

	var findSubset func(i, remaining int) bool
	findSubset = func(i, remaining int) bool {
		key := i*(target+1) + remaining

		if val, ok := dp[key]; ok {
			return val
		}

		if remaining == 0 {
			return true
		}
		if i == n {
			return false
		}

		// pick
		pick := false
		if nums[i] <= remaining {
			pick = findSubset(i+1, remaining-nums[i])
		}

		// skip
		skip := findSubset(i+1, remaining)

		dp[key] = pick || skip
		return dp[key]
	}

	return findSubset(0, target)
}

func canPartitionR(nums []int) bool {
	n := len(nums)
	total := 0
	for _, x := range nums {
		total += x
	}
	if total%2 == 1 {
		return false
	}
	target := total / 2

	var findSubset func(i, remaining int) bool
	findSubset = func(i, remaining int) bool {
		if remaining == 0 {
			return true
		}
		if i == n {
			return false
		}

		// pick
		pick := false
		if nums[i] <= remaining {
			pick = findSubset(i+1, remaining-nums[i])
		}

		// skip
		skip := findSubset(i+1, remaining)
		return pick || skip
	}

	return findSubset(0, target)
}

func main() {
	tests := []struct {
		nums string
		out  bool
	}{
		{
			nums: "[1,5,11,5]",
			out:  true,
		},
		{
			nums: "[1,2,3,5]",
			out:  false,
		},
		{
			nums: "[2,2,3,5]",
			out:  false,
		},
		{
			nums: "[3,3,3,4,5]",
			out:  true,
		},
	}

	for _, t := range tests {
		var nums []int
		json.Unmarshal([]byte(t.nums), &nums)
		start := time.Now()
		res := canPartition(nums)
		fmt.Printf("%v == %v | %v\n", t.out, res, time.Since(start))
	}
}
