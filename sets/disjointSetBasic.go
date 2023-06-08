package sets

type BasicElement struct {
	parent *BasicElement
}

func NewDisjointSetElement() *BasicElement {
	return &BasicElement{}
}

type DisjointSetFastUnionSlowFind struct{}

func NewDisjointSetFastUnionSlowFind() *DisjointSetFastUnionSlowFind {
	return &DisjointSetFastUnionSlowFind{}
}

func (s *DisjointSetFastUnionSlowFind) MakeSet() *BasicElement {
	e := NewDisjointSetElement()
	e.parent = e
	return e
}

func (s *DisjointSetFastUnionSlowFind) Find(e *BasicElement) *BasicElement {
	for e.parent != e {
		e = e.parent
	}
	return e
}

func (s *DisjointSetFastUnionSlowFind) FindRecursive(e *BasicElement) *BasicElement {
	if e == e.parent {
		return e
	}
	return s.FindRecursive(e.parent)
}

func (s *DisjointSetFastUnionSlowFind) Union(e1 *BasicElement, e2 *BasicElement) {
	e1.parent = e2
}
