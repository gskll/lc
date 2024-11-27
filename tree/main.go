package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Insert(root *TreeNode, val int) *TreeNode {
	if exists := Find(root, val); exists != nil {
		return nil
	}

	return insertNode(root, val)
}

func insertNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if root.Val > val {
		root.Left = Insert(root.Left, val)
	} else {
		root.Right = Insert(root.Right, val)
	}

	return root
}

// no child: delete
// 1 child: delete and replace with child
// 2 child: successor is smallest in right tree
func Delete(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > val {
		root.Left = Delete(root.Left, val)
	} else if root.Val < val {
		root.Right = Delete(root.Right, val)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		successor := findMin(root.Right)
		root.Val = successor.Val
		root.Right = Delete(root.Right, successor.Val)
	}
	return root
}

func Find(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if root.Val > val {
		return Find(root.Left, val)
	} else {
		return Find(root.Right, val)
	}
}

func findMin(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func findMax(root *TreeNode) *TreeNode {
	for root.Right != nil {
		root = root.Right
	}
	return root
}

func String(root *TreeNode) string {
	var result strings.Builder

	var preorder func(node *TreeNode, level int)
	preorder = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		indent := strings.Repeat(" ", level)
		result.WriteString(fmt.Sprintf("%s(%d - %p)\n", indent, node.Val, node))
		if node.Left != nil {
			preorder(node.Left, level+1)
		}

		if node.Right != nil {
			preorder(node.Right, level+1)
		}
	}

	preorder(root, 0)

	return result.String()
}

func IsEmpty(root *TreeNode) bool {
	return root == nil
}

func Height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftMax := Height(root.Left)
	rightMax := Height(root.Right)

	return max(leftMax, rightMax) + 1
}

func Size(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + Size(root.Left) + Size(root.Right)
}
