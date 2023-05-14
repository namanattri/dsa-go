package ops

import (
	"dsa-go/linkedList"
	"fmt"
)

func UnrolledLinkedListOps() {
	l := linkedList.NewUnrolledLinkedList([]int{7, 22, 13, 34, 76, 2, 88, 3, 12, 41, 67, 15, 72, 9, 4, 17, 25})
	fmt.Printf("Number of blocks in the unrolled linked list: %d\n", l.Length())
	fmt.Println("Traversing the unrolled linked list: ")
	l.Traverse()

	fmt.Printf("Inserting %d at the begining of the linked list\n", 4)
	l.InsertStart(4)
	l.Traverse()

	// fmt.Printf("Inserting %d at the end of the linked list\n", 8)
	// l.InsertEnd(8)
	// l.Traverse()

	// fmt.Printf("Inserting %d at index %d of the linked list\n", 15, 3)
	// l.InsertPos(15, 3)
	// l.Traverse()

	// fmt.Println("Deleting a node from the begining of the linked list")
	// l.DeleteStart()
	// l.Traverse()

	// fmt.Println("Deleting a node from the tail of the linked list")
	// l.DeleteEnd()
	// l.Traverse()

	// fmt.Printf("Deleting a node at index %d of the linked list\n", 4)
	// l.DeletePos(4)
	// l.Traverse()
}
