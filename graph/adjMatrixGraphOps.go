package graph

import "fmt"

func AdjMatrixGraphOps() {
	g := NewAdjMatrixGraph(4, 5, []rune{'A', 'B', 'C', 'D'}, true)

	fmt.Println("Empty undirected graph")
	fmt.Println(g)

	g.CreateEdge(0, 1)
	fmt.Println(g)
	g.CreateEdge(0, 3)
	fmt.Println(g)
	g.CreateEdge(1, 2)
	fmt.Println(g)
	g.CreateEdge(2, 0)
	fmt.Println(g)
	g.CreateEdge(2, 3)
	fmt.Println(g)

	fmt.Println(g)

	g.DFSTraversal()
}
