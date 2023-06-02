package queue

import (
	"fmt"
	"math/rand"
)

func DynamicCircularArrayQueueOps() {
	q := NewDynamicCircularArrayQueue(10)
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
		q.Enqueue(value)
		fmt.Printf("Enqueued %d to queue\n", value)
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
			fmt.Printf("Error encountered while dequeueing: %v\n", err)
		} else {
			fmt.Printf("Dequeue: %d\n", value)
		}
		fmt.Println(q)
	}

	// fill 20 elements in the queue
	fmt.Println("Enqueueing 20 elements to the queue")

	numCount := 20
	numbers := make([]int, numCount)
	for i := 0; i < numCount; i++ {
		numbers[i] = rand.Intn(99) + 1
	}

	for _, value := range numbers {
		q.Enqueue(value)
		fmt.Printf("Enqueued %d to queue\n", value)
		fmt.Println(q)
	}

	// dequeue 5 elements
	n = 5
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

	// fill 3 elements in the queue
	fmt.Println("Enqueueing 3 elements to the queue")

	numCount = 3
	numbers = make([]int, numCount)
	for i := 0; i < numCount; i++ {
		numbers[i] = rand.Intn(99) + 1
	}

	for _, value := range numbers {
		q.Enqueue(value)
		fmt.Printf("Enqueued %d to queue\n", value)
		fmt.Println(q)
	}

	// fill 10 elements and see if the elements are arranged
	fmt.Println("Enqueueing 10 elements to the queue")

	numCount = 10
	numbers = make([]int, numCount)
	for i := 0; i < numCount; i++ {
		numbers[i] = rand.Intn(99) + 1
	}

	for _, value := range numbers {
		q.Enqueue(value)
		fmt.Printf("Enqueued %d to queue\n", value)
		fmt.Println(q)
	}

	// fill 10 elements and see if the elements are arranged
	fmt.Println("Enqueueing 10 more elements to the queue")

	numCount = 10
	numbers = make([]int, numCount)
	for i := 0; i < numCount; i++ {
		numbers[i] = rand.Intn(99) + 1
	}

	for _, value := range numbers {
		q.Enqueue(value)
		fmt.Printf("Enqueued %d to queue\n", value)
		fmt.Println(q)
	}
}
