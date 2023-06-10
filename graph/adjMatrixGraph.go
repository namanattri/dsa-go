package graph

import "fmt"

type AdjMatrixGraph struct {
	V            int
	E            int
	adj          [][]int
	vertexLabels []rune
	directed     bool
	visited      []bool
}

func NewAdjMatrixGraph(V int, E int, vertexLabels []rune, directed bool) *AdjMatrixGraph {
	g := &AdjMatrixGraph{V: V, E: E, vertexLabels: vertexLabels, directed: directed}
	g.adj = make([][]int, V)
	g.visited = make([]bool, V)
	for u := 0; u < V; u++ {
		g.adj[u] = make([]int, V)
		for v := 0; v < V; v++ {
			g.adj[u][v] = 0
			if u == v && !directed {
				g.adj[u][v] = 1
			}
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

func (g *AdjMatrixGraph) CreateEdge(u int, v int) {
	fmt.Printf("Creating edge (%d, %d)\n", u, v)
	g.adj[u][v] = 1

	if !g.directed {
		g.adj[v][u] = 1
	}
}

func (g *AdjMatrixGraph) DFSTraversal() {
	for i := 0; i < g.V; i++ {
		if !g.visited[i] {
			g.DFS(i)
		}
	}
}

func (g *AdjMatrixGraph) DFS(u int) {
	fmt.Printf("%c ", g.vertexLabels[u])
	g.visited[u] = true

	for v := 0; v < g.V; v++ {
		if !g.visited[v] && g.adj[u][v] == 1 {
			g.DFS(v)
		}
	}
}
