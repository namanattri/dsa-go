package linkedList

import "fmt"

type DoublyLinkedListNode struct {
	Value int
	Prev  *DoublyLinkedListNode
	Next  *DoublyLinkedListNode
}

func NewDoublyLinkeListNode(value int) *DoublyLinkedListNode {
	return &DoublyLinkedListNode{Value: value}
}

type DoublyLinkedList struct {
	head *DoublyLinkedListNode
	tail *DoublyLinkedListNode
}

func NewDoublyLinkeList(values []int) *DoublyLinkedList {
	head := (*DoublyLinkedListNode)(nil)
	tail := (*DoublyLinkedListNode)(nil)
	for _, value := range values {
		node := NewDoublyLinkeListNode(value)
		if head == nil {
			head = node
			tail = node
			continue
		}

		node.Prev = tail
		tail.Next = node
		tail = node
	}

	return &DoublyLinkedList{head: head, tail: tail}
}

func (l *DoublyLinkedList) Length() int {
	len := 0
	cursor := l.head

	for cursor != nil {
		len++
		cursor = cursor.Next
	}

	return len
}

func (l *DoublyLinkedList) Traverse() {
	cursor := l.head
	for cursor != nil {
		fmt.Printf("%d -> ", cursor.Value)
		cursor = cursor.Next
	}
	fmt.Println("X")
}

func (l *DoublyLinkedList) TraverseBackwards() {
	cursor := l.tail
	for cursor != nil {
		fmt.Printf("%d -> ", cursor.Value)
		cursor = cursor.Prev
	}
	fmt.Println("X")
}

func (l *DoublyLinkedList) InsertStart(value int) {
	node := NewDoublyLinkeListNode(value)
	node.Next = l.head
	l.head.Prev = node
	l.head = node
}

func (l *DoublyLinkedList) InsertEnd(value int) {
	node := NewDoublyLinkeListNode(value)
	node.Prev = l.tail
	l.tail.Next = node
	l.tail = node
}

func (l *DoublyLinkedList) InsertPos(value int, pos int) {
	if pos == 0 {
		l.InsertStart(value)
		return
	}

	if pos > l.Length()-1 {
		return
	}

	cursor := l.head
	nodeIndex := 0

	for cursor != nil {
		if nodeIndex == pos {
			node := NewDoublyLinkeListNode(value)
			cursor.Prev.Next = node
			node.Prev = cursor.Prev
			cursor.Prev = node
			node.Next = cursor
			return
		}
		cursor = cursor.Next
		nodeIndex++
	}
}

func (l *DoublyLinkedList) DeleteStart() {
	l.head.Next.Prev = nil
	l.head = l.head.Next
}

func (l *DoublyLinkedList) DeleteEnd() {
	l.tail.Prev.Next = nil
	l.tail = l.tail.Prev
}

func (l *DoublyLinkedList) DeletePos(pos int) {
	if pos == 0 {
		l.DeleteStart()
		return
	}

	if pos > l.Length()-1 {
		return
	}

	cursor := l.head
	nodeIndex := 0

	for cursor != nil {
		if nodeIndex == pos {
			cursor.Prev.Next = cursor.Next
			cursor.Next.Prev = cursor.Prev
			return
		}
		cursor = cursor.Next
		nodeIndex++
	}
}
