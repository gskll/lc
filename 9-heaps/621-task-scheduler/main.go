package main

import (
	"fmt"
)

// https://leetcode.com/problems/task-scheduler/description/
// You are given an array of CPU tasks, each labeled with a letter from A to Z, and a number n. Each CPU interval can be idle or allow the completion of one task. Tasks can be completed in any order, but there's a constraint: there has to be a gap of at least n intervals between two tasks with the same label.
//
// Return the minimum number of CPU intervals required to complete all tasks.
//
// 1 <= tasks.length <= 10^4
// tasks[i] is an uppercase English letter.
// 0 <= n <= 100

// heap solution: O(t), O(1) memory -> constant as 26 characters for count
func leastInterval(tasks []byte, n int) int {
	var count [26]int
	for _, task := range tasks {
		count[task-'A']++
	}

	var heap MaxHeap
	for _, c := range count {
		if c > 0 {
			heap = append(heap, c)
		}
	}
	heap.build()

	queue := make(TaskQueue, 0, len(heap))
	time := 0

	for len(heap) > 0 || len(queue) > 0 {
		time++

		if top, exists := heap.pop(); exists {
			if top > 1 {
				queue.enqueue(top-1, time+n)
			}
		} else {
			time = queue.peek().nextAvailable
		}

		if len(queue) > 0 && time >= queue.peek().nextAvailable {
			freq := queue.dequeue().freq
			heap.insert(freq)
		}
	}

	return time
}

type QueueTask struct {
	freq          int
	nextAvailable int
}

type TaskQueue []QueueTask

func (q *TaskQueue) peek() QueueTask {
	return (*q)[0]
}

func (q *TaskQueue) enqueue(freq, ready int) {
	*q = append(*q, QueueTask{freq, ready})
}

func (q *TaskQueue) dequeue() QueueTask {
	first := (*q)[0]
	*q = (*q)[1:]
	return first
}

type MaxHeap []int

func (h *MaxHeap) build() {
	for i := len(*h)/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

func (h *MaxHeap) pop() (int, bool) {
	old := *h
	n := len(old)

	if n == 0 {
		return 0, false
	}

	top := old[0]
	old[0] = old[n-1]
	*h = old[:n-1]
	if len(*h) > 0 {
		h.heapifyDown(0)
	}
	return top, true
}

func (h *MaxHeap) insert(x int) {
	*h = append(*h, x)
	h.heapifyUp(len(*h) - 1)
}

func (hp *MaxHeap) heapifyUp(i int) {
	h := *hp
	for i > 0 {
		parent := (i - 1) / 2

		if h[parent] > h[i] {
			break
		}

		h[parent], h[i] = h[i], h[parent]
		i = parent
	}
	*hp = h
}

func (hp *MaxHeap) heapifyDown(i int) {
	h := *hp
	n := len(h)
	for {
		largest := i
		left := 2*i + 1
		right := left + 1

		if left < n && h[left] > h[largest] {
			largest = left
		}
		if right < n && h[right] > h[largest] {
			largest = right
		}

		if largest == i {
			break
		}

		h[i], h[largest] = h[largest], h[i]
		i = largest
	}
	*hp = h
}

// math solution: O(t) where t is len(tasks), O(1) memory
// find most frequent element ["A","A","A","B","B","B"]
// A__A__A, will be (n+1)*(maxFreq-1)+1 -> n+1 for each element, times how many of those element leaving out the last one, add the last one to the end
// but also 3 B's so AB_AB_AB so (n+1)*(maxFreq-1)+numMaxFreqElements
// could also have no gaps in between, in which case it's the num(tasks)
// formula is max(num(tasks), (n+1)*(maxFreq-1)+numMaxFreqElements)
func leastIntervalMath(tasks []byte, n int) int {
	var (
		count   [26]int
		maxN    int
		numMaxN int
	)

	for _, t := range tasks {
		count[t-'A']++
		if count[t-'A'] > maxN {
			maxN = count[t-'A']
		}
	}

	for _, c := range count {
		if c == maxN {
			numMaxN++
		}
	}

	return max(len(tasks), (n+1)*(maxN-1)+numMaxN)
}

func main() {
	tests := []struct {
		tasks []byte
		n     int
		out   int
	}{
		{
			tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:     2,
			out:   8,
		},
		{
			tasks: []byte{'A', 'C', 'A', 'B', 'D', 'B'},
			n:     1,
			out:   6,
		},
		{
			tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:     3,
			out:   10,
		},
	}

	for _, t := range tests {
		res := leastInterval(t.tasks, t.n)
		fmt.Println("exp:", t.out, "got:", res)
	}
}
