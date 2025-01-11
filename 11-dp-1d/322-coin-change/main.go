package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// https://leetcode.com/problems/coin-change/description/

// You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.
//
// Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.
//
// You may assume that you have an infinite number of each kind of coin.
//
// 1 <= coins.length <= 12
// 1 <= coins[i] <= 2^31 - 1
// 0 <= amount <= 10^4

/*
	- coins: [1,max_int32] - large
	- len(coins): [1,12] - small
	- amount: [0, 10000] - medium

	- return the fewest coins that can add up to amount or -1
	- optimization

	- infinite of each coin: can reuse coins
	- duplicates don't matter
	- take/skip not ideal as after we take we can still re-take the same coin

	- ordering? coins are in any order - might make sense to sort them big->small
	- greedy approach?

	- recurrence relation? subproblem?
	- for every n-c where n is amount and c is coins, we need to check which coin will give a smaller value
	- F(n) = 1 + min(F(n-c) for c in coins) // n >= c
	- F(n) = 0 // n == 0
	- F(n) = -1 // n < 0
*/

// bottom-up optimised: O(len(coins)*amount) / O(amount)
func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}

	dp[0] = 0
	for _, c := range coins {
		for amt := c; amt <= amount; amt++ {
			dp[amt] = min(dp[amt], dp[amt-c]+1)
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// top-down optimized
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = -1
	}

	var count func(remain int) int
	count = func(remain int) int {
		if remain < 0 {
			return -1
		}
		if remain == 0 {
			return 0
		}
		if dp[remain] != -1 {
			return dp[remain]
		}

		minCoins := amount + 1
		for _, c := range coins {
			if remain < c {
				continue
			}
			if minCoins <= 1 {
				break
			}
			minCoins = min(minCoins, count(remain-c)+1)
		}

		dp[remain] = minCoins
		return dp[remain]
	}

	res := count(amount)
	if res > amount {
		return -1
	}
	return res
}

// bottom-up O(len(coins)*amount) / O(amount)
// func coinChange(coins []int, amount int) int {
// 	dp := make([]int, amount+1)
// 	for i := range dp {
// 		dp[i] = math.MaxInt32
// 	}
//
// 	dp[0] = 0
// 	for i := 1; i <= amount; i++ {
// 		minCoins := math.MaxInt32
// 		for _, c := range coins {
// 			if i < c {
// 				continue
// 			}
// 			minCoins = min(minCoins, 1+dp[i-c])
// 		}
//
// 		dp[i] = minCoins
// 	}
//
// 	if dp[amount] == math.MaxInt32 {
// 		return -1
// 	}
// 	return dp[amount]
// }

// top-down O(len(coins)*amount) / O(amount)
// func coinChange(coins []int, amount int) int {
// 	dp := make([]int, amount+1)
//
// 	var count func(amount int) int
// 	count = func(amount int) int {
// 		if amount < 0 {
// 			return -1
// 		}
// 		if amount == 0 {
// 			return 0
// 		}
// 		if dp[amount] != 0 {
// 			return dp[amount]
// 		}
//
// 		minCoins := math.MaxInt32
// 		for _, c := range coins {
// 			subNum := count(amount - c)
// 			if subNum == -1 {
// 				continue
// 			}
// 			minCoins = min(minCoins, 1+subNum)
// 		}
//
// 		if minCoins == math.MaxInt32 {
// 			dp[amount] = -1
// 		} else {
// 			dp[amount] = minCoins
// 		}
//
// 		return dp[amount]
// 	}
//
// 	return count(amount)
// }

func main() {
	tests := []struct {
		coins  string
		amount int
		out    int
	}{
		{
			coins:  "[1,2,5]",
			amount: 11,
			out:    3,
		},
		{
			coins:  "[2]",
			amount: 3,
			out:    -1,
		},
		{
			coins:  "[1]",
			amount: 0,
			out:    0,
		},
		{
			coins:  "[186,419,83,408]",
			amount: 6249,
			out:    20,
		},
	}

	for _, t := range tests {
		var coins []int
		json.Unmarshal([]byte(t.coins), &coins)
		start := time.Now()
		res := coinChange(coins, t.amount)
		fmt.Printf("%v == %v | %v\n", t.out, res, time.Since(start))
	}
}
