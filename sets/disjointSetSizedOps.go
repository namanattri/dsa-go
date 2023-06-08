package sets

import "fmt"

func DisjointSetSizedOps() {
	d := NewSizedDisjointSet()

	e0 := d.MakeSet(0)
	e1 := d.MakeSet(1)
	e2 := d.MakeSet(2)
	e3 := d.MakeSet(3)
	e4 := d.MakeSet(4)
	e5 := d.MakeSet(5)
	e6 := d.MakeSet(6)

	d.Union(e5, e6)
	d.Union(e1, e2)
	d.Union(e0, e2)
	d.Union(e3, e1)
	d.Union(e4, e5)

	fmt.Println(e0)
	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println(e3)
	fmt.Println(e4)
	fmt.Println(e5)
	fmt.Println(e6)
}
