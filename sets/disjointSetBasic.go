package sets

import "fmt"

type BasicElement struct {
	parent *BasicElement
	value  int
}

func NewBasicElement(value int) *BasicElement {
	return &BasicElement{value: value}
}

func (e *BasicElement) String() string {
	return fmt.Sprintf("%d ", e.value)
}

type DisjointSetBasic struct{}

func NewDisjointSetBasic() *DisjointSetBasic {
	return &DisjointSetBasic{}
}

func (s *DisjointSetBasic) MakeSet(value int) *BasicElement {
	e := NewBasicElement(value)
	e.parent = e
	return e
}

func (s *DisjointSetBasic) Find(e *BasicElement) *BasicElement {
	for e.parent != e {
		e = e.parent
	}
	return e
}

func (s *DisjointSetBasic) FindRecursive(e *BasicElement) *BasicElement {
	if e == e.parent {
		return e
	}
	return s.FindRecursive(e.parent)
}

func (s *DisjointSetBasic) Union(e1 *BasicElement, e2 *BasicElement) {
	e1.parent = e2
}
