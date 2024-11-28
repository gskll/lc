package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// https://leetcode.com/problems/binary-tree-maximum-path-sum/description/
//
// A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.
//
// The path sum of a path is the sum of the node's values in the path.
//
// Given the root of a binary tree, return the maximum path sum of any non-empty path.
//
// The number of nodes in the tree is in the range [1, 3 * 10^4].
// -1000 <= Node.val <= 1000

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxSum := root.Val

	var walk func(node *TreeNode) int
	walk = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftSum := walk(node.Left)
		rightSum := walk(node.Right)

		leftGain := max(leftSum, 0)
		rightGain := max(rightSum, 0)

		maxSum = max(maxSum, node.Val+leftGain+rightGain)

		return node.Val + max(leftGain, rightGain)
	}

	walk(root)

	return maxSum
}

func main() {
	tests := []struct {
		root string
		out  int
	}{
		{
			root: "[1,2,3]",
			out:  6,
		},
		{
			root: "[-10,9,20,null,null,15,7]",
			out:  42,
		},
		{
			root: "[-3]",
			out:  -3,
		},
		{
			root: "[2,-1]",
			out:  2,
		},
		{
			root: "[-1,-2,10,-6,null,-3,-6]",
			out:  10,
		},
		{
			root: "[1,-2,3]",
			out:  4,
		},
		{
			root: "[9,6,-3,null,null,-6,2,null,null,2,null,-6,-6,-6]",
			out:  16,
		},
		{
			root: "[-1,null,9,-6,3,null,null,null,-2]",
			out:  12,
		},
		{
			root: "[7,-8,2,1,null,4,-8,null,null,null,-5,null,5]",
			out:  13,
		},
	}
	for _, test := range tests {
		start := time.Now()
		root := TreeFromLeetCode(test.root)
		res := maxPathSum(root)
		fmt.Printf("elapsed: %s\n", time.Since(start))
		fmt.Printf("got: %+v\n", res)
		fmt.Printf("exp: %+v\n", test.out)

	}
}

// Builds a tree from LeetCode format string "[1,2,3,null,null,4,5]"
func TreeFromLeetCode(input string) *TreeNode {
	nodes := parseJSON(input)
	return buildTreeT(nodes)
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

func buildTreeT(arr []*int) *TreeNode {
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
