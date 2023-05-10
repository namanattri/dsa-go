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
	Length int
	Head   *UnrolledLinkedListNode
	Next   *UnrolledLinkedListBlock
}

func NewUnrolledLinkedListBlock(length int, head *UnrolledLinkedListNode) *UnrolledLinkedListBlock {
	return &UnrolledLinkedListBlock{Length: length, Head: head}
}

type UnrolledLinkedList struct {
	Head *UnrolledLinkedListBlock
}

func NewUnrolledLinkedList(values []int) *UnrolledLinkedList {
	return &UnrolledLinkedList{}
}
