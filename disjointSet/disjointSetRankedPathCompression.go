package disjointSet

import "fmt"

type RankedPathCompressionDisjointSetElement struct {
	value  int
	parent *RankedPathCompressionDisjointSetElement
	rank   int
}

func NewRankedPathCompressionDisjointSetElement(value int) *RankedPathCompressionDisjointSetElement {
	e := &RankedPathCompressionDisjointSetElement{value: value, rank: 1}
	e.parent = e
	return e
}

func (e *RankedPathCompressionDisjointSetElement) String() string {
	return fmt.Sprintf("{[%d]->(%d:%d)}", e.parent.value, e.value, e.rank)
}

type RankedPathCompressionDisjointSet struct{}

func NewRankedPathCompressionDisjointSet() *RankedPathCompressionDisjointSet {
	return &RankedPathCompressionDisjointSet{}
}

func (s *RankedPathCompressionDisjointSet) MakeSet(value int) *RankedPathCompressionDisjointSetElement {
	e := NewRankedPathCompressionDisjointSetElement(value)
	return e
}

func (s *RankedPathCompressionDisjointSet) Find(e *RankedPathCompressionDisjointSetElement) *RankedPathCompressionDisjointSetElement {
	for e.parent != e {
		e.parent = e.parent.parent
		e = e.parent
	}
	return e
}

func (s *RankedPathCompressionDisjointSet) FindRecursive(e *RankedPathCompressionDisjointSetElement) *RankedPathCompressionDisjointSetElement {
	if e.parent == e {
		return e
	}
	e.parent = s.FindRecursive(e.parent)
	return e.parent
}

func (s *RankedPathCompressionDisjointSet) Union(e1 *RankedPathCompressionDisjointSetElement, e2 *RankedPathCompressionDisjointSetElement) {
	setOfE1 := s.Find(e1)
	setOfE2 := s.Find(e2)

	if setOfE1 == setOfE2 {
		return
	}

	if setOfE1.rank < setOfE2.rank {
		setOfE1.parent = setOfE2
	} else if setOfE1.rank > setOfE2.rank {
		setOfE2.parent = setOfE1
	} else { // the rank is same make one the parent of another and increment rank of parent
		setOfE2.parent = setOfE1
		setOfE1.rank++
	}
}
