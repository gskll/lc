package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
//
// Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.
//
// 1 <= preorder.length <= 3000
// inorder.length == preorder.length
// -3000 <= preorder[i], inorder[i] <= 3000
// preorder and inorder consist of unique values.
// Each value of inorder also appears in preorder.
// preorder is guaranteed to be the preorder traversal of the tree.
// inorder is guaranteed to be the inorder traversal of the tree.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// optimal solution - O(n)/O(n) for map/recursion stack
func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderIndices := make(map[int]int, len(inorder))
	for i, x := range inorder {
		inorderIndices[x] = i
	}

	var build func(preStart, preEnd, inStart, inEnd int) *TreeNode
	build = func(preStart, preEnd, inStart, inEnd int) *TreeNode {
		if preStart >= preEnd || inStart >= inEnd {
			return nil
		}

		node := &TreeNode{Val: preorder[preStart]}
		rootIndex := inorderIndices[node.Val]
		leftSize := rootIndex - inStart

		node.Left = build(preStart+1, preStart+1+leftSize, inStart, rootIndex)
		node.Right = build(preStart+1+leftSize, preEnd, rootIndex+1, inEnd)

		return node
	}

	return build(0, len(preorder), 0, len(inorder))
}

// second solution - fixed slices indices
func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	rootIndex := 0
	for i, val := range inorder {
		if val == preorder[0] {
			rootIndex = i
			break
		}
	}

	root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])

	return root
}

// first solution: passes tests but buggy
func buildTree1(preorder []int, inorder []int) *TreeNode {
	fmt.Println(preorder, inorder)

	preorderI, inorderI := 0, 0
	root := &TreeNode{Val: preorder[0]}

	for preorder[preorderI] != inorder[inorderI] {
		inorderI++
	}

	fmt.Println("start -> preorderI: ", preorderI, "; inorderI: ", inorderI)

	if inorderI > 0 {
		preorderI++
		fmt.Println("left -> preorderI: ", preorderI, "; inorderI: ", inorderI)
		root.Left = buildTree(preorder[preorderI:], inorder[:inorderI])
	}

	if inorderI < len(inorder)-1 {
		preorderI += max(inorderI, 1)
		fmt.Println("right -> preorderI: ", preorderI, "; inorderI: ", inorderI)
		// can't just increment - needs to be 1+index of last node in left tree
		root.Right = buildTree(preorder[preorderI:], inorder[inorderI+1:])
	}

	return root
}

func main() {
	tests := []struct {
		pre     string
		inorder string
		out     string
	}{
		{
			pre:     "[3,9,20,15,7]",
			inorder: "[9,3,15,20,7]",
			out:     "[3,9,20,null,null,15,7]",
		},
		{
			pre:     "[-1]",
			inorder: "[-1]",
			out:     "[-1]",
		},
		{
			pre:     "[3,1,2,4]",
			inorder: "[1,2,3,4]",
			out:     "[3,1,4,null,2]",
		},
	}
	for _, test := range tests {
		start := time.Now()
		var pre []int
		var inorder []int
		json.Unmarshal([]byte(test.pre), &pre)
		json.Unmarshal([]byte(test.inorder), &inorder)
		r := buildTree(pre, inorder)
		fmt.Printf("elapsed: %s\n", time.Since(start))
		fmt.Printf("got: %+v\n", r.ToLeetCode())
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
