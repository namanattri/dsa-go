package graph

import "fmt"

func AdjListGraphOps() {
	g := NewAdjListGraph([]rune{'A', 'B', 'C', 'D'})

	fmt.Println(g)
}
