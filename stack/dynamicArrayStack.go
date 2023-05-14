package stack

import "fmt"

const INITIAL_SIZE = 1

type DynamicArrayStack struct {
	top      int
	capacity int
	array    []int
}

func NewDynamicArrayStack() *DynamicArrayStack {
	return &DynamicArrayStack{top: -1, capacity: INITIAL_SIZE, array: make([]int, INITIAL_SIZE)}
}

func (s *DynamicArrayStack) doubleStack() {
	s.array = append(s.array, make([]int, s.capacity)...)
	s.capacity *= 2
}

// main stack operations

func (s *DynamicArrayStack) Push(value int) error {
	if s.IsFullStack() {
		s.doubleStack()
	}

	s.top += 1
	s.array[s.top] = value
	return nil
}

func (s *DynamicArrayStack) Pop() (int, error) {
	if s.IsEmptyStack() {
		err := fmt.Errorf("error: stack underflow")
		return 0, err
	}

	element := s.array[s.top]
	s.top--

	return element, nil
}

// auxiliary stack operations

func (s *DynamicArrayStack) Peek() (int, error) {
	if s.IsEmptyStack() {
		err := fmt.Errorf("error: stack empty")
		return 0, err
	}

	return s.array[s.top], nil
}

func (s *DynamicArrayStack) Top() int {
	return s.top
}

func (s *DynamicArrayStack) Size() int {
	return s.top + 1
}

func (s *DynamicArrayStack) Capacity() int {
	return s.capacity
}

func (s *DynamicArrayStack) Array() []int {
	return s.array
}

func (s *DynamicArrayStack) IsEmptyStack() bool {
	return s.top == -1
}

func (s *DynamicArrayStack) IsFullStack() bool {
	return s.top == s.capacity-1
}

func (s *DynamicArrayStack) String() string {
	return fmt.Sprintf("Stack:: Top: %d, Size: %d, Capacity: %d, Values: %v\n", s.Top(), s.Size(), s.Capacity(), s.Array())
}
