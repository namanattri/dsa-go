package graph

import (
	"fmt"
	"math"
)

type LinkedListNode struct {
	vertex rune
	weight int
	next   *LinkedListNode
}

func NewLinkedListNode(vertex rune, weight int) *LinkedListNode {
	return &LinkedListNode{vertex: vertex, weight: weight}
}

type AdjListGraph struct {
	vertices []rune
	adj      map[rune]*LinkedListNode // array of adj list pointers
	visited  map[rune]bool
	distance map[rune]int
	path     map[rune]rune
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

func (g *AdjListGraph) CreateEdge(u rune, v rune, weight int) {
	fmt.Printf("Creating edge (%c, %c)\n", u, v)
	newNode := NewLinkedListNode(v, weight)
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

func (g *AdjListGraph) TraverseByBFS() {
	g.visited = make(map[rune]bool)

	for _, vertex := range g.vertices {
		if !g.visited[vertex] {
			g.BFS(vertex)
		}
	}
}

func (g *AdjListGraph) BFS(vertex rune) {
	q := NewQueue()

	q.Enqueue(vertex)

	for !q.IsEmpty() {
		n, _ := q.Dequeue()

		vertex = n.value

		if !g.visited[vertex] {
			fmt.Printf("%c ", vertex)
			g.visited[vertex] = true
		}

		cursor := g.adj[vertex]

		for cursor != nil {
			if !g.visited[cursor.vertex] {
				q.Enqueue(cursor.vertex)
			}
			cursor = cursor.next
		}
	}
}

func (g *AdjListGraph) UnweightedShortestPathCalculation(source rune) {
	g.distance = make(map[rune]int)
	g.path = make(map[rune]rune)

	for _, vertex := range g.vertices {
		g.distance[vertex] = -1
	}
	g.distance[source] = 0

	q := NewQueue()

	q.Enqueue(source)

	for !q.IsEmpty() {
		v, _ := q.Dequeue()

		cursor := g.adj[v.value]

		for cursor != nil {
			if g.distance[cursor.vertex] == -1 {
				g.distance[cursor.vertex] = g.distance[v.value] + 1
				g.path[cursor.vertex] = v.value
				q.Enqueue(cursor.vertex)
			}
			cursor = cursor.next
		}
	}
}

func (g *AdjListGraph) DijkstrasShortestPathCalculation(source rune) {
	h := NewMinHeap()
	h.Insert(source, 0)

	g.distance = make(map[rune]int)
	g.path = make(map[rune]rune)

	for _, vertex := range g.vertices {
		g.distance[vertex] = -1
	}
	g.distance[source] = 0

	for !h.IsEmpty() {
		n, _ := h.DeleteMin()

		cursor := g.adj[n.value]

		for cursor != nil {
			d := g.distance[n.value] + cursor.weight
			if g.distance[cursor.vertex] == -1 {
				g.distance[cursor.vertex] = d
				h.Insert(cursor.vertex, d)
				g.path[cursor.vertex] = n.value
			}

			if g.distance[cursor.vertex] > d {
				g.distance[cursor.vertex] = d
				h.UpdatePriority(cursor.vertex, d)
			}
			cursor = cursor.next
		}
	}
}

func (g *AdjListGraph) BellmanFordAlgorithm(source rune) {
	g.distance = make(map[rune]int)
	g.path = make(map[rune]rune)

	for _, vertex := range g.vertices {
		g.distance[vertex] = math.MaxInt64
	}
	g.distance[source] = 0

	q := NewQueue()

	q.Enqueue(source)

	for !q.IsEmpty() {
		v, _ := q.Dequeue()

		cursor := g.adj[v.value]

		for cursor != nil {
			d := g.distance[v.value] + cursor.weight
			if g.distance[cursor.vertex] > d {
				g.distance[v.value] = g.distance[v.value] + cursor.weight
				g.path[cursor.vertex] = v.value

				if !q.Has(cursor.vertex) {
					q.Enqueue(cursor.vertex)
				}
			}
			cursor = cursor.next
		}
	}
}

func (g *AdjListGraph) DistanceTableFor() string {
	res := "Vertex\tDistance\tPathFrom\n"

	for _, vertex := range g.vertices {
		res += fmt.Sprintf("%c\t%d\t\t%c\n", vertex, g.distance[vertex], g.path[vertex])
	}

	return res
}
