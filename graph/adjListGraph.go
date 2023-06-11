package graph

import "fmt"

type LinkedListNode struct {
	vertex rune
	next   *LinkedListNode
}

func NewLinkedListNode(vertex rune) *LinkedListNode {
	return &LinkedListNode{vertex: vertex}
}

type AdjListGraph struct {
	vertices []rune
	adj      map[rune]*LinkedListNode // array of adj list pointers
	visited  map[rune]bool
}

func NewAdjListGraph(vertices []rune) *AdjListGraph {
	g := &AdjListGraph{}
	g.vertices = vertices
	g.adj = make(map[rune]*LinkedListNode)

	for _, vertex := range vertices {
		g.adj[vertex] = nil
	}
	return g
}

func (g *AdjListGraph) String() string {
	var res string
	for _, vertex := range g.vertices {
		res += fmt.Sprintf("%c: ", vertex)
		n := g.adj[vertex]
		for n != nil {
			res += fmt.Sprintf("%c -> ", n.vertex)
			n = n.next
		}
		res += "X\n"
	}
	return res
}

func (g *AdjListGraph) CreateEdge(u rune, v rune) {
	fmt.Printf("Creating edge (%c, %c)\n", u, v)
	newNode := NewLinkedListNode(v)
	n := g.adj[u]

	if n == nil {
		g.adj[u] = newNode
		return
	}

	for n.next != nil {
		n = n.next
	}

	n.next = newNode
}

func (g *AdjListGraph) TraverseByDFS() {
	g.visited = make(map[rune]bool)

	for _, vertex := range g.vertices {
		if !g.visited[vertex] {
			g.DFS(vertex)
		}
	}
}

func (g *AdjListGraph) DFS(vertex rune) {
	g.visited[vertex] = true
	fmt.Printf("%c ", vertex)

	cursor := g.adj[vertex]

	for cursor != nil {
		if !g.visited[cursor.vertex] {
			g.DFS(cursor.vertex)
		}
		cursor = cursor.next
	}
}
