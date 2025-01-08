package main

import "fmt"

// https://leetcode.com/problems/climbing-stairs/description/

// You are climbing a staircase. It takes n steps to reach the top.
//
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//
// 1 <= n <= 45

// O(n) / O(1)
func climbStairs(n int) int {
	one, two := 1, 1

	for i := 0; i < n-1; i++ {
		one, two = two+one, one
		// tmp := one
		// one = two + one
		// two = tmp
	}

	return one
}

// O(n) / O(n)
// func climbStairs(n int) int {
// 	if n <= 2 {
// 		return n
// 	}
// 	cache := make([]int, n+1)
//
// 	cache[1] = 1
// 	cache[2] = 2
//
// 	for i := 3; i <= n; i++ {
// 		cache[i] = cache[i-1] + cache[i-2]
// 	}
//
// 	return cache[n]
// }

// func climbStairs(n int) int {
// 	cache := make([]int, n+1)
// 	for i := range cache {
// 		cache[i] = -1
// 	}
//
// 	var solve func(i int) int
// 	solve = func(i int) int {
// 		if i > n {
// 			return 0
// 		}
//
// 		if i == n {
// 			return 1
// 		}
//
// 		if cache[i] != -1 {
// 			return cache[i]
// 		}
//
// 		cache[i] = solve(i+1) + solve(i+2)
// 		return cache[i]
// 	}
//
// 	return solve(0)
// }

// func climbStairs(n int) int {
// 	cache := make([]int, n+1)
// 	for i := range cache {
// 		cache[i] = -1
// 	}
//
// 	var solve func(i int) int
// 	solve = func(i int) int {
// 		if i < 0 {
// 			return 0
// 		}
//
// 		if i == 0 {
// 			return 1
// 		}
//
// 		if cache[i] != -1 {
// 			return cache[i]
// 		}
//
// 		cache[i] = solve(i-1) + solve(i-2)
// 		return cache[i]
// 	}
//
// 	return solve(n)
// }

// func climbStairs(n int) int {
// 	memo := make(map[int]int, n)
//
// 	var solve func(n int) int
// 	solve = func(n int) int {
// 		if x, ok := memo[n]; ok {
// 			fmt.Println("memo: ", n, x)
// 			return x
// 		}
//
// 		if n < 0 {
// 			return 0
// 		}
// 		if n == 0 {
// 			return 1
// 		}
//
// 		x := solve(n - 1)
// 		memo[n-1] = x
// 		y := solve(n - 2)
// 		memo[n-2] = y
//
// 		return x + y
// 	}
// 	return solve(n)
// }

func main() {
	tests := []struct {
		in  int
		out int
	}{
		{
			in:  2,
			out: 2,
		},
		{
			in:  3,
			out: 3,
		},
	}

	for _, t := range tests {
		res := climbStairs(t.in)
		fmt.Printf("%v == %v\n", t.out, res)
	}
}
