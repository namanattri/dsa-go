package trees

import "fmt"

func ExpressionTreeOps() {
	t := NewExpressionTree()

	t.CreateFromPostfixExpression([]rune{'A', 'B', 'C', '*', '+', 'D', '/'})

	fmt.Print("Preorder Traversal: ")
	t.PreOrderNonRecursive()
	fmt.Println()

	fmt.Print("Inorder Traversal: ")
	t.InOrderNonRecursive()
	fmt.Println()

	fmt.Print("Postorder Traversal: ")
	t.PostOrderNonRecursive()
	fmt.Println()
}
