package graph

import "fmt"

func HeapOps() {
	h := NewMinHeap()

	nodes := map[rune]int{
		'A': 5,
		'B': 3,
		'C': 2,
		'D': 1,
		'E': 4,
	}

	for value, priority := range nodes {
		fmt.Printf("Inserting %c with priority %d\n", value, priority)
		h.Insert(value, priority)
	}

	for i := 0; i < 10; i++ {
		n, error := h.DeleteMin()

		if error != nil {
			fmt.Println("Error encountered while delete min", error)
		} else {
			fmt.Printf("Dequeued %c with priority %d\n", n.value, n.priority)
		}
	}

	for value, priority := range nodes {
		fmt.Printf("Inserting %c with priority %d\n", value, priority)
		h.Insert(value, priority)
	}

	h.UpdatePriority('E', 0)
	h.UpdatePriority('D', 7)

	for i := 0; i < 5; i++ {
		n, error := h.DeleteMin()

		if error != nil {
			fmt.Println("Error encountered while delete min", error)
		} else {
			fmt.Printf("Dequeued %c with priority %d\n", n.value, n.priority)
		}
	}
}
