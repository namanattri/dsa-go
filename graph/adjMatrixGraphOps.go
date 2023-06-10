package graph

import "fmt"

func AdjMatrixGraphOps() {
	g := NewAdjMatrixGraph(4, 5, []rune{'A', 'B', 'C', 'D'})

	fmt.Println(g)
}
