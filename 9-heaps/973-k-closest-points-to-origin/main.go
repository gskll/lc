package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

// https://leetcode.com/problems/k-closest-points-to-origin/description/
//
// Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane and an integer k, return the k closest points to the origin (0, 0).
//
// The distance between two points on the X-Y plane is the Euclidean distance (i.e., √(x1 - x2)2 + (y1 - y2)2).
//
// You may return the answer in any order. The answer is guaranteed to be unique (except for the order that it is in).
//
// 1 <= k <= points.length <= 10^4
// -10^4 <= xi, yi <= 10^4

func euclidean(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}

func isCloser(p1, p2 []int) bool {
	return euclidean(p1) < euclidean(p2)
}

func maxHeapifyDown(h [][]int, i int) {
	n := len(h)
	for {
		largest := i
		left := 2*i + 1
		right := left + 1

		if left < n && isCloser(h[largest], h[left]) {
			largest = left
		}

		if right < n && isCloser(h[largest], h[right]) {
			largest = right
		}

		if largest == i {
			break
		}

		h[i], h[largest] = h[largest], h[i]
		i = largest
	}
}

func kClosestHeap(points [][]int, k int) [][]int {
	h := points[:k]

	for i := len(h)/2 - 1; i >= 0; i-- {
		maxHeapifyDown(h, i)
	}

	for _, p := range points[k:] {
		if isCloser(p, h[0]) {
			h[0] = p
			maxHeapifyDown(h, 0)
		}
	}

	return h
}

func kClosest(points [][]int, k int) [][]int {
	pi := quickselect(points, 0, len(points)-1, k-1)
	return points[:pi+1]
}

func quickselect(points [][]int, low, high, k int) int {
	for low < high {
		pi := partitionHoare(points, low, high)

		if k <= pi {
			high = pi
		} else {
			low = pi + 1
		}
	}
	return low
}

func partitionHoare(points [][]int, low, high int) int {
	pi := low + rand.Intn(high-low+1)
	points[low], points[pi] = points[pi], points[low]

	pivot := points[low]
	left := low - 1
	right := high + 1

	for {
		for {
			left++
			if euclidean(points[left]) >= euclidean(pivot) {
				break
			}
		}

		for {
			right--
			if euclidean(points[right]) <= euclidean(pivot) {
				break
			}
		}

		if left >= right {
			return right
		}

		points[left], points[right] = points[right], points[left]
	}
}

func main() {
	tests := []struct {
		points string
		k      int
		out    string
	}{
		{
			points: "[[1,3],[-2,2]]",
			k:      1,
			out:    "[[-2,2]]",
		},
		{
			points: "[[3,3],[5,-1],[-2,4]]",
			k:      2,
			out:    "[[3,3],[-2,4]]",
		},
	}

	for _, t := range tests {
		var points [][]int
		var out [][]int
		json.Unmarshal([]byte(t.points), &points)
		json.Unmarshal([]byte(t.out), &out)
		res := kClosest(points, t.k)
		fmt.Println("exp:", out, "got:", res)
	}
}
