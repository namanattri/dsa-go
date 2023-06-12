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
	head     *QueueNode
	tail     *QueueNode
	elements map[rune]bool
}

func NewQueue() *Queue {
	return &Queue{elements: make(map[rune]bool)}
}

func (q *Queue) String() string {
	res := "["

	cursor := q.head

	for cursor != nil {
		res += fmt.Sprintf("%c -> ", cursor.value)
		cursor = cursor.next
	}

	return res + "X]"
}

func (q *Queue) IsEmpty() bool {
	return q.head == nil
}

func (q *Queue) Has(value rune) bool {
	return q.elements[value]
}

func (q *Queue) Enqueue(value rune) {
	q.elements[value] = true
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

	q.elements[n.value] = false

	return n, nil
}
