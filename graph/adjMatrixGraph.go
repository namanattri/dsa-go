package graph

import "fmt"

type AdjMatrixGraph struct {
	adj          [][]int
	vertices     []rune
	directed     bool
	visited      []bool
	labelToIndex map[rune]int
}

func NewAdjMatrixGraph(vertices []rune, directed bool) *AdjMatrixGraph {
	g := &AdjMatrixGraph{vertices: vertices, directed: directed}

	countOfVertices := len(vertices)

	g.adj = make([][]int, countOfVertices)
	g.visited = make([]bool, countOfVertices)
	g.labelToIndex = make(map[rune]int)

	for u := 0; u < countOfVertices; u++ {
		g.adj[u] = make([]int, countOfVertices)
		g.labelToIndex[g.vertices[u]] = u

		for v := 0; v < countOfVertices; v++ {
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
	countOfVertices := len(g.vertices)

	for u := 0; u < countOfVertices; u++ {
		res += fmt.Sprintf("%c ", g.vertices[u])
	}
	res += "|\n"

	for u := 0; u < countOfVertices; u++ {
		row := fmt.Sprintf("%c| ", g.vertices[u])
		for v := 0; v < countOfVertices; v++ {
			row += fmt.Sprintf("%d ", g.adj[u][v])
		}
		res += row + "|\n"
	}
	return res
}

func (g *AdjMatrixGraph) CreateEdge(u rune, v rune) {
	fmt.Printf("Creating edge (%c, %c)\n", u, v)
	g.adj[g.labelToIndex[u]][g.labelToIndex[v]] = 1

	if !g.directed {
		g.adj[g.labelToIndex[v]][g.labelToIndex[u]] = 1
	}
}

func (g *AdjMatrixGraph) DFSTraversal() {
	countOfVertices := len(g.vertices)
	for i := 0; i < countOfVertices; i++ {
		if !g.visited[i] {
			g.DFS(i)
		}
	}
}

func (g *AdjMatrixGraph) DFS(u int) {
	fmt.Printf("%c ", g.vertices[u])
	g.visited[u] = true
	countOfVertices := len(g.vertices)

	for v := 0; v < countOfVertices; v++ {
		if !g.visited[v] && g.adj[u][v] == 1 {
			g.DFS(v)
		}
	}
}
