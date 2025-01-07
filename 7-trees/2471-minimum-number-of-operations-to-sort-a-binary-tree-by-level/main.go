package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/minimum-number-of-operations-to-sort-a-binary-tree-by-level/?envType=daily-question&envId=2024-12-23

// You are given the root of a binary tree with unique values.
//
// In one operation, you can choose any two nodes at the same level and swap their values.
//
// Return the minimum number of operations needed to make the values at each level sorted in a strictly increasing order.
//
// The level of a node is the number of edges along the path between it and the root node.
//
// The number of nodes in the tree is in the range [1, 10^5].
// 1 <= Node.val <= 10^5
// All the values of the tree are unique.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func minSwaps(arr []int) int {
	pairs := make([][2]int, len(arr))
	for i, n := range arr {
		pairs[i] = [2]int{i, n}
	}

	slices.SortFunc(pairs, func(x, y [2]int) int {
		return x[1] - y[1]
	})

	visited := make([]bool, len(arr))
	swaps := 0

	for i := range pairs {
		if visited[i] || pairs[i][0] == i {
			continue
		}

		cycleSize := 0
		j := i
		for !visited[j] {
			visited[j] = true
			j = pairs[j][0]
			cycleSize++
		}
		swaps += cycleSize - 1
	}
	return swaps
}

func minimumOperations(root *TreeNode) int {
	count := 0
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, 0, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			level = append(level, node.Val)
		}

		count += minSwaps(level)

		// for i := 0; i < numNodesInLevel-1; i++ {
		// 	levelMin := i
		// 	for j := i + 1; j < numNodesInLevel; j++ {
		// 		if levelNodes[j] < levelNodes[levelMin] {
		// 			levelMin = j
		// 		}
		// 	}
		//
		// 	if levelMin != i {
		// 		count++
		// 		levelNodes[i], levelNodes[levelMin] = levelNodes[levelMin], levelNodes[i]
		// 	}
		// }
	}
	return count
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tests := []struct {
		in  string
		out int
	}{
		{
			in:  "[1,4,3,7,6,8,5,null,null,null,null,9,null,10]",
			out: 3,
		},
		{
			in:  "[1,3,2,7,6,5,4]",
			out: 3,
		},
		{
			in:  "[1,2,3,4,5,6]",
			out: 0,
		},
	}
	for _, test := range tests {
		t := TreeFromLeetCode(test.in)
		r := minimumOperations(t)
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
