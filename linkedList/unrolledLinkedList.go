package linkedList

type UnrolledLinkedListNode struct {
	Value int
	Next  *UnrolledLinkedListNode
}

func NewUnrolledLinkedListNode(value int) *UnrolledLinkedListNode {
	return &UnrolledLinkedListNode{Value: value}
}

// circular linked list block
type UnrolledLinkedListBlock struct {
	Head *UnrolledLinkedListNode // Head pointing to the circular linked list inside the block
	Next *UnrolledLinkedListBlock
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
		cursor.Next = node
		cursor = node
	}
	cursor.Next = head

	return &UnrolledLinkedListBlock{Head: head}
}

type UnrolledLinkedList struct {
	Head *UnrolledLinkedListBlock
}

func NewUnrolledLinkedList(values []int) *UnrolledLinkedList {
	return &UnrolledLinkedList{}
}
