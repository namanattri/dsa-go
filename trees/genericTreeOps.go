package trees

import "fmt"

func GenericTreeOps() {
	t := NewGenericTree()

	t.Create()

	t.LevelOrder()
	fmt.Println()
}
