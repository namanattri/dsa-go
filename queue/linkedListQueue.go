package queue

import "fmt"

type LinkedListQueueNode struct {
	value int
	next  *LinkedListQueueNode
}

func NewLinkedListQueueNode(value int) *LinkedListQueueNode {
	return &LinkedListQueueNode{value: value}
}

type LinkedListQueue struct {
	front *LinkedListQueueNode
	rear  *LinkedListQueueNode
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{}
}

func (q *LinkedListQueue) IsEmpty() bool {
	return q.front == nil
}

func (q *LinkedListQueue) Size() int {
	size := 0

	cursor := q.front

	for cursor != nil {
		size++
	}

	return size
}

func (q *LinkedListQueue) Enqueue(value int) {
	node := NewLinkedListQueueNode(value)

	if q.front == nil {
		q.front = node
		q.rear = node
		return
	}

	q.rear.next = node
	q.rear = node
}

func (q *LinkedListQueue) Front() (int, error) {
	if q.front == nil {
		return -1, fmt.Errorf("queue: underflow")
	}

	return q.front.value, nil
}

func (q *LinkedListQueue) Dequeue() (int, error) {
	if q.front == nil {
		return -1, fmt.Errorf("queue: underflow")
	}

	node := q.front

	if q.front == q.rear {
		q.front = nil
		q.rear = nil
	} else {
		q.front = q.front.next
	}

	return node.value, nil
}

func (q *LinkedListQueue) String() string {
	str := "["
	cursor := q.front

	for cursor != nil {
		str += fmt.Sprintf("%d -> ", cursor.value)
		cursor = cursor.next
	}

	str += "X]"

	return str
}
