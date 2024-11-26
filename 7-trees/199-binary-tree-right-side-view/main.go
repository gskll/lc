package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/binary-tree-right-side-view/description/

// Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.
//
// The number of nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs: O(n), O(h)
func rightSideView(root *TreeNode) []int {
	var res []int

	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if len(res) == level {
			res = append(res, node.Val)
		}

		dfs(node.Right, level+1)
		dfs(node.Left, level+1)
	}

	dfs(root, 0)

	return res
}

// bfs : O(n), O(w)
func rightSideViewBFS(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		for levelNode := len(queue) - 1; levelNode >= 0; levelNode-- {
			node := queue[0]
			queue = queue[1:]

			if levelNode == 0 {
				res = append(res, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return res
}

func main() {
	tests := []struct {
		in  string
		out string
	}{
		{
			in:  "[1,2,3,null,5,null,4]",
			out: "[1,3,4]",
		},
		{
			in:  "[[1,2,3,4,null,null,null,5]]",
			out: "[1,3,4,5]",
		},
		{
			in:  "[1,null,3]",
			out: "[1,3]",
		},
		{
			in:  "[]",
			out: "[]",
		},
	}
	for _, test := range tests {
		t := TreeFromLeetCode(test.in)
		r := rightSideView(t)
		var out []int
		json.Unmarshal([]byte(test.out), &out)
		fmt.Printf("got: %+v\n", r)
		fmt.Printf("exp: %+v\n", out)

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
