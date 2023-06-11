package graph

import "fmt"

type QueueNode struct {
	value rune
	next  *QueueNode
}

func NewQueueNode(value rune) *QueueNode {
	return &QueueNode{value: value}
}

type Queue struct {
	head *QueueNode
	tail *QueueNode
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(value rune) {
	n := NewQueueNode(value)
	if q.head == nil {
		q.head = n
		q.tail = n
	} else {
		q.tail.next = n
		q.tail = n
	}
}

func (q *Queue) Dequeue() (*QueueNode, error) {
	if q.head == nil {
		return nil, fmt.Errorf("error: queue underflow")
	}
	n := q.head

	if q.head == q.tail {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}

	return n, nil
}
