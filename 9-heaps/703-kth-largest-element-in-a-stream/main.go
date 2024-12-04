package main

import "fmt"

// https://leetcode.com/problems/kth-largest-element-in-a-stream/description/
//
// You are part of a university admissions office and need to keep track of the kth highest test score from applicants in real-time. This helps to determine cut-off marks for interviews and admissions dynamically as new applicants submit their scores.
//
// You are tasked to implement a class which, for a given integer k, maintains a stream of test scores and continuously returns the kth highest test score after a new score has been submitted. More specifically, we are looking for the kth highest score in the sorted list of all scores.
//
// Implement the KthLargest class:
//
// KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of test scores nums.
// int add(int val) Adds a new test score val to the stream and returns the element representing the kth largest element in the pool of test scores so far.

// Time: O(n*logk) for construction, O(logk) for add
// Space: O(k)
type KthLargest struct {
	heap []int
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{
		heap: make([]int, 0, k),
		k:    k,
	}

	for _, n := range nums {
		kl.Add(n)
	}

	return kl
}

func (this *KthLargest) Add(val int) int {
	if len(this.heap) < this.k {
		this.heap = append(this.heap, val)
		this.heapifyUp(len(this.heap) - 1)
	} else if this.heap[0] < val {
		this.heap[0] = val
		this.heapifyDown(0)
	}
	return this.heap[0]
}

func (this *KthLargest) heapifyUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2

		if this.heap[i] >= this.heap[parent] {
			break
		}

		this.heap[i], this.heap[parent] = this.heap[parent], this.heap[i]
		i = parent
	}
}

func (this *KthLargest) heapifyDown(i int) {
	n := len(this.heap)
	for {

		smallest := i
		left := 2*i + 1
		right := 2*i + 2

		if left < n && this.heap[smallest] > this.heap[left] {
			smallest = left
		}

		if right < n && this.heap[smallest] > this.heap[right] {
			smallest = right
		}

		if smallest == i {
			break
		}
		this.heap[i], this.heap[smallest] = this.heap[smallest], this.heap[i]
		i = smallest
	}
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func main() {
	kthLargest := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(kthLargest.Add(3), "==", 4)  // return 4
	fmt.Println(kthLargest.Add(5), "==", 5)  // return 5
	fmt.Println(kthLargest.Add(10), "==", 5) // return 5
	fmt.Println(kthLargest.Add(9), "==", 8)  // return 8
	fmt.Println(kthLargest.Add(4), "==", 8)  // return 8
	fmt.Println()
	fmt.Println()
	kthLargest = Constructor(4, []int{7, 7, 7, 7, 8, 3})
	fmt.Println(kthLargest.Add(2), "==", 7)  // return 7
	fmt.Println(kthLargest.Add(10), "==", 7) // return 7
	fmt.Println(kthLargest.Add(9), "==", 7)  // return 7
	fmt.Println(kthLargest.Add(9), "==", 8)  // return 8
	fmt.Println()
	fmt.Println()
	kthLargest = Constructor(1, []int{})
	fmt.Println(kthLargest.Add(-3), "==", -3)
	fmt.Println(kthLargest.Add(-2), "==", -2)
	fmt.Println(kthLargest.Add(-4), "==", -2)
	fmt.Println(kthLargest.Add(0), "==", 0)
	fmt.Println(kthLargest.Add(4), "==", 4)
	fmt.Println()
	fmt.Println()
	kthLargest = Constructor(2, []int{0})
	fmt.Println(kthLargest.Add(-1), "==", -1)
	fmt.Println(kthLargest.Add(1), "==", 0)
	fmt.Println(kthLargest.Add(-2), "==", 0)
	fmt.Println(kthLargest.Add(-4), "==", 0)
	fmt.Println(kthLargest.Add(3), "==", 1)
}
