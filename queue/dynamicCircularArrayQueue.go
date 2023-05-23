package queue

import "fmt"

type DynamicCircularArrayQueue struct {
	front    int
	rear     int
	capacity int
	array    []int
}

func NewDynamicCircularArrayQueue(size int) *DynamicCircularArrayQueue {
	return &DynamicCircularArrayQueue{front: -1, rear: -1, capacity: size, array: make([]int, size)}
}

func (q *DynamicCircularArrayQueue) Capacity() int {
	return q.capacity
}

func (q *DynamicCircularArrayQueue) Array() []int {
	return q.array
}

func (q *DynamicCircularArrayQueue) IsEmpty() bool {
	return q.front == -1
}

func (q *DynamicCircularArrayQueue) isFull() bool {
	return (q.rear+1)%q.capacity == q.front
}

func (q *DynamicCircularArrayQueue) Size() int {
	if q.front == -1 {
		return 0
	}

	if q.front > q.rear {
		return q.capacity - q.front + q.rear + 1
	}
	return q.rear - q.front + 1
}

func (q *DynamicCircularArrayQueue) Front() (int, error) {
	if q.IsEmpty() {
		return -1, fmt.Errorf("queue: underflow")
	}

	return q.array[q.front], nil
}

func (q *DynamicCircularArrayQueue) resize() {
	size := q.capacity
	q.array = append(q.array, make([]int, q.capacity)...)
	q.capacity *= 2

	if q.front > q.rear {
		for i := 0; i < q.front; i++ {
			q.array[i+size] = q.array[i]
			q.array[i] = 0
		}
		q.rear += size
	}
}

func (q *DynamicCircularArrayQueue) Enqueue(value int) {
	if q.isFull() {
		q.resize()
	}

	q.rear = (q.rear + 1) % q.capacity
	q.array[q.rear] = value

	if q.front == -1 {
		q.front = q.rear
	}
}

func (q *DynamicCircularArrayQueue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return -1, fmt.Errorf("queue: underflow")
	}

	value := q.array[q.front]
	q.array[q.front] = 0

	if q.front == q.rear {
		q.front = -1
		q.rear = -1
	} else {
		q.front = (q.front + 1) % q.capacity
	}

	return value, nil
}

func (q *DynamicCircularArrayQueue) String() string {
	return fmt.Sprintf("Queue:: Front: %d Rear: %d Size: %d, Capacity: %d, Values: %v\n", q.front, q.rear, q.Size(), q.Capacity(), q.Array())
}
