package heaps

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
