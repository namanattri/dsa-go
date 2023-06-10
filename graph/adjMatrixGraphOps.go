package graph

import "fmt"

func AdjMatrixGraphOps() {
	g1 := NewAdjMatrixGraph(4, 5, []rune{'A', 'B', 'C', 'D'}, true)

	fmt.Println("Empty undirected graph")
	fmt.Println(g1)

	g1.CreateEdge(0, 1)
	fmt.Println(g1)
	g1.CreateEdge(0, 3)
	fmt.Println(g1)
	g1.CreateEdge(1, 2)
	fmt.Println(g1)
	g1.CreateEdge(2, 0)
	fmt.Println(g1)
	g1.CreateEdge(2, 3)
	fmt.Println(g1)

	fmt.Println(g1)
}
