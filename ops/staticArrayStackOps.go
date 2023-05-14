package ops

import (
	"dsa-go/stack"
	"fmt"
)

func StaticArrayStackOps() {
	s := stack.NewStaticArrayStack()
	printStackStats(s)

	fmt.Printf("Trying to Peek")
	value, err := s.Peek()

	if err != nil {
		fmt.Printf("Error encountered while peeking: %v\n", err)
	} else {
		fmt.Printf("Peek: %d\n", value)
	}

	// push multiple values to stack

	values := []int{45, 72, 83, 19, 61, 7, 92, 38, 51, 24, 86, 11}

	for _, value := range values {
		err := s.Push(value)

		if err != nil {
			fmt.Printf("Error while pushing %d: %v\n", value, err)
		} else {
			fmt.Printf("Pushed %d to stack\n", value)
		}
		printStackStats(s)
	}

	fmt.Printf("Trying to Peek")
	value, err = s.Peek()

	if err != nil {
		fmt.Printf("Error encountered while peeking: %v\n", err)
	} else {
		fmt.Printf("Peek: %d\n", value)
	}

	n := 5
	fmt.Printf("Popping %d elements\n", n)

	for i := 0; i < n; i++ {
		value, err := s.Pop()

		if err != nil {
			fmt.Printf("Error encountered while popping: %v\n", err)
		} else {
			fmt.Printf("Pop: %d\n", value)
		}
		printStackStats(s)
	}

	fmt.Printf("Trying to Peek")
	value, err = s.Peek()

	if err != nil {
		fmt.Printf("Error encountered while peeking: %v\n", err)
	} else {
		fmt.Printf("Peek: %d\n", value)
	}

	n = 8
	fmt.Printf("Popping %d elements\n", n)

	for i := 0; i < n; i++ {
		value, err := s.Pop()

		if err != nil {
			fmt.Printf("Error encountered while popping: %v\n", err)
		} else {
			fmt.Printf("Pop: %d\n", value)
		}
		printStackStats(s)
	}
}

func printStackStats(s *stack.StaticArrayStack) {
	fmt.Printf("Stack:: Top: %d, Size: %d, Capacity: %d, Values: %v\n", s.Top(), s.Size(), s.Capacity(), s.Array())
}
