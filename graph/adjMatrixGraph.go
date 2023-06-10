package graph

import "fmt"

type AdjMatrixGraph struct {
	V            int
	E            int
	adj          [][]int
	vertexLabels []rune
}

func NewAdjMatrixGraph(V int, E int, vertexLabels []rune) *AdjMatrixGraph {
	g := &AdjMatrixGraph{V: V, E: E, vertexLabels: vertexLabels}
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
	res := " | "
	for u := 0; u < g.V; u++ {
		res += fmt.Sprintf("%c ", g.vertexLabels[u])
	}
	res += "|\n"

	for u := 0; u < g.V; u++ {
		row := fmt.Sprintf("%c| ", g.vertexLabels[u])
		for v := 0; v < g.V; v++ {
			row += fmt.Sprintf("%d ", g.adj[u][v])
		}
		res += row + "|\n"
	}
	return res
}
