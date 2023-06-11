package graph

import "fmt"

type AdjMatrixGraph struct {
	adj      map[rune]map[rune]int
	vertices []rune
	directed bool
	visited  map[rune]bool
}

func NewAdjMatrixGraph(vertices []rune, directed bool) *AdjMatrixGraph {
	g := &AdjMatrixGraph{vertices: vertices, directed: directed}

	g.adj = make(map[rune]map[rune]int)

	for _, u := range vertices {
		g.adj[u] = make(map[rune]int)

		for _, v := range vertices {
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

	for _, u := range g.vertices {
		res += fmt.Sprintf("%c ", u)
	}
	res += "|\n"

	for _, u := range g.vertices {
		row := fmt.Sprintf("%c| ", u)
		for _, v := range g.vertices {
			row += fmt.Sprintf("%d ", g.adj[u][v])
		}
		res += row + "|\n"
	}
	return res
}

func (g *AdjMatrixGraph) CreateEdge(u rune, v rune) {
	fmt.Printf("Creating edge (%c, %c)\n", u, v)
	g.adj[u][v] = 1

	if !g.directed {
		g.adj[v][u] = 1
	}
}

func (g *AdjMatrixGraph) DFSTraversal() {
	g.visited = make(map[rune]bool)

	for _, vertex := range g.vertices {
		if !g.visited[vertex] {
			g.DFS(vertex)
		}
	}
}

func (g *AdjMatrixGraph) DFS(u rune) {
	fmt.Printf("%c ", u)
	g.visited[u] = true

	for _, v := range g.vertices {
		if !g.visited[v] && g.adj[u][v] == 1 {
			g.DFS(v)
		}
	}
}
