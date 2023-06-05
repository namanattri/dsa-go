package trees

import "fmt"

type NodeType interface {
	int | *BinaryTreeNode | *GenericTreeNode | *ExpressionTreeNode | *BinarySearchTreeNode
}

type GenericLinkedListNode[T NodeType] struct {
	value T
	next  *GenericLinkedListNode[T]
}

func NewGenericLinkedListNode[T NodeType](value T) *GenericLinkedListNode[T] {
	return &GenericLinkedListNode[T]{value: value}
}

type GenericStack[T NodeType] struct {
	head *GenericLinkedListNode[T]
}

func NewGenericStack[T NodeType]() *GenericStack[T] {
	return &GenericStack[T]{}
}

func (s *GenericStack[T]) IsEmpty() bool {
	return s.head == nil
}

func (s *GenericStack[T]) Push(value T) {
	node := NewGenericLinkedListNode(value)

	if s.head == nil {
		s.head = node
	} else {
		node.next = s.head
		s.head = node
	}
}

func (s *GenericStack[T]) Top() (T, error) {
	if s.IsEmpty() {
		err := fmt.Errorf("error: stack underflow")
		var res T // in order to return the default value declare a variable of type T
		return res, err
	}

	return s.head.value, nil
}

func (s *GenericStack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		err := fmt.Errorf("error: stack underflow")
		var res T // in order to return the default value declare a variable of type T
		return res, err
	}

	value := s.head.value
	s.head = s.head.next

	return value, nil
}

func (s *GenericStack[T]) String() string {
	str := "["
	cursor := s.head

	for cursor != nil {
		str += fmt.Sprintf("%v -> ", cursor.value)
		cursor = cursor.next
	}

	str += "X]"

	return str
}
