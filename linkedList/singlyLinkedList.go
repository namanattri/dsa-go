package linkedList

import (
	"fmt"
)

type SinglyLinkedListNode struct {
	Value int
	Next  *SinglyLinkedListNode
}

func NewNode(value int) *SinglyLinkedListNode {
	return &SinglyLinkedListNode{Value: value}
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
}

func NewSinglyLinkedList(values []int) *SinglyLinkedList {
	prevNode := (*SinglyLinkedListNode)(nil)
	head := (*SinglyLinkedListNode)(nil)

	for _, value := range values {
		if prevNode == nil {
			node := NewNode(value)
			head = node
			prevNode = node
		} else {
			prevNode.Next = NewNode(value)
			prevNode = prevNode.Next
		}
	}
	return &SinglyLinkedList{head: head}
}

func (l *SinglyLinkedList) Traverse() {
	cursor := l.head

	for cursor != nil {
		fmt.Printf("%d -> ", cursor.Value)
		cursor = cursor.Next
	}
	fmt.Println("X")
}

func (l *SinglyLinkedList) InsertStart(value int) {
	node := NewNode(value)
	node.Next = l.head
	l.head = node
}

func (l *SinglyLinkedList) InsertEnd(value int) {
	cursor := l.head

	for cursor.Next != nil {
		cursor = cursor.Next
	}

	node := NewNode(value)
	cursor.Next = node
}

func (l *SinglyLinkedList) InsertPos(value int, pos int) {
	if pos == 0 {
		l.InsertStart(value)
		return
	}

	nodeIndex := 0
	cursor := l.head
	node := NewNode(value)

	for cursor.Next != nil {
		if nodeIndex == pos-1 {
			node.Next = cursor.Next
			cursor.Next = node
			return
		}
		cursor = cursor.Next
		nodeIndex++
	}

	cursor.Next = node
}

func (l *SinglyLinkedList) DeleteStart() {
	l.head = l.head.Next
}

func (l *SinglyLinkedList) DeleteEnd() {
	cursor := l.head
	prevNode := cursor

	for cursor.Next != nil {
		prevNode = cursor
		cursor = cursor.Next
	}

	prevNode.Next = nil
}

func (l *SinglyLinkedList) DeletePos(pos int) {
	if pos == 0 {
		l.DeleteStart()
	}

	cursor := l.head
	nodeIndex := 0

	for cursor.Next != nil {
		if nodeIndex == pos-1 {
			cursor.Next = cursor.Next.Next
			return
		}
		cursor = cursor.Next
		nodeIndex++
	}
}
