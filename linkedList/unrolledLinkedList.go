package linkedList

import (
	"fmt"
	"math"
)

type UnrolledLinkedListNode struct {
	value int
	next  *UnrolledLinkedListNode
}

func NewUnrolledLinkedListNode(value int) *UnrolledLinkedListNode {
	return &UnrolledLinkedListNode{value: value}
}

// circular linked list block
type UnrolledLinkedListBlock struct {
	head *UnrolledLinkedListNode // Head pointing to the circular linked list inside the block
	next *UnrolledLinkedListBlock
}

func NewUnrolledLinkedListBlock(values []int) *UnrolledLinkedListBlock {
	head := (*UnrolledLinkedListNode)(nil)
	cursor := (*UnrolledLinkedListNode)(nil)

	for _, value := range values {
		node := NewUnrolledLinkedListNode(value)
		if head == nil {
			head = node
			cursor = node
			continue
		}
		cursor.next = node
		cursor = node
	}
	cursor.next = head

	return &UnrolledLinkedListBlock{head: head}
}

func (b *UnrolledLinkedListBlock) Length() int {
	len := 0
	cursor := b.head
	for cursor != nil {
		len++
		cursor = cursor.next
	}
	return len
}

func (b *UnrolledLinkedListBlock) Traverse() {
	cursor := b.head
	fmt.Print("[ ")
	for {
		fmt.Printf("%d -> ", cursor.value)
		if cursor.next == b.head {
			break
		}
		cursor = cursor.next
	}
	fmt.Print("X ]")
}

type UnrolledLinkedList struct {
	head      *UnrolledLinkedListBlock
	blockSize int
}

func NewUnrolledLinkedList(values []int) *UnrolledLinkedList {
	length := len(values)
	blockSize := math.Ceil(math.Sqrt(float64(length)))
	head := (*UnrolledLinkedListBlock)(nil)
	cursor := (*UnrolledLinkedListBlock)(nil)

	for i := 0; i < length; i += int(blockSize) {
		end := i + int(blockSize)
		if end > length {
			end = length
		}
		block := NewUnrolledLinkedListBlock(values[i:end])

		if head == nil {
			head = block
			cursor = block
			continue
		}

		cursor.next = block
		cursor = cursor.next
	}
	return &UnrolledLinkedList{head: head, blockSize: int(blockSize)}
}

func (l *UnrolledLinkedList) Length() int {
	len := 0
	cursor := l.head

	for cursor != nil {
		len++
		cursor = cursor.next
	}

	return len
}

func (l *UnrolledLinkedList) Traverse() {
	cursor := l.head
	for cursor != nil {
		cursor.Traverse()
		fmt.Print(" -> ")
		cursor = cursor.next
	}
	fmt.Println("X")
}

func (l *UnrolledLinkedList) InsertStart(value int) {
	node := NewUnrolledLinkedListNode(value)
	block := l.head

	for block != nil {
		fmt.Println(block.head.value)

		prev := (*UnrolledLinkedListNode)(nil)
		cursor := block.head

		blockLength := 0
		for cursor.next != block.head {
			prev = cursor
			cursor = cursor.next
			blockLength++
		}

		node.next = block.head
		block.head = node

		fmt.Printf("len: %d size: %d", blockLength, l.blockSize)
		if blockLength > l.blockSize {
			prev.next = block.head
			node = cursor
		} else {
			cursor.next = block.head
		}

		block = block.next
	}
}
