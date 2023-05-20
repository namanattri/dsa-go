package queue

import "fmt"

type StaticCircularArrayQueue struct {
	front    int
	rear     int
	capacity int
	array    []int
}

func NewStaticCircularArrayQueue(size int) *StaticCircularArrayQueue {
	return &StaticCircularArrayQueue{front: -1, rear: -1, capacity: size, array: make([]int, size)}
}

func (q *StaticCircularArrayQueue) Capacity() int {
	return q.capacity
}

func (q *StaticCircularArrayQueue) Array() []int {
	return q.array
}

func (q *StaticCircularArrayQueue) IsEmpty() bool {
	return q.front == -1
}

func (q *StaticCircularArrayQueue) IsFull() bool {
	return (q.rear+1)%q.capacity == q.front
}

func (q *StaticCircularArrayQueue) Size() int {
	if q.front == -1 {
		return 0
	}

	if q.front > q.rear {
		return q.capacity - q.front + q.rear + 1
	}
	return q.rear - q.front + 1
}

func (q *StaticCircularArrayQueue) Front() (int, error) {
	if q.IsEmpty() {
		return -1, fmt.Errorf("queue: underflow")
	}
	return q.array[q.front], nil
}

func (q *StaticCircularArrayQueue) Enqueue(value int) error {
	if q.IsFull() {
		return fmt.Errorf("queue: overflow")
	}

	q.rear = (q.rear + 1) % q.capacity
	q.array[q.rear] = value
	if q.front == -1 {
		q.front = q.rear
	}
	return nil
}

func (q *StaticCircularArrayQueue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("queue: underflow")
	}

	value := q.array[q.front]
	if q.front == q.rear {
		q.front = -1
		q.rear = -1
	} else {
		q.front = (q.front + 1) % q.capacity
	}

	return value, nil
}

func (q *StaticCircularArrayQueue) String() string {
	return fmt.Sprintf("Stack:: Front: %d Rear: %d Size: %d, Capacity: %d, Values: %v\n", q.front, q.rear, q.Size(), q.Capacity(), q.Array())
}
