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

// O(nlogn)
func (h *MaxHeap) Build(values []int) {
	for _, value := range values {
		h.Insert(value)
	}
}

// O(n)
func (h *MaxHeap) BuildOptimized(values []int) {
	n := len(values)
	for n > h.capacity {
		h.Resize()
	}

	copy(h.array[:n], values)
	h.count = n

	for i := h.ParentOf(h.count); i >= 0; i-- {
		h.PercolateDown(i)
	}
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

func (h *MaxHeap) PercolateDown(index int) {
	max := index
	left := h.LeftChildOf(index)
	right := h.RightChildOf(index)

	if left != -1 && h.array[left] > h.array[max] {
		max = left
	}

	if right != -1 && h.array[right] > h.array[max] {
		max = right
	}

	if max != index {
		// swap
		tmp := h.array[index]
		h.array[index] = h.array[max]
		h.array[max] = tmp

		h.PercolateDown(max)
	}
}

func (h *MaxHeap) DeleteMax() int {
	if h.count == 0 {
		return -1
	}
	value := h.array[0]
	h.array[0] = h.array[h.count-1]
	h.count--
	h.PercolateDown(0)
	return value
}
