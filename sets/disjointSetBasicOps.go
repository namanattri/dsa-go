package sets

import "fmt"

func DisjointSetBasicOps() {
	d := NewDisjointSetBasic()

	e1 := d.MakeSet(1)
	e2 := d.MakeSet(4)
	e3 := d.MakeSet(2)

	d.Union(e1, e2)
	d.Union(e3, e2)

	fmt.Printf("Parent of %v = %v and of %v = %v\n", e1, d.Find(e1), e3, d.Find(e3))
}
