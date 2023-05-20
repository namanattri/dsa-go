package ops

import (
	"dsa-go/queue"
	"fmt"
)

func StaticArrayQueueOps() {
	q := queue.NewStaticCircularArrayQueue(10)
	fmt.Println(q)

	fmt.Printf("Trying to Front\n")
	value, err := q.Front()

	if err != nil {
		fmt.Printf("Error encountered while fronting: %v\n", err)
	} else {
		fmt.Printf("Front: %d\n", value)
	}

	// push multiple values to stack

	values := []int{45, 72, 83, 19, 61, 7, 92, 38, 51, 24, 86, 11}

	for _, value := range values {
		err := q.Enqueue(value)

		if err != nil {
			fmt.Printf("Error while enqueueing %d: %v\n", value, err)
		} else {
			fmt.Printf("Enqueued %d to queue\n", value)
		}
		fmt.Println(q)
	}

	fmt.Printf("Trying to Front\n")
	value, err = q.Front()

	if err != nil {
		fmt.Printf("Error encountered while fronting: %v\n", err)
	} else {
		fmt.Printf("Front: %d\n", value)
	}

	n := 5
	fmt.Printf("Dequeueing %d elements\n", n)

	for i := 0; i < n; i++ {
		value, err := q.Dequeue()

		if err != nil {
			fmt.Printf("Error encountered while dequeueing: %v\n", err)
		} else {
			fmt.Printf("Dequeue: %d\n", value)
		}
		fmt.Println(q)
	}

	fmt.Printf("Trying to Front\n")
	value, err = q.Front()

	if err != nil {
		fmt.Printf("Error encountered while fronting: %v\n", err)
	} else {
		fmt.Printf("Front: %d\n", value)
	}

	n = 8
	fmt.Printf("Dequeueing %d elements\n", n)

	for i := 0; i < n; i++ {
		value, err := q.Dequeue()

		if err != nil {
			fmt.Printf("Error encountered while popping: %v\n", err)
		} else {
			fmt.Printf("Pop: %d\n", value)
		}
		fmt.Println(q)
	}
}
