package tree

import (
	"fmt"
	"log"
	"math"

	"github.com/m1gwings/treedrawer/drawer"
)

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

func stringify(t *Tree) *drawer.Drawer {
	dVal := t.val.Draw()
	dValW, dValH := dVal.Dimens()
	if t.left == nil && t.right == nil {
		d := drawer.NewDrawer(dValW+2, dValH)
		err := d.DrawDrawer(dVal, 1, 0)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing val with no child: %v", err))
		}
		return d
	}

	if (t.left != nil && t.right == nil) || (t.left == nil && t.right != nil) {
		var dChild *drawer.Drawer
		if t.left != nil {
			dChild = stringify(t.left)
		} else {
			dChild = stringify(t.right)
		}
		dChildW, dChildH := dChild.Dimens()
		w := int(math.Max(float64(dValW+2), float64(dChildW)))
		h := dValH + 1 + dChildH
		d := drawer.NewDrawer(w, h)
		err := d.DrawDrawer(dVal, (w-dValW)/2, 0)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing val with one child: %v", err))
		}
		err = d.DrawRune('│', w/2, dValH)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing | with one child: %v", err))
		}
		err = d.DrawDrawer(dChild, (w-dChildW)/2, dValH+1)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing child drawer with one child: %v", err))
		}
		return d
	}

	dLeft, dRight := stringify(t.left), stringify(t.right)
	dLeftW, dLeftH := dLeft.Dimens()
	dRightW, dRightH := dRight.Dimens()
	maxChildW := int(math.Max(float64(dLeftW), float64(dRightW)))
	w := maxChildW*2 + 1
	maxChildH := int(math.Max(float64(dLeftH), float64(dRightH)))
	h := dValH + 1 + maxChildH
	d := drawer.NewDrawer(w, h)
	err := d.DrawDrawer(dVal, (w-dValW)/2, 0)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing val with two children: %v", err))
	}
	err = d.DrawRune('┴', w/2, dValH)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing ┴ rune with two childern: %v", err))
	}
	for i := 1; i <= maxChildW/2; i++ {
		err = d.DrawRune('─', w/2-i, dValH)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing ─ rune with two childern with negative i: %v", err))
		}
		err = d.DrawRune('─', w/2+i, dValH)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing ─ rune with two childern with positive i: %v", err))
		}
	}
	err = d.DrawRune('╭', w/2-maxChildW/2-1, dValH)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing ╭ rune with two children: %v", err))
	}
	err = d.DrawRune('╮', w/2+maxChildW/2+1, dValH)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing ╮ rune with two children: %v", err))
	}
	err = d.DrawDrawer(dLeft, (maxChildW-dLeftW)/2, dValH+1)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing left child: %v", err))
	}
	err = d.DrawDrawer(dRight, maxChildW+1+(maxChildW-dRightW)/2, dValH+1)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing right child: %v", err))
	}
	return d
}
