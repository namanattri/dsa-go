package graph

import "fmt"

func AdjListGraphOps() {
	g := NewAdjListGraph([]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'})

	fmt.Println("Empty directed graph")
	fmt.Println(g)

	edges := [][]rune{
		{'A', 'B'},
		{'B', 'A'},
		{'B', 'C'},
		{'B', 'H'},
		{'C', 'B'},
		{'C', 'D'},
		{'C', 'E'},
		{'D', 'C'},
		{'E', 'B'},
		{'E', 'F'},
		{'E', 'G'},
		{'E', 'H'},
		{'F', 'E'},
		{'G', 'E'},
		{'H', 'E'},
		{'H', 'B'},
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

	edges = [][]rune{
		{'A', 'B'},
		{'A', 'D'},
		{'B', 'D'},
		{'B', 'E'},
		{'C', 'A'},
		{'C', 'F'},
		{'D', 'F'},
		{'D', 'G'},
		{'E', 'G'},
		{'G', 'F'},
	}

	createEdges(g2, edges)

	fmt.Println(g)

	g2.UnweightedShortestPathCalculation('C')
	fmt.Println(g2.DistanceTableFor())
}

func createEdges(g *AdjListGraph, edges [][]rune) {
	for _, edge := range edges {
		g.CreateEdge(edge[0], edge[1])
	}
}
