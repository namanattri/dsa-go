package trees

import "fmt"

func ExpressionTreeOps() {
	t := NewExpressionTree()

	t.CreateFromPostfixExpression([]rune{'A', 'B', 'C', '*', '+', 'D', '/'})

	fmt.Println("Inorder Traversal: ")
	t.InOrderNonRecursive()
}
