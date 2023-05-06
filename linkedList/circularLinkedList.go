package linkedList

import (
	"fmt"
)

type CircularLinkedListNode struct {
	Value int
	Next  *CircularLinkedListNode
}

func NewCircularLinkedListNode(value int) *CircularLinkedListNode {
	return &CircularLinkedListNode{Value: value}
}

type CircularLinkedList struct {
	head *CircularLinkedListNode
}

func NewCircularLinkedList(values []int) *CircularLinkedList {
	head := (*CircularLinkedListNode)(nil)
	cursor := (*CircularLinkedListNode)(nil)
	for _, value := range values {
		node := NewCircularLinkedListNode(value)
		if head == nil {
			head = node
			cursor = head
			continue
		}
		cursor.Next = node
		cursor = node
	}

	cursor.Next = head

	return &CircularLinkedList{head: head}
}

func (l *CircularLinkedList) Length() int {
	if l.head == nil {
		return 0
	}

	cursor := l.head
	len := 0

	for {
		len++
		cursor = cursor.Next

		if cursor == l.head {
			break
		}
	}

	return len
}

func (l *CircularLinkedList) Traverse() {
	if l.head == nil {
		return
	}

	cursor := l.head

	for {
		fmt.Printf("%d -> ", cursor.Value)

		cursor = cursor.Next

		if cursor == l.head {
			break
		}
	}

	fmt.Println("X")
}

func (l *CircularLinkedList) InsertStart(value int) {
	node := NewCircularLinkedListNode(value)
	node.Next = node

	if l.head == nil {
		l.head = node
		return
	}

	cursor := l.head

	for cursor.Next != l.head {
		cursor = cursor.Next
	}

	cursor.Next = node
	node.Next = l.head
	l.head = node
}

func (l *CircularLinkedList) InsertEnd(value int) {
	node := NewCircularLinkedListNode(value)
	node.Next = node

	if l.head == nil {
		l.head = node
		return
	}

	cursor := l.head

	for cursor.Next != l.head {
		cursor = cursor.Next
	}

	cursor.Next = node
	node.Next = l.head
}

func (l *CircularLinkedList) InsertPos(value int, pos int) {
	if pos == 0 {
		l.InsertStart(value)
		return
	}

	node := NewCircularLinkedListNode(value)
	node.Next = node

	prev := (*CircularLinkedListNode)(nil)
	cursor := l.head
	nodeIndex := 0

	for {
		if nodeIndex == pos {
			prev.Next = node
			node.Next = cursor
			return
		}
		prev = cursor
		cursor = cursor.Next
		nodeIndex++

		if cursor == l.head {
			break
		}
	}
}

func (l *CircularLinkedList) DeleteStart() {
	if l.head == nil {
		return
	}

	cursor := l.head

	for {
		cursor = cursor.Next
		if cursor.Next == l.head {
			break
		}
	}

	l.head = l.head.Next
	cursor.Next = l.head
}

func (l *CircularLinkedList) DeleteEnd() {
	if l.head == nil {
		return
	}

	prev := (*CircularLinkedListNode)(nil)
	cursor := l.head

	for {
		prev = cursor
		cursor = cursor.Next
		if cursor.Next == l.head {
			break
		}
	}

	prev.Next = l.head
}

func (l *CircularLinkedList) DeletePos(pos int) {
	if pos == 0 {
		l.DeleteStart()
		return
	}

	prev := (*CircularLinkedListNode)(nil)
	cursor := l.head
	nodeIndex := 0

	for {
		if nodeIndex == pos {
			prev.Next = cursor.Next
			return
		}

		prev = cursor
		cursor = cursor.Next
		nodeIndex++

		if cursor == l.head {
			break
		}
	}
}
