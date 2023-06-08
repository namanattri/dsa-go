package sets

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
}