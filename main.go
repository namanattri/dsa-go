package main

import "dsa-go/linkedList"

func main() {
	l := linkedList.NewSinglyLinkedList([]int{1, 2, 3})
	l.Traverse()

	l.InsertStart(4)
	l.Traverse()
	l.InsertEnd(8)
	l.Traverse()
	l.InsertPos(15, 3)
	l.Traverse()
}
