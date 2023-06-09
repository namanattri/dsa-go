package trees

import "fmt"

func BinarySearchTreeOps() {
	t := NewBinarySearchTree()

	t.Create()
	fmt.Println("Inorder Recursive Traversal: ")
	t.TraverseInorderRecursive(t.root)
	fmt.Println()
	fmt.Println("Inorder Non-Recursive Traversal: ")
	t.TraverseInorderNonRecursive()
	fmt.Println()

	fmt.Printf("Finding %d result(recursively): %s\n", 9, t.FindRecursive(t.root, 9))
	fmt.Printf("Finding %d result(recursively): %s\n", 11, t.FindRecursive(t.root, 11))

	fmt.Printf("Finding %d result(non-recursively): %s\n", 9, t.FindNonRecursive(9))
	fmt.Printf("Finding %d result(non-recursively): %s\n", 11, t.FindNonRecursive(11))

	fmt.Printf("Finding minimum(recursively): %s\n", t.FindMinimumRecursive(t.root))
	fmt.Printf("Finding minimum(non-recursively): %s\n", t.FindMinimumNonRecursive())

	fmt.Printf("Finding maximum(recursively): %s\n", t.FindMaximumRecursive(t.root))
	fmt.Printf("Finding maximum(non-recursively): %s\n", t.FindMaximumNonRecursive())

	fmt.Println("Inserting 11")
	t.InsertRecursive(t.root, 11)
	t.TraverseInorderNonRecursive()
	fmt.Println()

	fmt.Println("Inserting 20")
	t.InsertRecursive(t.root, 20)
	t.TraverseInorderNonRecursive()
	fmt.Println()

	fmt.Println("Inserting 2")
	t.InsertNonRecursive(2)
	t.TraverseInorderNonRecursive()
	fmt.Println()

	fmt.Println("Inserting 12")
	t.InsertNonRecursive(12)
	t.TraverseInorderNonRecursive()
	fmt.Println()

	fmt.Println("Delete 7")
	t.Delete(t.root, 7)
	t.TraverseInorderNonRecursive()
	fmt.Println()

	fmt.Println("Delete 10")
	t.Delete(t.root, 10)
	t.TraverseInorderNonRecursive()
	fmt.Println()
}
