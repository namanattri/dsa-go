package trees

import "fmt"

type GenericTreeNode struct {
	value       string
	firstChild  *GenericTreeNode
	nextSibling *GenericTreeNode
}

func NewGenericTreeNode(value string) *GenericTreeNode {
	return &GenericTreeNode{value: value}
}

func (n *GenericTreeNode) String() string {
	return n.value
}

type GenericTree struct {
	root *GenericTreeNode
}

func NewGenericTree() *GenericTree {
	return &GenericTree{}
}

func (t *GenericTree) Create() {
	t.root = NewGenericTreeNode("A")
	t.root.firstChild = NewGenericTreeNode("B")
	t.root.firstChild.nextSibling = NewGenericTreeNode("C")
	t.root.firstChild.nextSibling.nextSibling = NewGenericTreeNode("D")
	t.root.firstChild.nextSibling.nextSibling.nextSibling = NewGenericTreeNode("E")
	t.root.firstChild.nextSibling.nextSibling.nextSibling.nextSibling = NewGenericTreeNode("F")
	t.root.firstChild.nextSibling.nextSibling.nextSibling.nextSibling.nextSibling = NewGenericTreeNode("G")

	t.root.firstChild.nextSibling.nextSibling.firstChild = NewGenericTreeNode("H")

	t.root.firstChild.nextSibling.nextSibling.nextSibling.firstChild = NewGenericTreeNode("I")
	t.root.firstChild.nextSibling.nextSibling.nextSibling.firstChild.nextSibling = NewGenericTreeNode("J")

	t.root.firstChild.nextSibling.nextSibling.nextSibling.firstChild.nextSibling.firstChild = NewGenericTreeNode("P")
	t.root.firstChild.nextSibling.nextSibling.nextSibling.firstChild.nextSibling.firstChild.nextSibling = NewGenericTreeNode("Q")

	t.root.firstChild.nextSibling.nextSibling.nextSibling.nextSibling.firstChild = NewGenericTreeNode("K")
	t.root.firstChild.nextSibling.nextSibling.nextSibling.nextSibling.firstChild.nextSibling = NewGenericTreeNode("L")
	t.root.firstChild.nextSibling.nextSibling.nextSibling.nextSibling.firstChild.nextSibling.nextSibling = NewGenericTreeNode("M")

	t.root.firstChild.nextSibling.nextSibling.nextSibling.nextSibling.nextSibling.firstChild = NewGenericTreeNode("N")
}

func (t *GenericTree) LevelOrder() {
	root := t.root

	if root == nil {
		return
	}

	q := NewGenericQueue[*GenericTreeNode]()

	q.Enqueue(root)

	for !q.IsEmpty() {
		node, _ := q.Dequeue()

		fmt.Printf("%s ", node.value)

		if node.firstChild != nil {
			q.Enqueue(node.firstChild)
		}

		for node.nextSibling != nil {
			fmt.Printf("%s ", node.nextSibling.value)

			if node.nextSibling.firstChild != nil {
				q.Enqueue(node.nextSibling.firstChild)
			}

			node = node.nextSibling
		}
	}
}
