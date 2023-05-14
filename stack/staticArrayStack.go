package stack

import "fmt"

const MAXSIZE = 10

type StaticArrayStack struct {
	top      int
	capacity int
	array    []int
}

func NewStaticArrayStack() *StaticArrayStack {
	return &StaticArrayStack{top: -1, capacity: MAXSIZE, array: make([]int, MAXSIZE)}
}

// main stack operations

func (s *StaticArrayStack) Push(value int) error {
	if s.IsFullStack() {
		err := fmt.Errorf("error: stack overflow")
		return err
	}

	s.top += 1
	s.array[s.top] = value
	return nil
}

func (s *StaticArrayStack) Pop() (int, error) {
	if s.IsEmptyStack() {
		err := fmt.Errorf("error: stack underflow")
		return 0, err
	}

	element := s.array[s.top]
	s.top--

	return element, nil
}

// auxiliary stack operations

func (s *StaticArrayStack) Peek() (int, error) {
	if s.IsEmptyStack() {
		err := fmt.Errorf("error: stack empty")
		return 0, err
	}

	return s.array[s.top], nil
}

func (s *StaticArrayStack) Top() int {
	return s.top
}

func (s *StaticArrayStack) Size() int {
	return s.top + 1
}

func (s *StaticArrayStack) Capacity() int {
	return s.capacity
}

func (s *StaticArrayStack) Array() []int {
	return s.array
}

func (s *StaticArrayStack) IsEmptyStack() bool {
	return s.top == -1
}

func (s *StaticArrayStack) IsFullStack() bool {
	return s.top == s.capacity-1
}

func (s *StaticArrayStack) String() string {
	return fmt.Sprintf("Stack:: Top: %d, Size: %d, Capacity: %d, Values: %v\n", s.Top(), s.Size(), s.Capacity(), s.Array())
}
