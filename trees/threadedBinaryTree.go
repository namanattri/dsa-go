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
	dummyNode.rTag = 1
	dummyNode.left = dummyNode
	dummyNode.right = dummyNode

	return &ThreadedBinaryTree{root: dummyNode}
}

func (t *ThreadedBinaryTree) CreateInorderThreaded() {
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

func inOrderSuccessor(p *ThreadedBinaryTreeNode) *ThreadedBinaryTreeNode {
	if p.rTag == 0 {
		return p.right
	} else {
		position := p.right
		for position.lTag == 1 {
			position = position.left
		}
		return position
	}
}

func (t *ThreadedBinaryTree) InOrderTraversalInOrderTree() {
	p := inOrderSuccessor(t.root)
	for p != t.root {
		fmt.Printf("%d ", p.value)
		p = inOrderSuccessor(p)
	}
}

func (t *ThreadedBinaryTree) InOrderTraversalInOrderTreeAlt() {
	p := t.root
	for {
		p = inOrderSuccessor(p)

		if p == t.root {
			return
		}

		fmt.Printf("%d ", p.value)
	}
}

func preOrderSuccessor(p *ThreadedBinaryTreeNode) *ThreadedBinaryTreeNode {
	if p.lTag == 1 {
		return p.left
	} else {
		position := p
		for position.rTag == 0 {
			position = position.right
		}
		return position.right
	}
}

func (t *ThreadedBinaryTree) PreOrderTraversalInOrderTree() {
	p := preOrderSuccessor(t.root)

	for p != t.root {
		fmt.Printf("%d ", p.value)
		p = preOrderSuccessor(p)
	}
}

func (t *ThreadedBinaryTree) PreOrderTraversalInOrderTreeAlt() {
	p := t.root

	for {
		p = preOrderSuccessor(p)

		if p == t.root {
			return
		}

		fmt.Printf("%d ", p.value)
	}
}

func (t *ThreadedBinaryTree) InsertRightOfInOrderTree(p *ThreadedBinaryTreeNode, value int) {
	q := NewThreadedBinaryTreeNode(value)
	q.rTag = p.rTag
	q.right = p.right
	q.lTag = 0
	q.left = p
	p.rTag = 1
	p.right = q

	if q.rTag == 1 { // means p had a right child
		child := q.right

		// find the left most node starting from the right child and change the left pointer from p to q
		for child.lTag == 1 {
			child = child.left
		}
		child.left = q
	}
}
