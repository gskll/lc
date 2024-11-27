package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/

// Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.
//
// The number of nodes in the tree is n.
// 1 <= k <= n <= 104
// 0 <= Node.val <= 104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//// Optimal follow up

type OptimalTreeNode struct {
	Val       int
	Left      *TreeNode
	Right     *TreeNode
	LeftCount int
}

func kthSmallestFollowup(root *OptimalTreeNode, k int) int {
	curr := root
	for curr != nil {
		leftCount := 0
		if curr.Left != nil {
			leftCount = curr.LeftCount
		}

		if k == leftCount+1 {
			return curr.Val
		}
		if k <= leftCount {
			curr = curr.Left
		} else {
			k = k - (leftCount + 1)
			curr = curr.Right
		}
	}
	return -1
}

func insert(root *OptimalTreeNode, val int) *OptimalTreeNode {
	if root == nil {
		return &OptimalTreeNode{Val: val}
	}
	if val < root.Val {
		root.LeftCount++
		root.Left = insert(root.Left, val)
	} else {
		root.Right = insert(root.Right, val)
	}
	return root
}

func delete(root *OptimalTreeNode, val int) *OptimalTreeNode {
	if root == nil {
		return nil
	}
	if val < root.Val {
		root.LeftCount--
		root.Left = delete(root.Left, val)
	} else if val > root.Val {
		root.Right = delete(root.Right, val)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		minNode := findMin(root.Right)
		root.Val = minNode.Val
		root.Right = delete(root.Right, minNode.Val)
	}
	return root
}

func findMin(node *TreeNode) *TreeNode {
	curr := node
	for curr.Left != nil {
		curr = curr.Left
	}
	return curr
}

// Heap solution not optimal!! for follow up we use a left counter in the node that can be used to handle see in linear time how many nodes are smaller. we want to return
type MaxHeap struct {
	items []int
	size  int
}

func (h *MaxHeap) Len() int {
	return len(h.items)
}

func (h *MaxHeap) Less(i, j int) bool {
	return h.items[i] > h.items[j]
}

func (h *MaxHeap) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *MaxHeap) Push(x any) {
	h.items = append(h.items, x.(int))
}

func (h *MaxHeap) Pop() any {
	x := h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	return x
}

func (h *MaxHeap) PushVal(x int) {
	heap.Push(h, x)
}

func (h *MaxHeap) PopVal() int {
	return heap.Pop(h).(int)
}

func (h *MaxHeap) Top() int {
	if h.Len() == 0 {
		return -1
	}
	return h.items[0]
}

// Follow up: If the BST is modified often (i.e., we can do insert and delete operations) and you need to find the kth smallest frequently, how would you optimize?
// use heap - O(k*logk), O(k)
func kthSmallestHeap(root *TreeNode, k int) int {
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if maxHeap.Len() >= k || node == nil {
			return
		}

		if node.Left != nil {
			dfs(node.Left)
		}

		if maxHeap.Len() < k && maxHeap.Top() < node.Val {
			maxHeap.PushVal(node.Val)
		}

		if maxHeap.Len() < k && node.Right != nil {
			dfs(node.Right)
		}
	}

	dfs(root)
	return maxHeap.Top()
}

// in order traversal - O(k), O(k)
func kthSmallestTraverse(root *TreeNode, k int) int {
	var inOrder []int

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if len(inOrder) == k || node == nil {
			return
		}

		if node.Left != nil {
			dfs(node.Left)
		}

		inOrder = append(inOrder, node.Val)

		if node.Right != nil {
			dfs(node.Right)
		}
	}

	dfs(root)
	return inOrder[k-1]
}

// counter
func kthSmallest(root *TreeNode, k int) int {
	var counter int
	var result int

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if counter == k || node == nil {
			return
		}

		if node.Left != nil {
			dfs(node.Left)
		}

		counter++
		if counter == k {
			result = node.Val
			return
		}

		if node.Right != nil {
			dfs(node.Right)
		}
	}

	dfs(root)
	return result
}

func main() {
	tests := []struct {
		root string
		k    int
		out  int
	}{
		{
			root: "[3,1,4,null,2]",
			k:    1,
			out:  1,
		},
		{
			root: "[5,3,6,2,4,null,null,1]",
			k:    3,
			out:  3,
		},
	}
	for _, test := range tests {
		root := TreeFromLeetCode(test.root)
		r := kthSmallest(root, test.k)
		fmt.Printf("got: %+v\n", r)
		fmt.Printf("exp: %+v\n", test.out)

	}
}

// Builds a tree from LeetCode format string "[1,2,3,null,null,4,5]"
func TreeFromLeetCode(input string) *TreeNode {
	nodes := parseJSON(input)
	return buildTree(nodes)
}

// Prints tree in LeetCode array format
func (t *TreeNode) ToLeetCode() string {
	if t == nil {
		return "[]"
	}

	result := []*int{}
	queue := []*TreeNode{t}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, nil)
			continue
		}

		val := node.Val
		result = append(result, &val)

		if node.Left != nil || node.Right != nil {
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	// Trim trailing nulls
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	// Convert to string
	parts := make([]string, len(result))
	for i, v := range result {
		if v == nil {
			parts[i] = "null"
		} else {
			parts[i] = strconv.Itoa(*v)
		}
	}

	return "[" + strings.Join(parts, ",") + "]"
}

// Prints tree in a visual hierarchical format
func (t *TreeNode) String() string {
	var traverse func(*TreeNode, int) string
	traverse = func(root *TreeNode, level int) string {
		if root == nil {
			return ""
		}
		prefix := strings.Repeat("  ", level)
		result := fmt.Sprintf("%s(%d)\n", prefix, root.Val)

		// Print both children at the same level
		if root.Left != nil {
			result += traverse(root.Left, level+1)
		}
		if root.Right != nil {
			result += traverse(root.Right, level+1)
		}
		return result
	}
	return strings.TrimRight(traverse(t, 0), "\n")
}

// Helper functions
func parseJSON(input string) []*int {
	input = strings.TrimSpace(input)
	input = strings.Trim(input, "[]")

	if input == "" {
		return []*int{}
	}

	parts := strings.Split(input, ",")
	result := make([]*int, len(parts))

	for i, p := range parts {
		p = strings.TrimSpace(p)
		if p == "null" {
			result[i] = nil
			continue
		}
		num, _ := strconv.Atoi(p)
		result[i] = &num
	}

	return result
}

func buildTree(arr []*int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	if arr[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *arr[0]}
	nodes := []*TreeNode{root}
	i := 1

	for i < len(arr) {
		curr := nodes[0]
		nodes = nodes[1:]

		if i < len(arr) && arr[i] != nil {
			curr.Left = &TreeNode{Val: *arr[i]}
			nodes = append(nodes, curr.Left)
		}
		i++

		if i < len(arr) && arr[i] != nil {
			curr.Right = &TreeNode{Val: *arr[i]}
			nodes = append(nodes, curr.Right)
		}
		i++
	}

	return root
}
