package graph

import "fmt"

const INITIAL_CAPACITY = 1

type MinHeapNode struct {
	value    rune
	priority int
}

func NewMinHeapNode(value rune, priority int) *MinHeapNode {
	return &MinHeapNode{value: value, priority: priority}
}

type MinHeap struct {
	array    []*MinHeapNode
	count    int
	capacity int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{array: make([]*MinHeapNode, INITIAL_CAPACITY), count: 0, capacity: INITIAL_CAPACITY}
}

func (h *MinHeap) Resize() {
	h.array = append(h.array, make([]*MinHeapNode, h.capacity)...)
	h.capacity *= 2
}

func (h *MinHeap) IsEmpty() bool {
	return h.count == 0
}

func (h *MinHeap) ParentOf(index int) int {
	if index <= 0 || index >= h.count {
		return -1
	}
	return (index - 1) / 2
}

func (h *MinHeap) LeftChildOf(index int) int {
	left := 2*index + 1
	if left >= h.count {
		return -1
	}
	return left
}

func (h *MinHeap) RightChildOf(index int) int {
	right := 2*index + 2
	if right >= h.count {
		return -1
	}
	return right
}

func (h *MinHeap) DeleteMin() (*MinHeapNode, error) {
	if h.IsEmpty() {
		return nil, fmt.Errorf("error: heap underflow")
	}

	n := h.array[0]
	h.array[0] = h.array[h.count-1]
	h.count--

	h.PercolateDown(0)

	return n, nil
}

func (h *MinHeap) PercolateDown(index int) {
	min := index

	left := h.LeftChildOf(index)
	right := h.RightChildOf(index)

	if left != -1 && h.array[left].priority < h.array[min].priority {
		min = left
	}

	if right != -1 && h.array[right].priority < h.array[min].priority {
		min = right
	}

	if min != index {
		// swap
		tmp := h.array[index]
		h.array[index] = h.array[min]
		h.array[min] = tmp

		h.PercolateDown(min)
	}
}

func (h *MinHeap) Insert(value rune, priority int) {
	n := NewMinHeapNode(value, priority)

	if h.count == h.capacity {
		h.Resize()
	}

	h.count++

	index := h.count - 1

	for index > 0 && n.priority < h.array[h.ParentOf(index)].priority {
		h.array[index] = h.array[h.ParentOf(index)]
		index = h.ParentOf(index)
	}

	h.array[index] = n
}
