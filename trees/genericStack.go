package trees

import "fmt"

type NodeType interface {
	int | *BinaryTreeNode
}

type GenericStackNode[T NodeType] struct {
	value T
	next  *GenericStackNode[T]
}

func NewGenericStackNode[T NodeType](value T) *GenericStackNode[T] {
	return &GenericStackNode[T]{value: value}
}

type GenericStack[T NodeType] struct {
	head *GenericStackNode[T]
}

func NewGenericStack[T NodeType]() *GenericStack[T] {
	return &GenericStack[T]{}
}

func (s *GenericStack[T]) IsEmptyStack() bool {
	return s.head == nil
}

func (s *GenericStack[T]) Push(value T) {
	node := NewGenericStackNode(value)

	if s.head == nil {
		s.head = node
	} else {
		node.next = s.head
		s.head = node
	}
}

func (s *GenericStack[T]) Pop() (T, error) {
	if s.IsEmptyStack() {
		err := fmt.Errorf("error: stack underflow")
		var res T
		return res, err
	}

	value := s.head.value
	s.head = s.head.next

	return value, nil
}
