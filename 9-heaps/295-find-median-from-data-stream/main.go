package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.com/problems/find-median-from-data-stream/description/

// The median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value, and the median is the mean of the two middle values.
//
// For example, for arr = [2,3,4], the median is 3.
// For example, for arr = [2,3], the median is (2 + 3) / 2 = 2.5.
// Implement the MedianFinder class:
//
// MedianFinder() initializes the MedianFinder object.
// void addNum(int num) adds the integer num from the data stream to the data structure.
// double findMedian() returns the median of all elements so far. Answers within 10-5 of the actual answer will be accepted.
//
// Constraints:
//
// -10^5 <= num <= 10^5
// There will be at least one element in the data structure before calling findMedian.
// At most 5 * 10^4 calls will be made to addNum and findMedian.
//
//
// Follow up:
//
// If all integer numbers from the stream are in the range [0, 100], how would you optimize your solution?
// If 99% of all integer numbers from the stream are in the range [0, 100], how would you optimize your solution?

type minHeap []int

func (h minHeap) Len() int { return len(h) }

func (h minHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h minHeap) Peek() int { return h[0] }

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	top := old[n-1]
	*h = old[:n-1]
	return top
}

type maxHeap []int

func (h maxHeap) Len() int { return len(h) }

func (h maxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h maxHeap) Peek() int { return h[0] }
func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	top := old[n-1]
	*h = old[:n-1]
	return top
}

type MedianFinder struct {
	right *minHeap
	left  *maxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		left:  &maxHeap{},
		right: &minHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.left.Len() == 0 {
		heap.Push(this.left, num)
		return
	}

	if num < this.left.Peek() {
		heap.Push(this.left, num)
	} else {
		heap.Push(this.right, num)
	}

	// Balance the heaps
	// Left heap can have at most one more element than right heap
	if this.left.Len() > this.right.Len()+1 {
		heap.Push(this.right, heap.Pop(this.left))
	}
	if this.right.Len() > this.left.Len() {
		heap.Push(this.left, heap.Pop(this.right))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.Len() > this.right.Len() {
		return float64(this.left.Peek())
	} else {
		return float64(this.left.Peek()+this.right.Peek()) / 2
	}
}

// FOLLOW UP - NUMBERS BETWEEN [0-100]
type MedianFinderCount struct {
	counts [101]int
	size   int
}

func ConstructorCount() MedianFinderCount {
	return MedianFinderCount{}
}

func (this *MedianFinderCount) AddNum(num int) {
	this.counts[num]++
	this.size++
}

func (this *MedianFinderCount) FindMedian() float64 {
	if this.size == 0 {
		return 0
	}

	if this.size%2 == 1 {
		target := (this.size / 2) + 1
		count := 0
		for i := range this.counts {
			count += this.counts[i]
			if count >= target {
				return float64(i)
			}
		}
	}

	count := 0
	target1, target2 := this.size/2, (this.size/2)+1
	for i := range this.counts {
		count += this.counts[i]

		if count >= target2 {
			return float64(i)
		}

		if count >= target1 {
			return float64(i+i+1) / 2
		}
	}

	return 0
}

// FOLLOW UP 2 - 99% NUMBERS BETWEEN [0-100]

type MedianFinderHybrid struct {
	left    *maxHeap
	right   *minHeap
	counts  [101]int
	size    int
	inRange int
}

func ConstructorHybrid() MedianFinderHybrid {
	return MedianFinderHybrid{
		left:  &maxHeap{},
		right: &minHeap{},
	}
}

func (this *MedianFinderHybrid) AddNum(num int) {
	this.size++
	if num < 0 {
		heap.Push(this.left, num)
	} else if num > 100 {
		heap.Push(this.right, num)
	} else {
		this.counts[num]++
		this.inRange++
	}
}

func (this *MedianFinderHybrid) FindMedian() float64 {
	if this.size == 0 {
		return 0
	}

	var getValueAtPos func(pos int) int
	getValueAtPos = func(pos int) int {
		count := this.left.Len()

		// median is in left heap
		if count >= pos {
			tmp := make([]int, 0, count)
			for i := 0; i < pos; i++ {
				tmp = append(tmp, heap.Pop(this.left).(int))
			}
			val := tmp[len(tmp)-1]
			for _, n := range tmp {
				heap.Push(this.left, n)
			}
			return val
		}

		// check count array
		for i, c := range this.counts {
			count += c
			if count >= pos {
				return i
			}
		}

		// median is in right heap
		rightPos := pos - count
		tmp := make([]int, 0, rightPos)
		for i := 0; i < rightPos; i++ {
			tmp = append(tmp, heap.Pop(this.right).(int))
		}
		val := tmp[len(tmp)-1]
		for _, n := range tmp {
			heap.Push(this.right, n)
		}
		return val
	}

	if this.size%2 == 1 {
		return float64(getValueAtPos((this.size / 2) + 1))
	}

	return float64(getValueAtPos(this.size/2)+getValueAtPos((this.size/2)+1)) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

// func main() {
// 	medianFinder := ConstructorCount()
// 	medianFinder.AddNum(1)                           // arr = [1]
// 	medianFinder.AddNum(2)                           // arr = [1, 2]
// 	fmt.Println(medianFinder.FindMedian(), "== 1.5") // return 1.5 (i.e., (1 + 2) / 2)
// 	medianFinder.AddNum(3)                           // arr[1, 2, 3]
// 	fmt.Println(medianFinder.FindMedian(), "== 2")   // return 2.0
//
// 	mf := ConstructorCount()
// 	mf.AddNum(1)
// 	mf.AddNum(2)
// 	mf.AddNum(2)
// 	mf.AddNum(3)
// 	fmt.Println(mf.FindMedian(), "== 2") // Should print 2.0
// }

func main() {
	fmt.Println("Running MedianFinderHybrid tests...")

	// Test 1: Simple in-range numbers
	fmt.Println("\nTest 1: Simple in-range numbers")
	mf1 := ConstructorHybrid()
	mf1.AddNum(1)
	mf1.AddNum(2)
	fmt.Printf("After [1,2]: %.1f == 1.5\n", mf1.FindMedian()) // Expected: 1.5
	mf1.AddNum(3)
	fmt.Printf("After [1,2,3]: %.1f == 2\n", mf1.FindMedian()) // Expected: 2.0

	// Test 2: Numbers spanning all structures
	fmt.Println("\nTest 2: Numbers spanning all structures")
	mf2 := ConstructorHybrid()
	mf2.AddNum(-5)                                                  // left heap
	mf2.AddNum(50)                                                  // counts array
	mf2.AddNum(200)                                                 // right heap
	fmt.Printf("After [-5,50,200]: %.1f == 50\n", mf2.FindMedian()) // Expected: 50.0

	// Test 3: Duplicates in counts array
	fmt.Println("\nTest 3: Duplicates in counts array")
	mf3 := ConstructorHybrid()
	mf3.AddNum(50)
	mf3.AddNum(50)
	mf3.AddNum(50)
	fmt.Printf("After [50,50,50]: %.1f == 50\n", mf3.FindMedian()) // Expected: 50.0

	// Test 4: Complex case
	fmt.Println("\nTest 4: Complex case")
	mf4 := ConstructorHybrid()
	nums := []int{-10, -5, 0, 50, 150, 200}
	for _, num := range nums {
		mf4.AddNum(num)
	}
	fmt.Printf("After [-10,-5,0,50,150,200]: %.1f == 25\n", mf4.FindMedian()) // Expected: 25.0

	// Test 5: Edge cases with repeated findMedian
	fmt.Println("\nTest 5: Edge cases with repeated findMedian")
	mf5 := ConstructorHybrid()
	mf5.AddNum(-5)
	mf5.AddNum(-10)
	fmt.Printf("First call: %.1f == -7.5\n", mf5.FindMedian())  // Expected: -7.5
	fmt.Printf("Second call: %.1f == -7.5\n", mf5.FindMedian()) // Expected: -7.5 (testing heap restoration)
}
