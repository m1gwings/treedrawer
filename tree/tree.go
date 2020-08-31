package tree

// Tree describes the node of a tree with atmost two children.
type Tree struct {
	val                 NodeValue
	left, right, parent *Tree
}

// Val returns the value held by the current node of the tree.
func (t *Tree) Val() NodeValue {
	return t.val
}

// Left returns a pointer to the left child of t.
// It also returns true if the child exists or false otherwise.
// If the child doesn't exist the l *Tree returned is equal to t *Tree
func (t *Tree) Left() (l *Tree, ok bool) {
	if t.left == nil {
		return t, false
	}
	return t.left, true
}

// Right returns a pointer to the right child of t.
// It also returns true if the child exists or false otherwise.
// If the child doesn't exist the r *Tree returned is equal to t *Tree
func (t *Tree) Right() (r *Tree, ok bool) {
	if t.right == nil {
		return t, false
	}
	return t.right, true
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

// NewTree is the default constructor for Tree.
func NewTree(val NodeValue) *Tree {
	return &Tree{val: val}
}

// AddLeft adds a left child to the current node which will held val.
func (t *Tree) AddLeft(val NodeValue) {
	t.left = &Tree{val: val, parent: t}
}

// AddRight adds a right child to the current node which will held val.
func (t *Tree) AddRight(val NodeValue) {
	t.right = &Tree{val: val, parent: t}
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
