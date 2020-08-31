package tree

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/m1gwings/treedrawer/drawer"
)

func TestTreeBuilding(t *testing.T) {
	tr := NewTree(NodeInt64(5))
	tr.AddLeft(NodeInt64(3))
	tr.AddRight(NodeInt64(4))

	tr, ok := tr.Right()
	if !ok {
		t.Errorf("there should be a right child, expected true, received %t", ok)
	}
	v := tr.Val()
	if v != NodeInt64(4) {
		t.Errorf("the value in the right child should be 4, received %d", v)
	}

	tr, ok = tr.Parent()
	if !ok {
		t.Errorf("there should be a parent, expected true, received %t", ok)
	}
	v = tr.Val()
	if v != NodeInt64(5) {
		t.Errorf("the value in the root should be 5, received %d", v)
	}

	tr, ok = tr.Left()
	if !ok {
		t.Errorf("there should be a left child, expected true, received %t", ok)
	}
	v = tr.Val()
	if v != NodeInt64(3) {
		t.Errorf("the value in the left child should be 3, received %d", v)
	}

	fmt.Println(tr)
}

func TestRoot(t *testing.T) {
	tr := NewTree(NodeInt64(5))
	tr.AddLeft(NodeInt64(4))
	tr, ok := tr.Left()
	if !ok {
		t.Errorf("there should be a left child, expected true, received %t", ok)
	}
	tr.AddLeft(NodeInt64(3))
	tr, ok = tr.Left()
	if !ok {
		t.Errorf("there should be a left child, expected true, received %t", ok)
	}
	tr.AddLeft(NodeInt64(2))
	tr, ok = tr.Left()
	if !ok {
		t.Errorf("there should be a left child, expected true, received %t", ok)
	}
	tr.AddLeft(NodeInt64(1))
	tr, ok = tr.Left()
	if !ok {
		t.Errorf("there should be a left child, expected true, received %t", ok)
	}

	root := tr.Root()
	v := root.Val()
	if v != NodeInt64(5) {
		t.Errorf("the value in the root should be 5, received %d", v)
	}

	fmt.Println(tr)
}

type NodeWeird struct{}

func (nW NodeWeird) Draw() *drawer.Drawer {
	h := rand.Intn(6) + 1
	w := rand.Intn(6) + 1
	d, err := drawer.NewDrawer(w, h)
	if err != nil {
		log.Fatal(fmt.Errorf("error while allocating new drawer in NodeWeird.Draw: %v", err))
	}
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			d.DrawRune('*', x, y)
		}
	}
	return d
}

func TestString(t *testing.T) {
	rand.Seed(time.Now().Unix())
	tr := NewTree(NodeWeird{})
	tr.AddLeft(NodeWeird{})
	tr.AddRight(NodeWeird{})

	fmt.Println(tr)
}
