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

	t.PostOrder(t.root)
	fmt.Println()
	t.PostOrderNonRecursive()
	fmt.Println()

	t.LevelOrder()
	fmt.Println()
}
