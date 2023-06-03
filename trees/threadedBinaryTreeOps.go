package trees

import "fmt"

func ThreadedBinaryTreeOps() {
	t := NewThreadedBinaryTree()

	t.CreateInorderThreaded()

	t.InOrder()
	fmt.Println()

	t.InOrderAlt()
	fmt.Println()
}
