package rbtree

import "fmt"

/*
	RED BLACK TREE

	- uses extra bit on each node to mark red or black
	- ensure tree remains relatively balanced as we insert/delete
	- insert/delete and rotation/recoloring are all O(logn)

	Rules for red/black tree
	- the root is black
	- each node is either red or black
	- all leaf nodes NIL are black
	- if a node is red, both children are black
	- all paths from a node to any NIL leaf node pass through the same number of black nodes

	- tree is kept balanced by rotations, whenever a branch is getting too long it gets rotated to keep the tree shallow
	- a properly ordered tree pre-rotations stays properly ordered post-rotation
	- rotations are O(1)
	- LEFT ROTATION:
		- pivot node's initial parent becomes it's left child
		- pivot node's old left child becomes initial parent's new right child
*/

// Color represents the color of a Red-Black Tree node
type Color bool

const (
	RED   Color = true
	BLACK Color = false
)

// Node represents a node in the Red-Black Tree
type Node struct {
	Color  Color
	Parent *Node
	Left   *Node
	Right  *Node
	Val    int
}

// Tree represents a Red-Black Tree
type Tree struct {
	Root *Node
	NIL  *Node // Sentinel node
	size int
}

// NewTree creates a new Red-Black Tree
func NewTree() *Tree {
	nil_node := &Node{Color: BLACK}
	return &Tree{
		NIL:  nil_node,
		Root: nil_node,
		size: 0,
	}
}

// Size returns the number of nodes in the tree
func (t *Tree) Size() int {
	return t.size
}

// Insert adds a new value to the tree
func (t *Tree) Insert(val int) {
	node := &Node{
		Val:    val,
		Color:  RED,
		Left:   t.NIL,
		Right:  t.NIL,
		Parent: nil,
	}

	var parent *Node = nil
	current := t.Root

	// Find insertion point
	for current != t.NIL {
		parent = current
		if val < current.Val {
			current = current.Left
		} else if val > current.Val {
			current = current.Right
		} else {
			return // Duplicate value
		}
	}

	// Set up node connections
	node.Parent = parent
	if parent == nil {
		t.Root = node
	} else if val < parent.Val {
		parent.Left = node
	} else {
		parent.Right = node
	}

	t.size++
	t.fixInsert(node)
}

// fixInsert restores Red-Black Tree properties after insertion
func (t *Tree) fixInsert(node *Node) {
	for node.Parent != nil && node.Parent.Color == RED {
		if node.Parent == node.Parent.Parent.Left {
			uncle := node.Parent.Parent.Right

			if uncle.Color == RED {
				// Case 1: Uncle is red
				node.Parent.Color = BLACK
				uncle.Color = BLACK
				node.Parent.Parent.Color = RED
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Right {
					// Case 2: Uncle is black and node is a right child
					node = node.Parent
					t.leftRotate(node)
				}
				// Case 3: Uncle is black and node is a left child
				node.Parent.Color = BLACK
				node.Parent.Parent.Color = RED
				t.rightRotate(node.Parent.Parent)
			}
		} else {
			uncle := node.Parent.Parent.Left

			if uncle.Color == RED {
				// Case 1: Uncle is red
				node.Parent.Color = BLACK
				uncle.Color = BLACK
				node.Parent.Parent.Color = RED
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Left {
					// Case 2: Uncle is black and node is a left child
					node = node.Parent
					t.rightRotate(node)
				}
				// Case 3: Uncle is black and node is a right child
				node.Parent.Color = BLACK
				node.Parent.Parent.Color = RED
				t.leftRotate(node.Parent.Parent)
			}
		}
	}
	t.Root.Color = BLACK
}

// leftRotate performs a left rotation around the given node
func (t *Tree) leftRotate(x *Node) {
	y := x.Right
	x.Right = y.Left

	if y.Left != t.NIL {
		y.Left.Parent = x
	}

	y.Parent = x.Parent

	if x.Parent == nil {
		t.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}

	y.Left = x
	x.Parent = y
}

// rightRotate performs a right rotation around the given node
func (t *Tree) rightRotate(x *Node) {
	y := x.Left
	x.Left = y.Right

	if y.Right != t.NIL {
		y.Right.Parent = x
	}

	y.Parent = x.Parent

	if x.Parent == nil {
		t.Root = y
	} else if x == x.Parent.Right {
		x.Parent.Right = y
	} else {
		x.Parent.Left = y
	}

	y.Right = x
	x.Parent = y
}

// Delete removes a value from the tree
func (t *Tree) Delete(val int) bool {
	z := t.search(val)
	if z == t.NIL {
		return false
	}

	t.delete(z)
	t.size--
	return true
}

// search finds a node with the given value
func (t *Tree) search(val int) *Node {
	current := t.Root
	for current != t.NIL {
		if val < current.Val {
			current = current.Left
		} else if val > current.Val {
			current = current.Right
		} else {
			return current
		}
	}
	return t.NIL
}

// delete removes the given node from the tree
func (t *Tree) delete(z *Node) {
	y := z
	yOriginalColor := y.Color
	var x *Node

	if z.Left == t.NIL {
		x = z.Right
		t.transplant(z, z.Right)
	} else if z.Right == t.NIL {
		x = z.Left
		t.transplant(z, z.Left)
	} else {
		y = t.minimum(z.Right)
		yOriginalColor = y.Color
		x = y.Right

		if y.Parent == z {
			x.Parent = y
		} else {
			t.transplant(y, y.Right)
			y.Right = z.Right
			y.Right.Parent = y
		}

		t.transplant(z, y)
		y.Left = z.Left
		y.Left.Parent = y
		y.Color = z.Color
	}

	if yOriginalColor == BLACK {
		t.fixDelete(x)
	}
}

// minimum finds the minimum value in the subtree rooted at x
func (t *Tree) minimum(x *Node) *Node {
	for x.Left != t.NIL {
		x = x.Left
	}
	return x
}

// transplant replaces subtree rooted at u with subtree rooted at v
func (t *Tree) transplant(u, v *Node) {
	if u.Parent == nil {
		t.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	v.Parent = u.Parent
}

// fixDelete restores Red-Black Tree properties after deletion
func (t *Tree) fixDelete(x *Node) {
	for x != t.Root && x.Color == BLACK {
		if x == x.Parent.Left {
			w := x.Parent.Right

			if w.Color == RED {
				w.Color = BLACK
				x.Parent.Color = RED
				t.leftRotate(x.Parent)
				w = x.Parent.Right
			}

			if w.Left.Color == BLACK && w.Right.Color == BLACK {
				w.Color = RED
				x = x.Parent
			} else {
				if w.Right.Color == BLACK {
					w.Left.Color = BLACK
					w.Color = RED
					t.rightRotate(w)
					w = x.Parent.Right
				}

				w.Color = x.Parent.Color
				x.Parent.Color = BLACK
				w.Right.Color = BLACK
				t.leftRotate(x.Parent)
				x = t.Root
			}
		} else {
			w := x.Parent.Left

			if w.Color == RED {
				w.Color = BLACK
				x.Parent.Color = RED
				t.rightRotate(x.Parent)
				w = x.Parent.Left
			}

			if w.Right.Color == BLACK && w.Left.Color == BLACK {
				w.Color = RED
				x = x.Parent
			} else {
				if w.Left.Color == BLACK {
					w.Right.Color = BLACK
					w.Color = RED
					t.leftRotate(w)
					w = x.Parent.Left
				}

				w.Color = x.Parent.Color
				x.Parent.Color = BLACK
				w.Left.Color = BLACK
				t.rightRotate(x.Parent)
				x = t.Root
			}
		}
	}
	x.Color = BLACK
}

// Contains checks if a value exists in the tree
func (t *Tree) Contains(val int) bool {
	return t.search(val) != t.NIL
}

// InOrderTraversal performs an in-order traversal of the tree
func (t *Tree) InOrderTraversal(fn func(int)) {
	var inorder func(*Node)
	inorder = func(node *Node) {
		if node != t.NIL {
			inorder(node.Left)
			fn(node.Val)
			inorder(node.Right)
		}
	}
	inorder(t.Root)
}

// ValidateProperties checks if the tree maintains Red-Black properties
func (t *Tree) ValidateProperties() error {
	if t.Root.Color != BLACK {
		return fmt.Errorf("property violation: root must be black")
	}

	// Check all properties recursively
	blackHeight, err := t.validateNode(t.Root, 0)
	if err != nil {
		return err
	}

	// Verify black height
	if _, err := t.verifyBlackHeight(t.Root, 0, blackHeight); err != nil {
		return err
	}

	return nil
}

func (t *Tree) validateNode(node *Node, blackCount int) (int, error) {
	if node == t.NIL {
		return blackCount, nil
	}

	// Property: Red nodes should have black children
	if node.Color == RED {
		if node.Left.Color == RED || node.Right.Color == RED {
			return 0, fmt.Errorf("property violation: red node has red child")
		}
	}

	// Count black nodes
	if node.Color == BLACK {
		blackCount++
	}

	// Recursively validate children
	leftCount, err := t.validateNode(node.Left, blackCount)
	if err != nil {
		return 0, err
	}

	rightCount, err := t.validateNode(node.Right, blackCount)
	if err != nil {
		return 0, err
	}

	// Verify equal black height
	if leftCount != rightCount {
		return 0, fmt.Errorf("property violation: unequal black height")
	}

	return leftCount, nil
}

func (t *Tree) verifyBlackHeight(node *Node, height int, expectedHeight int) (int, error) {
	if node == t.NIL {
		if height != expectedHeight {
			return 0, fmt.Errorf("property violation: inconsistent black height")
		}
		return height, nil
	}

	if node.Color == BLACK {
		height++
	}

	if _, err := t.verifyBlackHeight(node.Left, height, expectedHeight); err != nil {
		return 0, err
	}

	if _, err := t.verifyBlackHeight(node.Right, height, expectedHeight); err != nil {
		return 0, err
	}

	return height, nil
}
