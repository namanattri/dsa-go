package trees

import "fmt"

type GenericQueue[T NodeType] struct {
	front *GenericLinkedListNode[T]
	rear  *GenericLinkedListNode[T]
}

func NewGenericQueue[T NodeType]() *GenericQueue[T] {
	return &GenericQueue[T]{}
}

func (s *GenericQueue[T]) IsEmpty() bool {
	return s.front == nil
}

func (q *GenericQueue[T]) Enqueue(value T) {
	node := NewGenericLinkedListNode(value)

	if q.front == nil {
		q.front = node
		q.rear = node
		return
	}

	q.rear.next = node
	q.rear = node
}

func (q *GenericQueue[T]) Front() (T, error) {
	if q.front == nil {
		var res T
		return res, fmt.Errorf("queue: underflow")
	}

	return q.front.value, nil
}

func (q *GenericQueue[T]) Dequeue() (T, error) {
	if q.front == nil {
		var res T
		return res, fmt.Errorf("queue: underflow")
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
