package main

import (
	"fmt"
)

// https://leetcode.com/problems/find-the-duplicate-number/description/

// Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.
//
// There is only one repeated number in nums, return this repeated number.
//
// You must solve the problem without modifying the array nums and using only constant extra space.
//
// 1 <= n <= 105
// nums.length == n + 1
// 1 <= nums[i] <= n
// All the integers in nums appear only once except for precisely one integer which appears two or more times.

type ListNode struct {
	Val  int
	Next *ListNode
}

// PIGEONHOLE THEORY
// if you have n items to put into m containers, and n > m, then at least one container must contain more than one item.

// since we have length n+1, and [1,n] possibilities for each number, we have n pigeonholes and n+1 pigeons, so there must be a duplicate

// INTUITION
// since we have n+1 length and [1,n] number possibilities, we can think of the array as a linked list where the value at each index points to that index
// duplicates will create a cycle

// [1,3,4,2,2]
// 0->1->3->2->4
//          _|
//
// we can use floyd's cycle detection to find the start of the cycle which will be the duplicate value

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]] // head.next, head.next.next

	// find intersection point in cycle
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	// find cycle start == duplicate
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}

func main() {
	type test struct {
		nums1 []int
		out   int
	}
	tests := []test{
		{nums1: []int{1, 3, 4, 2, 2}, out: 2},
		{nums1: []int{3, 1, 3, 4, 2}, out: 3},
		{nums1: []int{3, 3, 3, 3, 3}, out: 3},
		{nums1: []int{1, 3, 4, 2, 1}, out: 1},
	}
	for i, t := range tests {
		res := findDuplicate(t.nums1)
		fmt.Printf("%d: %+v --> %+v == %+v\n", i, t.nums1, res, t.out)
	}
}
