package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/

// Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.
//
// According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”
//
// The number of nodes in the tree is in the range [2, 10^5].
// -10^9 <= Node.val <= 10^9
// All Node.val are unique.
// p != q
// p and q will exist in the BST.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// O(h) time, O(h) space for recursion stack
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}

// O(h) time, O(1) space
func lowestCommonAncestorIterative(root, p, q *TreeNode) *TreeNode {
	curr := root
	for curr != nil {
		if p.Val < curr.Val && q.Val < curr.Val {
			curr = curr.Left
		} else if p.Val > curr.Val && q.Val > curr.Val {
			curr = curr.Right
		} else {
			return curr
		}
	}

	return curr
}

func main() {
	tests := []struct {
		root string
		p    int
		q    int
		out  int
	}{
		{
			root: "[6,2,8,0,4,7,9,null,null,3,5]",
			p:    2,
			q:    8,
			out:  6,
		},
		{
			root: "[6,2,8,0,4,7,9,null,null,3,5]",
			p:    2,
			q:    4,
			out:  2,
		},
		{
			root: "[2,1]",
			p:    2,
			q:    1,
			out:  2,
		},
	}
	for _, test := range tests {
		r, p, q := TreeFromLeetCode(test.root, test.p, test.q)
		start := time.Now()
		res := lowestCommonAncestor(r, p, q)
		fmt.Printf("elapsed: %s\n", time.Since(start))
		fmt.Printf("got: %+v\n", res.Val)
		fmt.Printf("exp: %+v\n", test.out)

	}
}

// Builds a tree from LeetCode format string "[1,2,3,null,null,4,5]"
func TreeFromLeetCode(input string, p, q int) (*TreeNode, *TreeNode, *TreeNode) {
	nodes := parseJSON(input)
	return buildTree(nodes, p, q)
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

func buildTree(arr []*int, p, q int) (*TreeNode, *TreeNode, *TreeNode) {
	if len(arr) == 0 {
		return nil, nil, nil
	}

	if arr[0] == nil {
		return nil, nil, nil
	}

	var pN, qN *TreeNode
	root := &TreeNode{Val: *arr[0]}
	nodes := []*TreeNode{root}
	i := 1

	for i < len(arr) {
		curr := nodes[0]
		nodes = nodes[1:]

		if curr.Val == p {
			pN = curr
		} else if curr.Val == q {
			qN = curr
		}

		if i < len(arr) && arr[i] != nil {
			curr.Left = &TreeNode{Val: *arr[i]}
			nodes = append(nodes, curr.Left)
			if curr.Left.Val == p {
				pN = curr
			} else if curr.Left.Val == q {
				qN = curr
			}
		}
		i++

		if i < len(arr) && arr[i] != nil {
			curr.Right = &TreeNode{Val: *arr[i]}
			nodes = append(nodes, curr.Right)
			if curr.Right.Val == p {
				pN = curr
			} else if curr.Right.Val == q {
				qN = curr
			}
		}
		i++
	}

	return root, pN, qN
}
