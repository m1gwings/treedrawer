package tree

import (
	"fmt"
	"log"
	"testing"
)

func TestShowcase(t *testing.T) {
	tr := NewTree(NodeInt64(9))
	tr.AddChild(NodeString("I can handle strings"))
	tr.AddChild(NodeInt64(1))
	tr.AddChild(NodeInt64(2))
	tr.AddChild(NodeInt64(3))
	tr.AddChild(NodeInt64(4))
	tr, err := tr.Child(1)
	if err != nil {
		log.Fatal(err)
	}
	tr.AddChild(NodeInt64(124))
	tr.AddChild(NodeInt64(13))
	tr.AddChild(NodeString("a string"))
	tr, ok := tr.Parent()
	if !ok {
		log.Fatal(fmt.Errorf("this child should have a parent"))
	}
	tr, err = tr.Child(0)
	if err != nil {
		log.Fatal(err)
	}
	tr.AddChild(NodeString("with as many children as you want"))
	tr, err = tr.Child(0)
	if err != nil {
		log.Fatal(err)
	}
	tr.AddChild(NodeString("with as many layers as you want"))
	tr, err = tr.Child(0)
	if err != nil {
		log.Fatal(err)
	}
	tr.AddChild(NodeString("actually I can handle everything..."))
	tr, err = tr.Child(0)
	if err != nil {
		log.Fatal(err)
	}
	tr.AddChild(NodeString("...that satisfies NodeValue interface"))

	fmt.Println(tr)
}
