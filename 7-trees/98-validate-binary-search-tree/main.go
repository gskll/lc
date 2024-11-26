package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/validate-binary-search-tree/description/

// Given the root of a binary tree, determine if it is a valid binary search tree (BST).
//
// A valid BST is defined as follows:
//
// The left subtree of a node contains only nodes with keys less than the node's key.
// The right subtree of a node contains only nodes with keys greater than the node's key.
// Both the left and right subtrees must also be binary search trees.
//
// The number of nodes in the tree is in the range [1, 10^4].
// -23^1 <= Node.val <= 2^31 - 1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// valid ranges: preferred: more intuitive, easier to adapt to BST variants, catches invalid nodes earlier
// O(n)/ O(h)
func isValidBST1(root *TreeNode) bool {
	invalid := false

	var dfs func(node *TreeNode, minVal, maxVal int)
	dfs = func(node *TreeNode, minVal, maxVal int) {
		if invalid || node == nil {
			return
		}

		if minVal >= node.Val || maxVal <= node.Val {
			invalid = true
			return
		}

		dfs(node.Left, minVal, (node.Val))
		dfs(node.Right, (node.Val), maxVal)
	}

	dfs(root, math.MinInt, math.MaxInt)

	return !invalid
}

// in order traversal O(n) / O(h)
func isValidBST(root *TreeNode) bool {
	invalid := false

	var prev *int
	var inOrderTraversal func(node *TreeNode)
	inOrderTraversal = func(node *TreeNode) {
		if node == nil || invalid {
			return
		}

		inOrderTraversal(node.Left)

		if prev != nil && *prev >= node.Val {
			invalid = true
			return
		}

		prev = &node.Val

		inOrderTraversal(node.Right)
	}
	inOrderTraversal(root)

	return !invalid
}

func main() {
	tests := []struct {
		in  string
		out bool
	}{
		{
			in:  "[2,1,3]",
			out: true,
		},
		{
			in:  "[5,1,4,null,null,3,6]",
			out: false,
		},
		{
			in:  "[5,4,6,null,null,3,7]",
			out: false,
		},
	}
	for _, test := range tests {
		t := TreeFromLeetCode(test.in)
		r := isValidBST(t)
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
