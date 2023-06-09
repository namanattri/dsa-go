package linkedList

import (
	"fmt"
	"unsafe"
)

type XORLinkedListNode struct {
	Value             int
	PointerDifference uintptr
}

func NewXORLinkedListNode(value int) *XORLinkedListNode {
	return &XORLinkedListNode{Value: value}
}

type XORLinkedList struct {
	head *XORLinkedListNode
	tail *XORLinkedListNode
}

func NewXORLinkedList(values []int) *XORLinkedList {
	head := (*XORLinkedListNode)(nil)
	tail := (*XORLinkedListNode)(nil)

	for _, value := range values {
		node := NewXORLinkedListNode(value)
		if head == nil {
			head = node
			tail = node
		} else {
			node.PointerDifference = uintptr(unsafe.Pointer(tail)) ^ uintptr(0)
			tail.PointerDifference = tail.PointerDifference ^ uintptr(unsafe.Pointer(node))
			tail = node
		}
	}

	return &XORLinkedList{head: head, tail: tail}
}

func (l *XORLinkedList) Length() int {
	len := 0
	prev := uintptr(0)
	cursor := l.head

	for uintptr(unsafe.Pointer(cursor)) != uintptr(0) {
		next := prev ^ cursor.PointerDifference
		prev = uintptr(unsafe.Pointer(cursor))
		cursor = (*XORLinkedListNode)(unsafe.Pointer(next))
		len++
	}

	return len
}

func (l *XORLinkedList) Traverse() {
	prev := uintptr(0)
	cursor := l.head

	for uintptr(unsafe.Pointer(cursor)) != uintptr(0) {
		fmt.Printf("%d -> ", cursor.Value)
		next := cursor.PointerDifference ^ prev
		prev = uintptr(unsafe.Pointer(cursor))
		cursor = (*XORLinkedListNode)(unsafe.Pointer(next))
	}
	fmt.Println("X")
}

func (l *XORLinkedList) TraverseBackwards() {
	next := uintptr(0)
	cursor := l.tail

	for uintptr(unsafe.Pointer(cursor)) != uintptr(0) {
		fmt.Printf("%d -> ", cursor.Value)
		prev := cursor.PointerDifference ^ next
		next = uintptr(unsafe.Pointer(cursor))
		cursor = (*XORLinkedListNode)(unsafe.Pointer(prev))
	}
	fmt.Println("X")
}

func (l *XORLinkedList) InsertStart(value int) {
	node := NewXORLinkedListNode(value)
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.PointerDifference = uintptr(0) ^ uintptr(unsafe.Pointer(l.head))
		l.head.PointerDifference = uintptr(unsafe.Pointer(node)) ^ l.head.PointerDifference
		l.head = node
	}
}

func (l *XORLinkedList) InsertEnd(value int) {
	node := NewXORLinkedListNode(value)
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.PointerDifference = uintptr(unsafe.Pointer(l.tail)) ^ uintptr(0)
		l.tail.PointerDifference = l.tail.PointerDifference ^ uintptr(unsafe.Pointer(node))
		l.tail = node
	}
}

func (l *XORLinkedList) InsertPos(value int, pos int) {
	// Check if pos is out of bounds
	if pos < 0 || pos > l.Length() {
		return
	}

	// Create the new node
	node := NewXORLinkedListNode(value)

	// Traverse the list to find the node at the specified position
	prev := uintptr(0)
	cursor := l.head
	i := 0
	for i < pos && cursor != nil {
		next := prev ^ cursor.PointerDifference
		prev = uintptr(unsafe.Pointer(cursor))
		cursor = (*XORLinkedListNode)(unsafe.Pointer(next))
		i++
	}

	// Insert the new node between the previous and current nodes
	if prev == uintptr(0) { // Insert at the beginning of the list
		node.PointerDifference = uintptr(0) ^ uintptr(unsafe.Pointer(cursor))
		cursor.PointerDifference = cursor.PointerDifference ^ uintptr(unsafe.Pointer(node))
		l.head = node
	} else if cursor == nil { // Insert at the end of the list
		node.PointerDifference = prev ^ uintptr(0)
		prevNode := (*XORLinkedListNode)(unsafe.Pointer(prev))
		prevNode.PointerDifference = prevNode.PointerDifference ^ uintptr(unsafe.Pointer(node))
		l.tail = node
	} else { // Insert in the middle of the list
		node.PointerDifference = prev ^ uintptr(unsafe.Pointer(cursor))
		prevNode := (*XORLinkedListNode)(unsafe.Pointer(prev))
		prevNode.PointerDifference = prevNode.PointerDifference ^ uintptr(unsafe.Pointer(node)) ^ uintptr(unsafe.Pointer(cursor))
		cursor.PointerDifference = cursor.PointerDifference ^ prev ^ uintptr(unsafe.Pointer(node))
	}
}

func (l *XORLinkedList) DeleteStart() {
	if l.Length() == 0 {
		return
	}

	if l.Length() == 1 {
		l.head = nil
		l.tail = nil
		return
	}

	next := l.head.PointerDifference ^ uintptr(0)
	nextNode := (*XORLinkedListNode)(unsafe.Pointer(next))
	nextNode.PointerDifference = nextNode.PointerDifference ^ uintptr(unsafe.Pointer(l.head)) ^ uintptr(0)
	l.head = nextNode
}

func (l *XORLinkedList) DeleteEnd() {
	if l.Length() == 0 {
		return
	}

	if l.Length() == 1 {
		l.head = nil
		l.tail = nil
		return
	}

	prev := l.tail.PointerDifference ^ uintptr(0)
	prevNode := (*XORLinkedListNode)(unsafe.Pointer(prev))
	prevNode.PointerDifference = prevNode.PointerDifference ^ uintptr(unsafe.Pointer(l.tail)) ^ uintptr(0)
	l.tail = prevNode
}

func (l *XORLinkedList) DeletePos(pos int) {
	// Check if pos is out of bounds
	if pos < 0 || pos > l.Length() {
		return
	}

	// Traverse the list to find the node at the specified position
	prev := uintptr(0)
	cursor := l.head
	i := 0
	for i < pos && cursor != nil {
		next := prev ^ cursor.PointerDifference
		prev = uintptr(unsafe.Pointer(cursor))
		cursor = (*XORLinkedListNode)(unsafe.Pointer(next))
		i++
	}

	prevNode := (*XORLinkedListNode)(unsafe.Pointer(prev))
	nextNode := (*XORLinkedListNode)(unsafe.Pointer(prev ^ cursor.PointerDifference))
	prevNode.PointerDifference = prevNode.PointerDifference ^ uintptr(unsafe.Pointer(cursor)) ^ uintptr(unsafe.Pointer(nextNode))
	nextNode.PointerDifference = nextNode.PointerDifference ^ uintptr(unsafe.Pointer(cursor)) ^ uintptr(unsafe.Pointer(prevNode))
}
