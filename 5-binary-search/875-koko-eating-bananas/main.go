package main

import (
	"fmt"
)

// https://leetcode.com/problems/koko-eating-bananas/description/

// Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. The guards have gone and will come back in h hours.
//
// Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas from that pile. If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.
//
// Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.
//
// Return the minimum integer k such that she can eat all the bananas within h hours.
//
// 1 <= piles.length <= 10^4
// piles.length <= h <= 10^9
// 1 <= piles[i] <= 10^9

func minEatingSpeed(piles []int, h int) int {
	sum, maxP := 0, 0
	for _, p := range piles {
		if p > maxP {
			maxP = p
		}
		sum += p
	}

	l, r := (sum+h-1)/h, maxP

	for l < r {
		k := l + (r-l)/2

		hours := 0
		for _, p := range piles {
			hours += (p + k - 1) / k
			if hours > h {
				break
			}
		}

		if hours <= h {
			r = k
		} else {
			l = k + 1
		}
	}
	return l
}

func main() {
	type test struct {
		piles []int
		h     int
		out   int
	}
	tests := []test{
		{piles: []int{3, 6, 7, 11}, h: 8, out: 4},
		{piles: []int{30, 11, 23, 4, 20}, h: 5, out: 30},
		{piles: []int{30, 11, 23, 4, 20}, h: 6, out: 23},
		{piles: []int{312884470}, h: 968709470, out: 1},
	}
	for i, t := range tests {
		res := minEatingSpeed(t.piles, t.h)
		fmt.Printf("%d: %+v %d --> %+v == %+v\n", i, t.piles, t.h, res, t.out)
	}
}
