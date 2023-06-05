package trees

import "fmt"

func BinarySearchTreeOps() {
	t := NewBinarySearchTree()

	t.Create()

	fmt.Printf("Finding %d result(recursively): %s\n", 9, t.FindRecursive(t.root, 9))
	fmt.Printf("Finding %d result(recursively): %s\n", 11, t.FindRecursive(t.root, 11))

	fmt.Printf("Finding %d result(non-recursively): %s\n", 9, t.FindNonRecursive(9))
	fmt.Printf("Finding %d result(non-recursively): %s\n", 11, t.FindNonRecursive(11))
}
