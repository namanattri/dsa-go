package sets

import "fmt"

type DisjointSet struct {
	array []int
	size  int
}

func NewDisjointSet(size int) *DisjointSet {
	return &DisjointSet{array: make([]int, size), size: size}
}

func (d *DisjointSet) Make() {
	for i := 0; i < d.size; i++ {
		d.array[i] = i
	}
}

func (d *DisjointSet) Find(x int) int {
	if x < 0 || x >= d.size {
		return -1
	}

	if d.array[x] == x {
		return x
	} else {
		return d.Find(d.array[x])
	}
}

func (d *DisjointSet) Union(a int, b int) {
	if d.Find(a) == d.Find(b) {
		return // do nothing: already in set
	}

	if a < 0 || a >= d.size || b < 0 || b >= d.size {
		return // either or both a, b don't exist in set
	}

	// set the root of a to b
	d.array[a] = b
}

func (d *DisjointSet) String() string {
	res := "[ "

	for index, value := range d.array {
		res += fmt.Sprintf("(%d:%d) ", index, value)
	}

	return res + "]"
}

func (d *DisjointSet) MakeBySize() {
	for i := 0; i < d.size; i++ {
		d.array[i] = -1
	}
}

func (d *DisjointSet) FindBySize(x int) int {
	if x < 0 || x >= d.size {
		return -1
	}

	if d.array[x] == -1 {
		return x
	} else {
		return d.Find(d.array[x])
	}
}

func (d *DisjointSet) UnionBySize(a int, b int) {
	if d.FindBySize(a) == d.FindBySize(b) && d.FindBySize(a) != -1 {
		return
	}

	if d.array[a] < -d.array[b] {
		d.array[b] += d.array[a]
		d.array[a] = b
	} else {
		d.array[a] += d.array[b]
		d.array[b] = a
	}
}

// @todo array implementation not working
func (d *DisjointSet) UnionByHeight(a int, b int) {
	if d.FindBySize(a) == d.FindBySize(b) && d.FindBySize(a) != -1 {
		return
	}

	if d.array[a] < -d.array[b] {
		d.array[a] = b
	} else {
		if d.array[b] == d.array[a] {
			d.array[a] = b
			d.array[b]--
		}
	}
}
