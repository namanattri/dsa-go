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

func (t *BinarySearchTree) FindMinimumNonRecursive() *BinarySearchTreeNode {
	if t.root == nil {
		return nil
	}

	root := t.root

	for root.left != nil {
		root = root.left
	}

	return root
}

func (t *BinarySearchTree) FindMaximumRecursive(root *BinarySearchTreeNode) *BinarySearchTreeNode {
	if root == nil {
		return nil
	}

	if root.right == nil {
		return root
	} else {
		return t.FindMaximumRecursive(root.right)
	}
}

func (t *BinarySearchTree) FindMaximumNonRecursive() *BinarySearchTreeNode {
	if t.root == nil {
		return nil
	}

	root := t.root

	for root.right != nil {
		root = root.right
	}

	return root
}

func (t *BinarySearchTree) TraverseInorderRecursive(root *BinarySearchTreeNode) {
	if root == nil {
		return
	}
	t.TraverseInorderRecursive(root.left)
	fmt.Print(root)
	t.TraverseInorderRecursive(root.right)
}

func (t *BinarySearchTree) TraverseInorderNonRecursive() {
	s := NewGenericStack[*BinarySearchTreeNode]()
	root := t.root

	for {
		for root != nil {
			s.Push(root)
			root = root.left
		}

		if s.IsEmpty() {
			break
		}

		popped, _ := s.Pop()

		fmt.Print(popped)

		root = popped.right
	}
}

func (t *BinarySearchTree) InsertRecursive(root *BinarySearchTreeNode, value int) *BinarySearchTreeNode {
	if root == nil {
		root = NewBinarySearchTreeNode(value)
	} else if value < root.value {
		root.left = t.InsertRecursive(root.left, value)
	} else if value > root.value {
		root.right = t.InsertRecursive(root.right, value)
	} // if value is same as root do nothing
	return root
}

func (t *BinarySearchTree) InsertNonRecursive(value int) {
	root := t.root

	for root != nil {
		if value == root.value { // do nothing
			return
		} else if value < root.value {
			if root.left == nil {
				root.left = NewBinarySearchTreeNode(value)
				return
			} else {
				root = root.left
			}
		} else {
			if root.right == nil {
				root.right = NewBinarySearchTreeNode(value)
				return
			} else {
				root = root.right
			}
		}
	}
}
