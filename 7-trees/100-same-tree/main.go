package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// https://leetcode.com/problems/same-tree/description/

// Given the roots of two binary trees p and q, write a function to check if they are the same or not.
//
// Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.
//
// The number of nodes in both trees is in the range [0, 100].
// -10^4 <= Node.val <= 10^4

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	// early return if left is not same
	if !isSameTree(p.Left, q.Left) {
		return false
	}

	return isSameTree(p.Right, q.Right)

	// return isSameTree(p.Left,q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	tests := []struct {
		p   string
		q   string
		out bool
	}{
		{
			p:   "[1,2,3]",
			q:   "[1,2,3]",
			out: true,
		},
		{
			p:   "[1,2]",
			q:   "[1,null,2]",
			out: false,
		},
		{
			p:   "[1,2,1]",
			q:   "[1,1,2]",
			out: true,
		},
	}
	for _, test := range tests {
		t := TreeFromLeetCode(test.p)
		t2 := TreeFromLeetCode(test.q)
		start := time.Now()
		r := isSameTree(t, t2)
		fmt.Printf("elapsed: %s\n", time.Since(start))
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
