package trees

import "fmt"

func ThreadedBinaryTreeOps() {
	t := NewThreadedBinaryTree()

	t.CreateInorderThreaded()

	t.InOrderTraversalInOrderTree()
	fmt.Println()

	t.InOrderTraversalInOrderTreeAlt()
	fmt.Println()

	t.PreOrderTraversalInOrderTree()
	fmt.Println()

	t.PreOrderTraversalInOrderTreeAlt()
	fmt.Println()
}
