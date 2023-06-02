package linkedList

import (
	"fmt"
)

func CircularLinkedListOps() {
	l := NewCircularLinkedList([]int{7, 22, 13, 34, 76, 2, 88})
	fmt.Printf("Length of the circular linked list: %d\n", l.Length())
	fmt.Println("Traversing the circular linked list: ")
	l.Traverse()

	fmt.Printf("Inserting %d at the begining of the linked list\n", 4)
	l.InsertStart(4)
	l.Traverse()

	fmt.Printf("Inserting %d at the end of the linked list\n", 8)
	l.InsertEnd(8)
	l.Traverse()

	pos := 3
	val := 15
	fmt.Printf("Inserting %d at index %d of the linked list\n", val, pos)
	l.InsertPos(15, 3)
	l.Traverse()

	pos = 0
	val = 33
	fmt.Printf("Inserting %d at index %d of the linked list\n", val, pos)
	l.InsertPos(val, pos)
	l.Traverse()

	pos = l.Length() - 1
	val = 27

	fmt.Printf("Inserting %d at index %d of the linked list\n", val, pos)
	l.InsertPos(val, pos)
	l.Traverse()

	fmt.Println("Deleting a node from the begining of the linked list")
	l.DeleteStart()
	l.Traverse()

	fmt.Println("Deleting a node from the tail of the linked list")
	l.DeleteEnd()
	l.Traverse()

	pos = 4
	fmt.Printf("Deleting a node at index %d of the linked list\n", pos)
	l.DeletePos(pos)
	l.Traverse()

	pos = 0
	fmt.Printf("Deleting a node at index %d of the linked list\n", pos)
	l.DeletePos(pos)
	l.Traverse()

	pos = l.Length() - 1
	fmt.Printf("Deleting a node at index %d of the linked list\n", pos)
	l.DeletePos(pos)
	l.Traverse()

	fmt.Printf("Inserting %d at the begining of the linked list\n", 52)
	l.InsertStart(52)
	l.Traverse()

	fmt.Printf("Inserting %d at the end of the linked list\n", 85)
	l.InsertEnd(85)
	l.Traverse()
}
