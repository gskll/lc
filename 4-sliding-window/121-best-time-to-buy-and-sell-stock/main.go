package main

import (
	"fmt"
)

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
// You are given an array prices where prices[i] is the price of a given stock on the ith day.
//
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
//
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

func maxProfit(prices []int) int {
	profit := 0
	l, r := 0, 1

	for r < len(prices) {
		buy, sell := prices[l], prices[r]
		if buy < sell {
			profit = max(sell-buy, profit)
		} else {
			l = r
		}
		r++
	}

	return profit
}

func main() {
	type test struct {
		prices []int
		out    int
	}
	tests := []test{
		{prices: []int{7, 1, 5, 3, 6, 4}, out: 5},
		{prices: []int{7, 6, 4, 3, 1}, out: 0},
	}
	for i, t := range tests {
		res := maxProfit(t.prices)
		fmt.Printf("%d: %+v --> %+v == %+v\n", i, t.prices, res, t.out)
	}
}
