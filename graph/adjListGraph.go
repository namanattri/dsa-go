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
	adj map[rune]*LinkedListNode // array of adj list pointers
}

func NewAdjListGraph(vertices []rune) *AdjListGraph {
	g := &AdjListGraph{}
	g.adj = make(map[rune]*LinkedListNode)

	for _, vertex := range vertices {
		g.adj[vertex] = nil
	}
	return g
}

func (g *AdjListGraph) String() string {
	var res string
	for vertex, n := range g.adj {
		res += fmt.Sprintf("%c: ", vertex)
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
