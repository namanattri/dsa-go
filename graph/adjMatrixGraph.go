package graph

import "fmt"

type AdjMatrixGraph struct {
	V   int
	E   int
	adj [][]int
}

func NewAdjMatrixGraph(V int, E int) *AdjMatrixGraph {
	g := &AdjMatrixGraph{V: V, E: E}
	g.adj = make([][]int, V)
	for u := 0; u < V; u++ {
		g.adj[u] = make([]int, V)
		for v := 0; v < V; v++ {
			g.adj[u][v] = 0
		}
	}
	return g
}

func (g *AdjMatrixGraph) String() string {
	var res string
	for u := 0; u < g.V; u++ {
		row := "| "
		for v := 0; v < g.V; v++ {
			row += fmt.Sprintf("%d ", g.adj[u][v])
		}
		res += row + "|\n"
	}
	return res
}
