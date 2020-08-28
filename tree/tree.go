package tree

import (
	"fmt"
	"log"
	"math"
	random "math/rand"
	"strconv"
	"time"

	"github.com/m1gwings/treedrawer/drawer"
)

func init() {
	random.Seed(time.Now().Unix())
}

// Tree describes the node of a tree with atmost two children.
type Tree struct {
	val                 int64
	left, right, parent *Tree
}

// Val returns the value held by the current node of the tree.
func (t *Tree) Val() int64 {
	return t.val
}

// Left moves the current node to its left child.
// returns false if there is no left child, otherwise it returns true.
func (t *Tree) Left() (ok bool) {
	if t.left == nil {
		return false
	}
	t = t.left
	return true
}

// Right moves the current node to its right child.
// returns false if there is no right child, otherwise it returns true.
func (t *Tree) Right() (ok bool) {
	if t.right == nil {
		return false
	}
	t = t.right
	return true
}

// Parent moves the current node to its parent child.
// returns false if this node is the root of the whole tree, otherwise it returns true.
func (t *Tree) Parent() (ok bool) {
	if t.parent == nil {
		return false
	}
	t = t.parent
	return true
}

// AddLeft adds a left child to the current node which will held val.
func (t *Tree) AddLeft(val int64) {
	t.left = &Tree{val: val, parent: t}
}

// AddRight adds a right child to the current node which will held val.
func (t *Tree) AddRight(val int64) {
	t.right = &Tree{val: val, parent: t}
}

// Rand returns the root of a random three with at most n layers.
func Rand(n int) *Tree {
	t := new(Tree)
	rand(t, 0, n-1)
	return t
}

func rand(t *Tree, curr, maxRecursion int) {
	t.val = random.Int63n(100)
	if curr == maxRecursion {
		return
	}
	if random.Int()%2 == 1 {
		t.AddLeft(0)
		rand(t.left, curr+1, maxRecursion)
	}
	if random.Int()%2 == 1 {
		t.AddRight(0)
		rand(t.right, curr+1, maxRecursion)
	}
}

// String returns the string representation of the tree.
func (t *Tree) String() string {
	return stringify(t).String()
}

func stringify(t *Tree) *drawer.Drawer {
	dVal := drawer.NewDrawerFromString(strconv.Itoa(int(t.val)))
	dValW, _ := dVal.Dimens()
	if t.left == nil && t.right == nil {
		d := drawer.NewDrawer(dValW+2, 1)
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
		h := dChildH + 2
		d := drawer.NewDrawer(w, h)
		err := d.DrawDrawer(dVal, (w-dValW)/2, 0)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing val with one child: %v", err))
		}
		err = d.DrawByte('|', w/2, 1)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing | with one child: %v", err))
		}
		err = d.DrawDrawer(dChild, (w-dChildW)/2, 2)
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
	slashEndI := maxChildW/2 + 1
	edgeH := maxChildW - slashEndI + 1
	h := maxChildH + edgeH + 1
	d := drawer.NewDrawer(w, h)
	err := d.DrawDrawer(dVal, (w-dValW)/2, 0)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing val with two children: %v", err))
	}
	for i := 0; i < edgeH; i++ {
		err = d.DrawByte('/', w/2-1-i, i+1)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing / with two children: %v", err))
		}
		err = d.DrawByte('\\', w/2+1+i, i+1)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing \\ with two children: %v", err))
		}
	}
	d.DrawDrawer(dLeft, (maxChildW-dLeftW)/2, edgeH)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing left child: %v", err))
	}
	d.DrawDrawer(dRight, maxChildW+1+(maxChildW-dRightW)/2, edgeH)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing right child: %v", err))
	}
	return d
}
