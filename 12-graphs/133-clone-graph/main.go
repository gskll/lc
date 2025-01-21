package main

// https://leetcode.com/problems/clone-graph/description/

// Given a reference of a node in a connected undirected graph.
//
// Return a deep copy (clone) of the graph.
//
// Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.
//
// class Node {
//     public int val;
//     public List<Node> neighbors;
// }
//
//
// Test case format:
//
// For simplicity, each node's value is the same as the node's index (1-indexed). For example, the first node with val == 1, the second node with val == 2, and so on. The graph is represented in the test case using an adjacency list.
//
// An adjacency list is a collection of unordered lists used to represent a finite graph. Each list describes the set of neighbors of a node in the graph.
//
// The given node will always be the first node with val = 1. You must return the copy of the given node as a reference to the cloned graph.
//
// The number of nodes in the graph is in the range [0, 100].
// 1 <= Node.val <= 100
// Node.val is unique for each node.
// There are no repeated edges and no self-loops in the graph.
// The Graph is connected and all nodes can be visited starting from the given node.

/*
	- max 101 nodes, val [1,100]
	- for each node, clone it, and clone neighbours, add neighbours to queue?

	- copy node, make copies of neighbours and store them in node copy
	- add neighbours to bfs queue if not visited
*/

/**
 * Definition for a Node.
 */
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	var clonedNodes [101]*Node
	clonedNodes[node.Val] = &Node{Val: node.Val}
	queue := []*Node{node}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		cloneCur := clonedNodes[cur.Val]

		for _, neighbour := range cur.Neighbors {
			if clonedNodes[neighbour.Val] == nil {
				clonedNodes[neighbour.Val] = &Node{Val: neighbour.Val}
				queue = append(queue, neighbour)
			}
			cloneCur.Neighbors = append(cloneCur.Neighbors, clonedNodes[neighbour.Val])
		}
	}

	return clonedNodes[node.Val]
}
