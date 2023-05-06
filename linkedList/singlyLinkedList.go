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
	fmt.Println("Traversing the singly linked list: ")
	cursor := l.head

	for cursor != nil {
		fmt.Printf("%d -> ", cursor.Value)
		cursor = cursor.Next
	}
	fmt.Println("X")
}

func (l *SinglyLinkedList) InsertStart(value int) {
	fmt.Printf("Inserting %d at the begining of the linked list\n", value)
	node := NewNode(value)
	node.Next = l.head
	l.head = node
}

func (l *SinglyLinkedList) InsertEnd(value int) {
	fmt.Printf("Inserting %d at the begining of the linked list\n", value)
	cursor := l.head

	for cursor.Next != nil {
		cursor = cursor.Next
	}

	node := NewNode(value)
	cursor.Next = node
}

func (l *SinglyLinkedList) InsertPos(value int, pos int) {
	fmt.Printf("Inserting %d as node %d of the linked list\n", value, pos)
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
