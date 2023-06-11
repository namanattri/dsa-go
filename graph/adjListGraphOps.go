package graph

import "fmt"

func AdjListGraphOps() {
	g := NewAdjListGraph([]rune{'A', 'B', 'C', 'D'})

	fmt.Println("Empty directed graph")
	fmt.Println(g)

	g.CreateEdge('A', 'B')
	fmt.Println(g)
	g.CreateEdge('A', 'D')
	fmt.Println(g)
	g.CreateEdge('B', 'C')
	fmt.Println(g)
	g.CreateEdge('C', 'A')
	fmt.Println(g)
	g.CreateEdge('C', 'D')
	fmt.Println(g)
}
