package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := depth(node.Left)
		if left == -1 {
			return -1
		}
		right := depth(node.Right)
		if right == -1 {
			return -1
		}

		if left-right > 1 || right-left > 1 {
			return -1
		}

		return max(left, right) + 1
	}

	return depth(root) != -1
}

func isBalanced2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	balanced := true
	counter := 0

	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		counter++
		if node == nil || !balanced {
			return 0
		}

		left := depth(node.Left)
		right := depth(node.Right)

		if math.Abs(float64(left-right)) > 1 {
			balanced = false
		}

		return max(left, right) + 1
	}

	depth(root)
	fmt.Println("counter: ", counter)

	return balanced
}

func main() {
	tests := []struct {
		in  string
		out bool
	}{
		// {
		// 	in:  "[3,9,20,null,null,15,7]",
		// 	out: true,
		// },
		{
			in:  "[1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4,1,2,2,3,3,null,null,4,4]",
			out: false,
		},
		// {
		// 	in:  "[]",
		// 	out: true,
		// },
	}
	for _, test := range tests {
		t := TreeFromLeetCode(test.in)
		start := time.Now()
		r := isBalanced(t)
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
