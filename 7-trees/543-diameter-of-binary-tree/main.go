package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/diameter-of-binary-tree/description/

// Given the root of a binary tree, return the length of the diameter of the tree.
//
// The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.
//
// The length of a path between two nodes is represented by the number of edges between them.
//
// The number of nodes in the tree is in the range [1, 10^4].
// -100 <= Node.val <= 100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeInfo struct {
	height   int
	diameter int
}

// O(n) / O(n) --> less memory efficient than the recusrive solution
// have to track info in a map, recursive solution uses the call stack
func diameterOfBinaryTreeIterative(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := []*TreeNode{root}
	seen := map[*TreeNode]NodeInfo{nil: {}}

	for len(stack) > 0 {
		node := stack[len(stack)-1]

		if _, seen := seen[node.Left]; !seen && node.Left != nil {
			stack = append(stack, node.Left)
			continue
		}

		if _, seen := seen[node.Right]; !seen && node.Right != nil {
			stack = append(stack, node.Right)
			continue
		}

		// both children have been processed -> pop stack
		stack = stack[:len(stack)-1]

		// if nil we use 0 values
		left := seen[node.Left]
		right := seen[node.Right]

		height := 1 + max(left.height, right.height)
		diameter := max(
			left.height+right.height,           // diameter through current node
			max(left.diameter, right.diameter), // longest diameter in subtrees
		)

		seen[node] = NodeInfo{height, diameter}

	}

	return seen[root].diameter
}

// recursive O(n) / O(h) (O(h) is O(logn) for balanced tree or O(n) for linear)
func diameterOfBinaryTreeRecursive(root *TreeNode) int {
	var diameter int
	var height func(node *TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := height(node.Left)
		right := height(node.Right)

		diameter = max(diameter, left+right)

		return max(left, right) + 1
	}

	height(root)
	return diameter
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxRight := maxDepth(root.Right)
	maxLeft := maxDepth(root.Left)

	return max(maxLeft, maxRight) + 1
}

// brute force: O(n^2) / O(n)
func diameterOfBinaryTreeBruteForce(root *TreeNode) int {
	if root == nil {
		return 0
	}

	currDiameter := maxDepth(root.Left) + maxDepth(root.Right)
	leftDiameter := diameterOfBinaryTree(root.Left)
	rightDiameter := diameterOfBinaryTree(root.Right)

	return max(currDiameter, leftDiameter, rightDiameter)
}

func main() {
	tests := []struct {
		in  string
		out int
	}{
		{
			in:  "[1,2,3,4,5]",
			out: 3,
		},
		{
			in:  "[1,2]",
			out: 1,
		},
		{
			in:  "[2,3,null,1]",
			out: 2,
		},
		{
			in:  "[1]",
			out: 0,
		},
	}
	for _, test := range tests {
		t := TreeFromLeetCode(test.in)
		r := diameterOfBinaryTree(t)
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
