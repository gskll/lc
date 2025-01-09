package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// https://leetcode.com/problems/min-cost-climbing-stairs/description/

// You are given an integer array cost where cost[i] is the cost of ith step on a staircase. Once you pay the cost, you can either climb one or two steps.
//
// You can either start from the step with index 0, or the step with index 1.
//
// Return the minimum cost to reach the top of the floor
//
// 2 <= cost.length <= 1000
// 0 <= cost[i] <= 999

/*
	- Using an array of costs find the lowest cost to iterate to the end of the array starting at 0 or 1 and taking 1 or 2 steps in each iteration
	- finding the minimum. need to find the minimum starting at 0, and the minimum starting at 1, then the minimum between them
	- since we can take 1 step, we'll calculate the minimum starting at 1 by removing the cost[0] from the minimum path starting at 0 then taking 1 step
	- len(costs) is medium
	- costs aren't unique
	- costs are positive or 0 - don't have to worry about negatives
	- taking a step will never decrease the total cost

	cost [0,999]
	- can go either direction, from i = [0,len(costs)-1] or reverse
	- either stop when i >= len(costs) || i < 0
	- top-down or bottom-up
	- probably easier to go forwards to handle two start points

	- Decision: take 1 step, take 2 steps
	- find the minimum cost to iterate through costs starting at 0 or 1
	- from i
		- take 1 step
			- have to get from i+1 to end, i < len(costs)
		- take 2 steps
			- have to get from i+2 to end, i < len(costs)

	- recurrence relation
	- greedy wouldn't work, have to explore both paths
	- minCost_i = cost[i] + min(minCost_i1, minCost_i2)

	- Big O
		- unique states O(len(cost))
		- cache complexity O(1)
*/

// O(n) / O(1)
func minCostClimbingStairs(cost []int) int {
	plusOne, plusTwo := 0, 0

	for i := len(cost) - 1; i >= 0; i-- {
		plusOne, plusTwo = cost[i]+min(plusOne, plusTwo), plusOne
	}

	return min(plusOne, plusTwo)
}

// O(n) / O(n)
// func minCostClimbingStairs(cost []int) int {
// 	cache := make([]int, len(cost)+2)
//
// 	for i := len(cost) - 1; i >= 0; i-- {
// 		cache[i] = cost[i] + min(cache[i+1], cache[i+2])
// 	}
//
// 	return min(cache[0], cache[1])
// }

// O(n) / O(n)
// func minCostClimbingStairs(cost []int) int {
// 	cache := make([]int, len(cost))
// 	for i := range cache {
// 		cache[i] = -1
// 	}
//
// 	var minCost func(i int) int
// 	minCost = func(i int) int {
// 		if i >= len(cost) {
// 			return 0
// 		}
// 		if i == len(cost)-1 {
// 			return cost[i]
// 		}
//
// 		if cache[i] == -1 {
// 			oneStep := minCost(i + 1)
// 			twoStep := minCost(i + 2)
//
// 			cache[i] = cost[i] + min(oneStep, twoStep)
// 		}
// 		return cache[i]
// 	}
//
// 	return min(minCost(0), minCost(1))
// }

func main() {
	tests := []struct {
		in  string
		out int
	}{
		{
			in:  "[10,15,20]",
			out: 15,
		},
		{
			in:  "[1,100,1,1,1,100,1,1,100,1]",
			out: 6,
		},
	}

	for _, t := range tests {
		var in []int
		json.Unmarshal([]byte(t.in), &in)
		start := time.Now()
		res := minCostClimbingStairs(in)
		fmt.Printf("%v == %v | %v\n", t.out, res, time.Since(start))
	}
}
