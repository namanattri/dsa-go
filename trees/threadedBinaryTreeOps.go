package trees

import "fmt"

func ThreadedBinaryTreeOps() {
	t := NewThreadedBinaryTree()

	t.CreateInorderThreaded()

	fmt.Print("Inorder traversal: ")
	t.InOrderTraversalInOrderTree()
	fmt.Println()

	fmt.Print("Inorder traversal(alt): ")
	t.InOrderTraversalInOrderTreeAlt()
	fmt.Println()

	fmt.Print("Preorder traversal: ")
	t.PreOrderTraversalInOrderTree()
	fmt.Println()

	fmt.Print("Preorder traversal(alt): ")
	t.PreOrderTraversalInOrderTreeAlt()
	fmt.Println()

	t.InsertRightOfInOrderTree(t.root.left, 7)
	fmt.Printf("Inserting %d to the right of %d\n", 7, t.root.left.value)

	fmt.Print("Inorder traversal: ")
	t.InOrderTraversalInOrderTree()
	fmt.Println()

	fmt.Print("Inorder traversal(alt): ")
	t.InOrderTraversalInOrderTreeAlt()
	fmt.Println()

	fmt.Print("Preorder traversal: ")
	t.PreOrderTraversalInOrderTree()
	fmt.Println()

	fmt.Print("Preorder traversal(alt): ")
	t.PreOrderTraversalInOrderTreeAlt()
	fmt.Println()
}
