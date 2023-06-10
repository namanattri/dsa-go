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
	adj          []*LinkedListNode // array of adj list pointers
	labelToIndex map[rune]int
}

func NewAdjListGraph(vertices []rune) *AdjListGraph {
	g := &AdjListGraph{}
	g.adj = make([]*LinkedListNode, len(vertices))
	g.labelToIndex = make(map[rune]int)

	for i, vertex := range vertices {
		g.adj[i] = NewLinkedListNode(vertex)
		g.labelToIndex[vertex] = i
	}
	return g
}

func (g *AdjListGraph) String() string {
	var res string
	for _, n := range g.adj {
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
	n := g.adj[g.labelToIndex[u]]

	for n.next != nil {
		n = n.next
	}

	n.next = NewLinkedListNode(v)
}
