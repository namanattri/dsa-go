package linkedList

import "fmt"

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
	fmt.Println("Traversion singly linked list: ")
	cursor := l.head

	for cursor != nil {
		fmt.Println(cursor.Value)
		cursor = cursor.Next
	}
}
