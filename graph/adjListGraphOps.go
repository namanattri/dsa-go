package graph

import "fmt"

func AdjListGraphOps() {
	g := NewAdjListGraph([]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'})

	fmt.Println("Empty directed graph")
	fmt.Println(g)

	g.CreateEdge('A', 'B')
	fmt.Println(g)
	g.CreateEdge('B', 'A')
	fmt.Println(g)
	g.CreateEdge('B', 'C')
	fmt.Println(g)
	g.CreateEdge('B', 'H')
	fmt.Println(g)
	g.CreateEdge('C', 'B')
	fmt.Println(g)
	g.CreateEdge('C', 'D')
	fmt.Println(g)
	g.CreateEdge('C', 'E')
	fmt.Println(g)
	g.CreateEdge('D', 'C')
	fmt.Println(g)
	g.CreateEdge('E', 'B')
	fmt.Println(g)
	g.CreateEdge('E', 'F')
	fmt.Println(g)
	g.CreateEdge('E', 'G')
	fmt.Println(g)
	g.CreateEdge('E', 'H')
	fmt.Println(g)
	g.CreateEdge('F', 'E')
	fmt.Println(g)
	g.CreateEdge('G', 'E')
	fmt.Println(g)
	g.CreateEdge('H', 'E')
	fmt.Println(g)
	g.CreateEdge('H', 'B')
	fmt.Println(g)

	g.TraverseByDFS()
	fmt.Println()
	g.TraverseByBFS()
	fmt.Println()

}
