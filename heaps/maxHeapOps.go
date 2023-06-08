package heaps

import "fmt"

func MaxHeapOps() {
	h := NewMaxHeap()

	h.Build([]int{1, 2, 3, 5, 7, 8, 10, 12, 14, 16, 19, 31})
	fmt.Printf("h = %v\n", h)

	fmt.Printf("Deleting Max = %d : h = %v\n", h.DeleteMax(), h)

}
