package heaps

import "fmt"

func MaxHeapOps() {
	h := NewMaxHeap()

	h.Build([]int{1, 2, 3, 5, 7, 8, 10, 12, 14, 16, 19, 31})
	fmt.Printf("h = %v\n", h)

	fmt.Printf("Deleting Max = %d : h = %v\n", h.DeleteMax(), h)

	h2 := NewMaxHeap()

	h2.BuildOptimized([]int{1, 5, 14, 2, 10, 21, 18, 3, 11, 8, 7, 12})
	fmt.Printf("h2 = %v\n", h2)

	fmt.Printf("Deleting Max = %d : h2 = %v\n", h2.DeleteMax(), h2)

	h3 := NewMaxHeap()

	h3.BuildOptimized([]int{1, 5, 14, 2, 10, 21, 18, 3, 11, 8, 7, 12})
	fmt.Printf("h3 = %v\n", h3)

	h3.Sort()
	fmt.Printf("h3 = %v\n", h3)

}
