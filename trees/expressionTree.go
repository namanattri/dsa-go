package trees

import (
	"fmt"
	"unicode"
)

type ExpressionTreeNode struct {
	value rune
	left  *ExpressionTreeNode
	right *ExpressionTreeNode
}

func NewExpressionTreeNode(value rune) *ExpressionTreeNode {
	return &ExpressionTreeNode{value: value}
}

func (n *ExpressionTreeNode) String() string {
	return fmt.Sprintf("%c ", n.value)
}

type ExpressionTree struct {
	root *ExpressionTreeNode
}

func NewExpressionTree() *ExpressionTree {
	return &ExpressionTree{}
}

// postfix expression is of the type B C * A + D / which is equivalent to ((B + C) * A)/D
func (t *ExpressionTree) CreateFromPostfixExpression(postfixExpression []rune) {
	s := NewGenericStack[*ExpressionTreeNode]()
	var r *ExpressionTreeNode
	for _, value := range postfixExpression {
		if unicode.IsUpper(value) { // value is an operand
			// push to stack
			s.Push(NewExpressionTreeNode(value))
		} else { // value is an operator
			//create a tree with operator as root and top two nodes in stack as right and left child of the tree
			r = NewExpressionTreeNode(value)
			r.right, _ = s.Pop()
			r.left, _ = s.Pop()
			s.Push(r)
		}
	}
	t.root = r
}

func (t *ExpressionTree) InOrderNonRecursive() {
	s := NewGenericStack[*ExpressionTreeNode]()
	root := t.root

	for {
		for root != nil {
			s.Push(root)
			root = root.left
		}

		if s.IsEmptyStack() {
			break
		}

		popped, _ := s.Pop()

		fmt.Printf("%c ", popped.value)

		root = popped.right
	}
}

func (t *ExpressionTree) PreOrderNonRecursive() {
	s := NewGenericStack[*ExpressionTreeNode]()
	root := t.root

	for {
		for root != nil {
			fmt.Printf("%c ", root.value)
			s.Push(root)
			root = root.left
		}

		if s.IsEmptyStack() {
			break
		}

		popped, _ := s.Pop()

		root = popped.right
	}
}

func (t *ExpressionTree) PostOrderNonRecursive() {
	s := NewGenericStack[*ExpressionTreeNode]()
	root := t.root
	var previous *ExpressionTreeNode

	for {
		for root != nil {
			s.Push(root)
			root = root.left
		}

		for root == nil && !s.IsEmptyStack() {
			top, _ := s.Top()

			if top.right == nil || top.right == previous {
				fmt.Printf("%c ", top.value)
				s.Pop()
				previous = top
				root = nil
			} else {
				root = top.right
			}
		}

		if s.IsEmptyStack() {
			break
		}
	}
}
