package linkedList

import (
	"fmt"
)

func XORLinkedListOps() {
	// concept
	// a ^ a = 0
	// a ^ 0 = a
	// a ^ b = b ^ a
	// a ^ (b ^ c) = (a ^ b) ^ c
	// (a ^ b) ^ a = b
	// (a ^ b) ^ b = a
	// a -> b -> c
	// b.pointerDiff = a ^ c
	// b.next == (a ^ c) ^ a == c
	// b.prev == (a ^ c) ^ c == a

	l := NewXORLinkedList([]int{7, 22, 13, 34, 76, 2, 88})
	fmt.Printf("Length of the xor linked list: %d\n", l.Length())
	fmt.Println("Traversing the XOR linked list: ")
	l.Traverse()
	fmt.Println("Traversing the XOR linked list backwards: ")
	// l.TraverseBackwards()

	fmt.Printf("Inserting %d at the begining of the linked list\n", 4)
	l.InsertStart(4)
	l.Traverse()
	// l.TraverseBackwards()

	fmt.Printf("Inserting %d at the end of the linked list\n", 8)
	l.InsertEnd(8)
	l.Traverse()
	// l.TraverseBackwards()

	fmt.Printf("Inserting %d at index %d of the linked list\n", 15, 3)
	l.InsertPos(15, 3)
	l.Traverse()
	l.TraverseBackwards()

	fmt.Println("Deleting a node from the begining of the linked list")
	l.DeleteStart()
	l.Traverse()
	l.TraverseBackwards()

	fmt.Println("Deleting a node from the tail of the linked list")
	l.DeleteEnd()
	l.Traverse()
	l.TraverseBackwards()

	fmt.Printf("Deleting a node at index %d of the linked list\n", 4)
	l.DeletePos(4)
	l.Traverse()
	l.TraverseBackwards()
}
