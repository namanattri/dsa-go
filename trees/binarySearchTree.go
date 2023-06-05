package trees

import "fmt"

type BinarySearchTreeNode struct {
	value int
	left  *BinarySearchTreeNode
	right *BinarySearchTreeNode
}

func NewBinarySearchTreeNode(value int) *BinarySearchTreeNode {
	return &BinarySearchTreeNode{value: value}
}

func (n *BinarySearchTreeNode) String() string {
	return fmt.Sprintf("%d ", n.value)
}

type BinarySearchTree struct {
	root *BinarySearchTreeNode
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (t *BinarySearchTree) Create() {
	t.root = NewBinarySearchTreeNode(10)
	t.root.left = NewBinarySearchTreeNode(6)
	t.root.right = NewBinarySearchTreeNode(16)

	t.root.left.left = NewBinarySearchTreeNode(4)
	t.root.left.right = NewBinarySearchTreeNode(9)

	t.root.right.left = NewBinarySearchTreeNode(13)

	t.root.left.right.left = NewBinarySearchTreeNode(7)
}

func (t *BinarySearchTree) FindRecursive(root *BinarySearchTreeNode, value int) *BinarySearchTreeNode {
	if root == nil {
		return nil
	}

	if value < root.value {
		return t.FindRecursive(root.left, value)
	} else if value > root.value {
		return t.FindRecursive(root.right, value)
	} else {
		return root
	}
}

func (t *BinarySearchTree) FindNonRecursive(value int) *BinarySearchTreeNode {
	root := t.root

	for root != nil {
		if value == root.value {
			return root
		} else if value < root.value {
			root = root.left
		} else {
			root = root.right
		}
	}

	return nil
}

func (t *BinarySearchTree) FindMinimumRecursive(root *BinarySearchTreeNode) *BinarySearchTreeNode {
	if root == nil {
		return nil
	}

	if root.left == nil {
		return root
	} else {
		return t.FindMinimumRecursive(root.left)
	}
}
