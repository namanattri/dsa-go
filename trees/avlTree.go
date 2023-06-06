package trees

type AVLTreeNode struct {
	value  int
	left   *AVLTreeNode
	right  *AVLTreeNode
	height int
}

func NewAVLTreeNode(value int) *AVLTreeNode {
	return &AVLTreeNode{value: value}
}

type AVLTree struct {
	root *AVLTreeNode
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

func (t *AVLTree) Create() {
	t.root = NewAVLTreeNode(6)
	t.root.height = max(height(t.root.left), height(t.root.right)) + 1

	t.root.left = NewAVLTreeNode(5)
	t.root.left.height = max(height(t.root.left.left), height(t.root.right.right)) + 1
	t.root.right = NewAVLTreeNode(9)
	t.root.right.height = max(height(t.root.right.left), height(t.root.right.right)) + 1

	t.root.left.left = NewAVLTreeNode(5)
	t.root.left.left.height = 0
	t.root.right.left = NewAVLTreeNode(8)
	t.root.right.left.height = 0
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
