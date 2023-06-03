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

func (n *BinaryTreeNode) String() string {
	return fmt.Sprintf("%d", n.value)
}

type BinaryTree struct {
	root *BinaryTreeNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
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
	root := t.root
	for {
		for root != nil {
			fmt.Printf("%d ", root.value)
			s.Push(root)
			root = root.left
		}

		if s.IsEmptyStack() {
			break
		}

		poppedNode, _ := s.Pop()

		root = poppedNode.right
	}
}

func (t *BinaryTree) InOrder(node *BinaryTreeNode) {
	if node != nil {
		t.InOrder(node.left)
		fmt.Printf("%d ", node.value)
		t.InOrder(node.right)
	}
}

func (t *BinaryTree) InOrderNonRecursive() {
	s := NewGenericStack[*BinaryTreeNode]()
	root := t.root

	for {
		for root != nil {
			s.Push(root)
			root = root.left
		}

		if s.IsEmptyStack() {
			break
		}

		poppedNode, _ := s.Pop()

		fmt.Printf("%d ", poppedNode.value)

		root = poppedNode.right
	}
}

func (t *BinaryTree) PostOrder(node *BinaryTreeNode) {
	if node != nil {
		t.PostOrder(node.left)
		t.PostOrder(node.right)
		fmt.Printf("%d ", node.value)
	}
}

func (t *BinaryTree) PostOrderNonRecursive() {
	s := NewGenericStack[*BinaryTreeNode]()
	root := t.root
	previous := (*BinaryTreeNode)(nil)

	for {
		for root != nil {
			s.Push(root)
			root = root.left
		}

		for root == nil && !s.IsEmptyStack() {
			topNode, _ := s.Top()
			root = topNode

			if root.right == nil || root.right == previous {
				fmt.Printf("%d ", root.value)
				s.Pop()
				previous = root
				root = nil
			} else {
				root = root.right
			}
		}

		if s.IsEmptyStack() {
			break
		}
	}
}

func (t *BinaryTree) LevelOrder() {
	if t.root == nil {
		return
	}

	q := NewGenericQueue[*BinaryTreeNode]()

	q.Enqueue(t.root)

	for !q.IsEmpty() {
		node, _ := q.Dequeue()

		fmt.Printf("%d ", node.value)

		if node.left != nil {
			q.Enqueue(node.left)
		}

		if node.right != nil {
			q.Enqueue(node.right)
		}
	}
}
