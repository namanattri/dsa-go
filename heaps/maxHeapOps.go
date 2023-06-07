package heaps

import "fmt"

func MaxHeapOps() {
	h := NewMaxHeap()

	fmt.Printf("h: %v\n", h)

	h.Insert(1)
	fmt.Printf("h: %v\n", h)

	h.Insert(2)
	fmt.Printf("h: %v\n", h)

	h.Insert(9)
	fmt.Printf("h: %v\n", h)

	h.Insert(5)
	fmt.Printf("h: %v\n", h)

	h.Insert(15)
	fmt.Printf("h: %v\n", h)

	h.Insert(3)
	fmt.Printf("h: %v\n", h)
}
