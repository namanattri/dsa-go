package trees

import "fmt"

type ThreadedBinaryTreeNode struct {
	left  *ThreadedBinaryTreeNode
	lTag  int
	value int
	rTag  int
	right *ThreadedBinaryTreeNode
}

func NewThreadedBinaryTreeNode(value int) *ThreadedBinaryTreeNode {
	return &ThreadedBinaryTreeNode{value: value}
}

func (n *ThreadedBinaryTreeNode) String() string {
	return fmt.Sprintf("%d", n.value)
}

type ThreadedBinaryTree struct {
	root *ThreadedBinaryTreeNode
}

func NewThreadedBinaryTree() *ThreadedBinaryTree {
	// create a dummy node
	dummyNode := NewThreadedBinaryTreeNode(0)
	dummyNode.lTag = 0
	dummyNode.rTag = 0
	dummyNode.left = dummyNode
	dummyNode.right = dummyNode

	return &ThreadedBinaryTree{root: dummyNode}
}

func (t *ThreadedBinaryTree) Create() {
	t.root.lTag = 1
	t.root.left = NewThreadedBinaryTreeNode(1)

	t.root.left.lTag = 1
	t.root.left.left = NewThreadedBinaryTreeNode(5)

	t.root.left.left.lTag = 1
	t.root.left.left.left = NewThreadedBinaryTreeNode(2)

	t.root.left.left.rTag = 0
	t.root.left.left.right = t.root.left

	t.root.left.left.left.lTag = 0
	t.root.left.left.left.left = t.root
	t.root.left.left.left.rTag = 0
	t.root.left.left.left.right = t.root.left.left

	t.root.left.rTag = 1
	t.root.left.right = NewThreadedBinaryTreeNode(11)

	t.root.left.right.lTag = 1
	t.root.left.right.left = NewThreadedBinaryTreeNode(16)
	t.root.left.right.rTag = 1
	t.root.left.right.right = NewThreadedBinaryTreeNode(31)

	t.root.left.right.left.lTag = 0
	t.root.left.right.left.left = t.root.left
	t.root.left.right.left.rTag = 0
	t.root.left.right.left.right = t.root.left.right

	t.root.left.right.right.lTag = 0
	t.root.left.right.right.left = t.root.left.right
	t.root.left.right.right.rTag = 0
	t.root.left.right.right.right = t.root
}
