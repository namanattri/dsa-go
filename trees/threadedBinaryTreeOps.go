package trees

import "fmt"

func ThreadedBinaryTreeOps() {
	t := NewThreadedBinaryTree()

	t.Create()

	t.InOrder()
	fmt.Println()
}
