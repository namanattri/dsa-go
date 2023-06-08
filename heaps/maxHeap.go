package heaps

import "fmt"

const INITIAL_CAPACITY = 1

type MaxHeap struct {
	capacity int
	count    int
	array    []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{capacity: INITIAL_CAPACITY, count: 0, array: make([]int, INITIAL_CAPACITY)}
}

func (h *MaxHeap) Resize() {
	h.array = append(h.array, make([]int, h.capacity)...)
	h.capacity *= 2
}

func (h *MaxHeap) Insert(value int) {
	if h.count == h.capacity {
		h.Resize()
	}
	h.count++
	index := h.count - 1

	for index > 0 && value > h.array[(index-1)/2] {
		h.array[index] = h.array[(index-1)/2]
		index = (index - 1) / 2
	}
	h.array[index] = value
}

func (h *MaxHeap) String() string {
	res := "[ "

	for i := 0; i < h.count; i++ {
		res += fmt.Sprintf("%d ", h.array[i])
	}

	return res + "]"
}

func (h *MaxHeap) ParentOf(index int) int {
	if index <= 0 || index >= h.count {
		return -1
	}
	return (index - 1) / 2
}

func (h *MaxHeap) LeftChildOf(index int) int {
	left := 2*index + 1
	if left >= h.count {
		return -1
	}
	return left
}

func (h *MaxHeap) RightChildOf(index int) int {
	right := 2*index + 2
	if right >= h.count {
		return -1
	}
	return right
}

func (h *MaxHeap) GetMax() int {
	if h.count == 0 {
		return -1
	}
	return h.array[0]
}
