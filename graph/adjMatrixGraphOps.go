package graph

import "fmt"

func AdjMatrixGraphOps() {
	g := NewAdjMatrixGraph([]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'}, false)

	fmt.Println("Empty undirected graph")
	fmt.Println(g)

	g.CreateEdge('A', 'B')
	fmt.Println(g)
	g.CreateEdge('B', 'C')
	fmt.Println(g)
	g.CreateEdge('B', 'H')
	fmt.Println(g)
	g.CreateEdge('C', 'D')
	fmt.Println(g)
	g.CreateEdge('C', 'E')
	fmt.Println(g)
	g.CreateEdge('E', 'H')
	fmt.Println(g)
	g.CreateEdge('E', 'F')
	fmt.Println(g)
	g.CreateEdge('E', 'G')
	fmt.Println(g)

	g.DFSTraversal()
	fmt.Println()
	g.BFSTraversal()
	fmt.Println()
}
