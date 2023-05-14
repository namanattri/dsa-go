package stack

import "fmt"

type LinkedListStackNode struct {
	value int
	next  *LinkedListStackNode
}

func NewLinkedListStackNode(value int) *LinkedListStackNode {
	return &LinkedListStackNode{value: value}
}

type LinkedListStack struct {
	head *LinkedListStackNode
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{}
}

func (s *LinkedListStack) String() string {
	str := "["
	cursor := s.head

	for cursor != nil {
		str += fmt.Sprintf("%d -> ", cursor.value)
		cursor = cursor.next
	}

	str += "X]"

	return str
}

func (s *LinkedListStack) IsEmptyStack() bool {
	return s.head == nil
}

func (s *LinkedListStack) Push(value int) {
	node := NewLinkedListStackNode(value)
	if s.head == nil {
		s.head = node
	} else {
		node.next = s.head
		s.head = node
	}
}

func (s *LinkedListStack) Peek() (int, error) {
	if s.IsEmptyStack() {
		err := fmt.Errorf("error: stack underflow")
		return 0, err
	}

	return s.head.value, nil
}

func (s *LinkedListStack) Pop() (int, error) {
	if s.IsEmptyStack() {
		err := fmt.Errorf("error: stack underflow")
		return 0, err
	}

	value := s.head.value
	s.head = s.head.next

	return value, nil
}
