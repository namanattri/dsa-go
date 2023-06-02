package trees

import "fmt"

func BinaryTreeOps() {
	t := NewBinaryTree()

	t.Create()

	t.PreOrder(t.root)
	fmt.Println()
	t.PreOrderNonRecursive()
	fmt.Println()

	t.InOrder(t.root)
	fmt.Println()
	t.InOrderNonRecursive()
	fmt.Println()
}
