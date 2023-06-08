package disjointSet

import "fmt"

type RankedDisjointSetElement struct {
	value  int
	parent *RankedDisjointSetElement
	rank   int
}

func NewRankedDisjointSetElement(value int) *RankedDisjointSetElement {
	return &RankedDisjointSetElement{value: value}
}

func (e *RankedDisjointSetElement) String() string {
	return fmt.Sprintf("{[%d]->(%d:%d)}", e.parent.value, e.value, e.rank)
}

type RankedDisjointSet struct{}

func NewRankedDisjointSet() *RankedDisjointSet {
	return &RankedDisjointSet{}
}

func (s *RankedDisjointSet) MakeSet(value int) *RankedDisjointSetElement {
	e := NewRankedDisjointSetElement(value)
	e.parent = e
	e.rank = 1
	return e
}

func (s *RankedDisjointSet) Find(e *RankedDisjointSetElement) *RankedDisjointSetElement {
	for e.parent != e {
		e = e.parent
	}
	return e
}

func (s *RankedDisjointSet) FindRecursive(e *RankedDisjointSetElement) *RankedDisjointSetElement {
	if e.parent == e {
		return e
	}
	return s.FindRecursive(e.parent)
}

func (s *RankedDisjointSet) Union(e1 *RankedDisjointSetElement, e2 *RankedDisjointSetElement) {
	setOfE1 := s.Find(e1)
	setOfE2 := s.Find(e2)

	if setOfE1 == setOfE2 {
		return // both sets already have same parents, and hence already unionised
	}

	if setOfE1.rank < setOfE2.rank {
		setOfE1.parent = setOfE2
	} else if setOfE1.rank > setOfE2.rank {
		setOfE2.parent = setOfE1
	} else {
		setOfE2.parent = setOfE1
		setOfE1.rank++
	}
}
