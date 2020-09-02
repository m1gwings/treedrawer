package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/m1gwings/treedrawer/tree"
	"golang.org/x/net/html"
)

func main() {
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}

	t := tree.NewTree(tree.NodeString(""))

	var f func(*html.Node, *tree.Tree)
	f = func(n *html.Node, t *tree.Tree) {
		t.SetVal(tree.NodeString(n.Data))
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			tChild := t.AddChild(tree.NodeString(""))
			f(c, tChild)
		}
	}
	// Starting from FirstChild because the DocumentRoot has an empty Val
	f(doc.FirstChild, t)

	fmt.Println(t)
}
