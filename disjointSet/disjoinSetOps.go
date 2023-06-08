package disjointSet

import "fmt"

func DisjointSetOps() {
	d := NewDisjointSet(7)

	d.Make()

	fmt.Printf("d = %v\n", d)

	d.Union(5, 6)
	fmt.Printf("union(5, 6) d = %v\n", d)

	d.Union(1, 2)
	fmt.Printf("union(1, 2) d = %v\n", d)

	d.Union(0, 2)
	fmt.Printf("union(0, 2) d = %v\n", d)

	d2 := NewDisjointSet(7)

	d2.MakeBySize()

	fmt.Printf("d = %v\n", d2)

	d2.UnionBySize(5, 6)
	fmt.Printf("union(5, 6) d = %v\n", d2)

	d2.UnionBySize(1, 2)
	fmt.Printf("union(1, 2) d = %v\n", d2)

	d2.UnionBySize(0, 2)
	fmt.Printf("union(0, 2) d = %v\n", d2)

	d3 := NewDisjointSet(7)

	d3.MakeBySize()

	fmt.Printf("d = %v\n", d3)

	d3.UnionByHeight(5, 6)
	fmt.Printf("union(5, 6) d = %v\n", d3)

	d3.UnionByHeight(1, 2)
	fmt.Printf("union(1, 2) d = %v\n", d3)

	d3.UnionByHeight(0, 2)
	fmt.Printf("union(0, 2) d = %v\n", d3)
}
