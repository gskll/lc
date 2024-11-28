package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/description/

// Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.
//
// Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.
//
// Clarification: The input/output format is the same as how LeetCode serializes a binary tree. You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.
//
// The number of nodes in the tree is in the range [0, 10^4].
// -1000 <= Node.val <= 1000

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	null      string
	separator string
}

func Constructor() Codec {
	return Codec{
		null:      "null",
		separator: ",",
	}
}

// dfs solution - more efficient than bfs
// for optimal could use dfs for skewed tree and bfs for balanced
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var ser []string
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			ser = append(ser, this.null)
			return
		}

		ser = append(ser, strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	for i := len(ser) - 1; i >= 0; i-- {
		if ser[i] != this.null {
			ser = ser[:i+1]
			break
		}
	}

	return strings.Join(ser, this.separator)
}

func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" || data == this.null {
		return nil
	}

	vals := strings.Split(data, this.separator)
	i := 0

	var helper func() *TreeNode
	helper = func() *TreeNode {
		if i >= len(vals) || vals[i] == this.null {
			i++
			return nil
		}

		val, _ := strconv.Atoi(vals[i])
		node := &TreeNode{Val: val}
		i++

		node.Left = helper()
		node.Right = helper()

		return node
	}

	return helper()
}

/// LEETCODE BFS SERIALIZATION

// Serializes a tree to a single string.
// level by level with bfs
func (this *Codec) serializeBFS(root *TreeNode) string {
	if root == nil {
		return "[]"
	}

	var ser []string

	queue := []*TreeNode{root}

	currLevel := 1
	for len(queue) > 0 {
		levelCount := len(queue)
		for nodeInLevel := 0; nodeInLevel < levelCount; nodeInLevel++ {
			node := queue[0]
			queue = queue[1:]

			if node == nil {
				ser = append(ser, this.null)
				continue
			}

			ser = append(ser, strconv.Itoa(node.Val))

			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
		currLevel++
	}

	for i := len(ser) - 1; i >= 0; i-- {
		if ser[i] != this.null {
			ser = ser[:i+1]
			break
		}
	}

	return "[" + strings.Join(ser, this.separator) + "]"
}

// Deserializes your encoded data to tree.
func (this *Codec) deserializeBFS(data string) *TreeNode {
	if data == "[]" {
		return nil
	}
	data = strings.Trim(data, "[]")
	vals := strings.Split(data, this.separator)

	if len(vals) == 0 || vals[0] == this.null {
		return nil
	}

	rootValue, _ := strconv.Atoi(vals[0])
	root := &TreeNode{Val: rootValue}
	nodes := []*TreeNode{root}
	i := 1

	for i < len(vals) {
		curr := nodes[0]
		nodes = nodes[1:]

		if i < len(vals) && vals[i] != this.null {
			val, _ := strconv.Atoi(vals[i])
			curr.Left = &TreeNode{Val: val}
			nodes = append(nodes, curr.Left)
		}
		i++

		if i < len(vals) && vals[i] != this.null {
			val, _ := strconv.Atoi(vals[i])
			curr.Right = &TreeNode{Val: val}
			nodes = append(nodes, curr.Right)
		}
		i++
	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

func main() {
	tests := []struct {
		root string
	}{
		{
			root: "[1,2,3,null,null,4,5]",
		},
		{
			root: "[]",
		},
	}
	for _, test := range tests {
		start := time.Now()
		ser := Constructor()
		deser := Constructor()
		root := TreeFromLeetCode(test.root)

		serialized := ser.serialize(root)

		fmt.Printf("ser: elapsed: %s\n", time.Since(start))
		fmt.Printf("got: %+v\n", serialized)
		fmt.Printf("exp: %+v\n", test.root)

		deserialized := deser.deserialize(serialized)

		fmt.Printf("deser: elapsed: %s\n", time.Since(start))
		fmt.Printf("got: %+v\n", deserialized.ToLeetCode())
		fmt.Printf("exp: %+v\n", test.root)
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
