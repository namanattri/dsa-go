package graph

import "fmt"

type edge struct {
	from   rune
	to     rune
	weight int
}

func AdjListGraphOps() {
	g := NewAdjListGraph([]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'})

	fmt.Println("Empty directed graph")
	fmt.Println(g)

	edges := []edge{
		{'A', 'B', 1},
		{'B', 'A', 1},
		{'B', 'C', 1},
		{'B', 'H', 1},
		{'C', 'B', 1},
		{'C', 'D', 1},
		{'C', 'E', 1},
		{'D', 'C', 1},
		{'E', 'B', 1},
		{'E', 'F', 1},
		{'E', 'G', 1},
		{'E', 'H', 1},
		{'F', 'E', 1},
		{'G', 'E', 1},
		{'H', 'E', 1},
		{'H', 'B', 1},
	}

	createEdges(g, edges)

	fmt.Println(g)

	fmt.Print("DFS: ")
	g.TraverseByDFS()
	fmt.Println()

	fmt.Print("BFS: ")
	g.TraverseByBFS()
	fmt.Println()

	g2 := NewAdjListGraph([]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G'})

	edges = []edge{
		{'A', 'B', 1},
		{'A', 'D', 1},
		{'B', 'D', 1},
		{'B', 'E', 1},
		{'C', 'A', 1},
		{'C', 'F', 1},
		{'D', 'F', 1},
		{'D', 'G', 1},
		{'E', 'G', 1},
		{'G', 'F', 1},
	}

	createEdges(g2, edges)

	fmt.Println(g2)

	g2.UnweightedShortestPathCalculation('C')
	fmt.Println(g2.DistanceTableFor())

	g3 := NewAdjListGraph([]rune{'A', 'B', 'C', 'D', 'E'})

	edges = []edge{
		{'A', 'B', 4},
		{'A', 'C', 1},
		{'B', 'E', 4},
		{'C', 'B', 2},
		{'C', 'D', 4},
		{'D', 'E', 4},
	}

	createEdges(g3, edges)

	fmt.Println(g3)

	g3.DijkstrasShortestPathCalculation('A')
	fmt.Println(g3.DistanceTableFor())
}

func createEdges(g *AdjListGraph, edges []edge) {
	for _, edge := range edges {
		g.CreateEdge(edge.from, edge.to, edge.weight)
	}
}
