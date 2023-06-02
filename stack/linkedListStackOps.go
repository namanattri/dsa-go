package stack

import (
	"fmt"
)

func LinkedListStackOps() {
	s := NewLinkedListStack()
	fmt.Println(s)

	fmt.Printf("Trying to Peek\n")
	value, err := s.Peek()

	if err != nil {
		fmt.Printf("Error encountered while peeking: %v\n", err)
	} else {
		fmt.Printf("Peek: %d\n", value)
	}
	fmt.Println(s)

	// push multiple values to stack

	values := []int{45, 72, 83, 19, 61, 7, 92, 38, 51, 24, 86, 11}

	for _, value := range values {
		s.Push(value)
		fmt.Println(s)
	}

	fmt.Printf("Trying to Peek\n")
	value, err = s.Peek()

	if err != nil {
		fmt.Printf("Error encountered while peeking: %v\n", err)
	} else {
		fmt.Printf("Peek: %d\n", value)
	}
	fmt.Println(s)

	n := 5
	fmt.Printf("Popping %d elements\n", n)

	for i := 0; i < n; i++ {
		value, err := s.Pop()

		if err != nil {
			fmt.Printf("Error encountered while popping: %v\n", err)
		} else {
			fmt.Printf("Pop: %d\n", value)
		}
		fmt.Println(s)
	}

	fmt.Printf("Trying to Peek\n")
	value, err = s.Peek()

	if err != nil {
		fmt.Printf("Error encountered while peeking: %v\n", err)
	} else {
		fmt.Printf("Peek: %d\n", value)
	}
	fmt.Println(s)

	n = 8
	fmt.Printf("Popping %d elements\n", n)

	for i := 0; i < n; i++ {
		value, err := s.Pop()

		if err != nil {
			fmt.Printf("Error encountered while popping: %v\n", err)
		} else {
			fmt.Printf("Pop: %d\n", value)
		}
		fmt.Println(s)
	}
}
