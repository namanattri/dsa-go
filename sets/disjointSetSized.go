package sets

import "fmt"

type SizedDisjointSetElement struct {
	value  int
	parent *SizedDisjointSetElement
	size   int
}

func NewSizedDisjointSetElement(value int) *SizedDisjointSetElement {
	return &SizedDisjointSetElement{value: value}
}

func (e *SizedDisjointSetElement) String() string {
	return fmt.Sprintf("{[%d]->(%d:%d)}", e.parent.value, e.value, e.size)
}

type SizedDisjointSet struct{}

func NewSizedDisjointSet() *SizedDisjointSet {
	return &SizedDisjointSet{}
}

func (s *SizedDisjointSet) MakeSet(value int) *SizedDisjointSetElement {
	e := NewSizedDisjointSetElement(value)
	e.parent = e
	e.size = 1
	return e
}

func (s *SizedDisjointSet) Find(e *SizedDisjointSetElement) *SizedDisjointSetElement {
	for e.parent != e {
		e = e.parent
	}
	return e
}

func (s *SizedDisjointSet) FindRecursive(e *SizedDisjointSetElement) *SizedDisjointSetElement {
	if e.parent == e {
		return e
	}
	return s.FindRecursive(e.parent)
}

func (s *SizedDisjointSet) Union(e1 *SizedDisjointSetElement, e2 *SizedDisjointSetElement) {
	setOfE1 := s.Find(e1)
	setOfE2 := s.Find(e2)

	if setOfE1 == setOfE2 {
		return // both sets already have same parents, and hence already unionised
	}

	if setOfE1.size <= setOfE2.size {
		setOfE1.parent = setOfE2
		setOfE2.size += setOfE1.size
	} else {
		setOfE2.parent = setOfE1
		setOfE1.size += setOfE2.size
	}
}
