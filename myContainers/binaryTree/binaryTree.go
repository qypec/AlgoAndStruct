package binaryTree

/* binaryTree implementation */
type Tree struct {
	Root *Element
}

type Element struct {
	parent *Element
	ChildLeft *Element
	ChildRight *Element

	Value interface{}
}

func MakeParent(ChildLeft *Element, ChildRight *Element) *Element {
	var t Element

	t.parent = nil
	t.Value = nil
	t.ChildLeft = ChildLeft
	t.ChildRight = ChildRight
	return &t
}

func (t Element) HasChilds() bool {
	return t.ChildLeft != nil || t.ChildRight != nil
}

func (t Element) HasParent() bool {
	return t.parent != nil
}

func MakeElement(value interface{}) *Element {
	var t Element

	t.parent = nil
	t.ChildLeft = nil
	t.ChildRight = nil
	t.Value = value
	return &t
}

func MakeRoot(element *Element) *Tree {
	var t Tree

	t.Root = element
	return &t
}
