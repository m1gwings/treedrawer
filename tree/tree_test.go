package tree

import (
	"fmt"
	"testing"
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

func TestBuiltInNodes(t *testing.T) {
	tr := NewTree(NodeInt64(1))
	tr.AddChild(NodeFloat64(1.5))
	tr.AddChild(NodeComplex128(1 + 1i))
	tr.AddChild(NodeString("string"))
	fmt.Println(tr)
}
