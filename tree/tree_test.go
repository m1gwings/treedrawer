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
	tr.AddChild(NodeInt64(3))
	tr.AddChild(NodeInt64(4))

	tr, err := tr.Child(1)
	if err != nil {
		t.Errorf("there should be a right child: %v", err)
	}
	v := tr.Val()
	if v != NodeInt64(4) {
		t.Errorf("the value in the right child should be 4, received %d", v)
	}

	tr, ok := tr.Parent()
	if !ok {
		t.Errorf("there should be a parent, expected true, received %t", ok)
	}
	v = tr.Val()
	if v != NodeInt64(5) {
		t.Errorf("the value in the root should be 5, received %d", v)
	}

	tr, err = tr.Child(0)
	if err != nil {
		t.Errorf("there should be a left child: %v", err)
	}
	v = tr.Val()
	if v != NodeInt64(3) {
		t.Errorf("the value in the left child should be 3, received %d", v)
	}

	fmt.Println(tr)
}

func TestRoot(t *testing.T) {
	tr := NewTree(NodeInt64(5))
	tr.AddChild(NodeInt64(4))
	tr, err := tr.Child(0)
	if err != nil {
		t.Errorf("there should be a child: %v", err)
	}
	tr.AddChild(NodeInt64(3))
	tr, err = tr.Child(0)
	if err != nil {
		t.Errorf("there should be a child: %v", err)
	}
	tr.AddChild(NodeInt64(2))
	tr, err = tr.Child(0)
	if err != nil {
		t.Errorf("there should be a child: %v", err)
	}
	tr.AddChild(NodeInt64(1))
	tr, err = tr.Child(0)
	if err != nil {
		t.Errorf("there should be a child: %v", err)
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

func TestParentBiggerThanBothChildren(t *testing.T) {
	tr := NewTree(NodeString("qwertyuiopasdfghjkl"))
	tr.AddChild(NodeString("sa"))
	tr.AddChild(NodeString("as"))

	fmt.Println(tr)
}

func TestNodeStringWithNewLine(t *testing.T) {
	tr := NewTree(NodeString("abcd\nab\nababab\n"))
	fmt.Println(tr)
}

func TestNodeWeird(t *testing.T) {
	rand.Seed(time.Now().Unix())

	fmt.Println(WeirdTree(5))
}

func WeirdTree(depth int) *Tree {
	t := NewTree(NodeWeird{})
	nChildren := rand.Intn(depth)
	for i := 0; i < nChildren; i++ {
		t.children = append(t.children, WeirdTree(depth-1))
		t.children[i].val = NodeWeird{}
	}
	return t
}
