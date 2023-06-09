package trees

import "fmt"

type AVLTreeNode struct {
	value  int
	left   *AVLTreeNode
	right  *AVLTreeNode
	height int
}

func NewAVLTreeNode(value int) *AVLTreeNode {
	return &AVLTreeNode{value: value}
}

func (n *AVLTreeNode) String() string {
	return fmt.Sprintf("%d ", n.value)
}

func (n *AVLTreeNode) SingleRotateLeft() *AVLTreeNode {
	leftChild := n.left
	n.left = leftChild.right
	leftChild.right = n
	leftChild.height = max(height(leftChild.left), height(leftChild.right)) + 1
	n.height = max(height(n.left), height(n.right)) + 1
	return leftChild
}

func (n *AVLTreeNode) SingleRotateRight() *AVLTreeNode {
	rightChild := n.right
	n.right = rightChild.left
	rightChild.left = n
	rightChild.height = max(height(rightChild.left), height(rightChild.right)) + 1
	n.height = max(height(n.left), height(n.right)) + 1
	return rightChild
}

func (n *AVLTreeNode) DoubleRotateLeft() *AVLTreeNode {
	n.left = n.left.SingleRotateRight()
	return n.SingleRotateLeft()
}

func (n *AVLTreeNode) DoubleRotateRight() *AVLTreeNode {
	n.right = n.right.SingleRotateLeft()
	return n.SingleRotateRight()
}

type AVLTree struct {
	root *AVLTreeNode
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

func (t *AVLTree) Insert(root *AVLTreeNode, value int) *AVLTreeNode {
	if root == nil {
		root = NewAVLTreeNode(value)
	} else if value < root.value {
		root.left = t.Insert(root.left, value)
		// check for imbalance
		if height(root.left)-height(root.right) == 2 {
			if value < root.left.value { // insertion happened in left subtree of left child
				root = root.SingleRotateLeft()
			} else { // insertion happened in right subtree of left child
				root = root.DoubleRotateLeft()
			}
		}
	} else if value > root.value {
		root.right = t.Insert(root.right, value)
		// check for imbalance
		if height(root.right)-height(root.left) == 2 {
			if value > root.right.value { // insertion happened in right subtree of right child
				root = root.SingleRotateRight()
			} else { // insertion happened in left subtree of right child
				root = root.DoubleRotateRight()
			}
		}
	} // else value == root.value which mean the value already exists in the avl tree

	root.height = max(height(root.left), height(root.right)) + 1
	return root
}

func (t *AVLTree) InOrder(node *AVLTreeNode) {
	if node != nil {
		t.InOrder(node.left)
		fmt.Printf("%d ", node.value)
		t.InOrder(node.right)
	}
}

func (t *AVLTree) LevelOrder() {
	if t.root == nil {
		return
	}

	q := NewGenericQueue[*AVLTreeNode]()

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

func height(n *AVLTreeNode) int {
	if n == nil {
		return -1
	}
	return n.height
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
