package trees

import "fmt"

func AVLTreeOps() {
	t := NewAVLTree()

	t.root = t.Insert(t.root, 12)
	traverse(t)

	// insert multiple nodes
	t.root = t.Insert(t.root, 14)
	traverse(t)

	t.root = t.Insert(t.root, 15)
	traverse(t)

	t.root = t.Insert(t.root, 16)
	t.root = t.Insert(t.root, 17)
	t.root = t.Insert(t.root, 18)
	traverse(t)

	t.root = t.Insert(t.root, 5)
	t.root = t.Insert(t.root, 8)
	t.root = t.Insert(t.root, 9)
	traverse(t)

	t.root = t.Insert(t.root, 3)
	t.root = t.Insert(t.root, 2)
	t.root = t.Insert(t.root, 1)
	traverse(t)

	fmt.Println("Height: ", t.root.height)
}

func traverse(t *AVLTree) {
	fmt.Print("Inorder: ")
	t.InOrder(t.root)
	fmt.Println()

	fmt.Print("Levelorder: ")
	t.LevelOrder()
	fmt.Println()
}
