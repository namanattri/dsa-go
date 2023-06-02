package trees

import (
	"fmt"
)

type BinaryTreeNode struct {
	value int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func NewBinaryTreeNode(value int) *BinaryTreeNode {
	return &BinaryTreeNode{value: value}
}

type BinaryTree struct {
	root *BinaryTreeNode
}

func (t *BinaryTree) Create() {
	t.root = NewBinaryTreeNode(1)

	t.root.left = NewBinaryTreeNode(2)
	t.root.left.left = NewBinaryTreeNode(4)
	t.root.left.right = NewBinaryTreeNode(5)

	t.root.right = NewBinaryTreeNode(3)
	t.root.right.left = NewBinaryTreeNode(6)
	t.root.right.right = NewBinaryTreeNode(7)
}

func (t *BinaryTree) PreOrder(node *BinaryTreeNode) {
	if node != nil {
		fmt.Printf("%d ", node.value)
		t.PreOrder(node.left)
		t.PreOrder(node.right)
	}
}

func (t *BinaryTree) PreOrderNonRecursive() {
	s := NewGenericStack[*BinaryTreeNode]()
	node := t.root
	for {
		for node != nil {
			fmt.Printf("%d ", node.value)
			s.Push(node)
			node = node.left
		}

		if s.IsEmptyStack() {
			break
		}

		node, _ := s.Pop()

		node = node.right
	}
}
