package tree

import (
	"fmt"
)

// Tree describes the node of a tree with almost two children.
type Tree struct {
	val      NodeValue
	parent   *Tree
	children []*Tree
}

// Val returns the value held by the current node of the tree.
func (t *Tree) Val() NodeValue {
	return t.val
}

// SetVal sets the value of the current node of the tree.
func (t *Tree) SetVal(n NodeValue) {
	t.val = n
}

// Parent returns a pointer to the parent of t.
// It also returns false if this node is the root of the tree or true otherwise.
// If this node is the root of the tree the p *Tree returned is equal to t *Tree
func (t *Tree) Parent() (p *Tree, ok bool) {
	if t.parent == nil {
		return t, false
	}
	return t.parent, true
}

// Children returns a slice of pointers to children of t.
func (t *Tree) Children() []*Tree {
	return t.children
}

// Child returns the i-th child of t.
func (t *Tree) Child(i int) (child *Tree, err error) {
	if i < 0 || i >= len(t.children) {
		return nil, fmt.Errorf("there is no child with index %d", i)
	}
	return t.children[i], nil
}

// AddChild adds a child to t with value n.
// Returns the child that has been added.
func (t *Tree) AddChild(n NodeValue) (tChild *Tree) {
	tChild = &Tree{val: n, parent: t}
	t.children = append(t.children, tChild)
	return
}

// NewTree is the default constructor for Tree.
func NewTree(val NodeValue) *Tree {
	return &Tree{val: val}
}

// Root returns a pointer to the root of the tree
func (t *Tree) Root() (root *Tree) {
	for root = t; root.parent != nil; root = root.parent {
	}
	return root
}

// String returns the string representation of the tree.
func (t *Tree) String() string {
	return stringify(t.Root()).String()
}
