package tree

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/m1gwings/treedrawer/drawer"
)

// NodeValue is the interface that wraps the Draw method.
//
// The Draw method allows to convert data into its unicode canvas representation.
// With the Draw method you can control how your data is going to appear on the tree.
type NodeValue interface {
	Draw() *drawer.Drawer
}

// NodeInt64 is the default type for drawing int64s on the tree.
type NodeInt64 int64

// Draw satisfies the NodeValue interface.
func (i NodeInt64) Draw() *drawer.Drawer {
	return NodeString(strconv.Itoa(int(i))).Draw()
}

// NodeString is the default type for drawing strings on the tree.
type NodeString string

// Draw satisfies the NodeValue interface.
func (s NodeString) Draw() *drawer.Drawer {
	lines := strings.Split(string(s), "\n")
	var maxLineLength int
	for _, line := range lines {
		realLineLength := utf8.RuneCountInString(line)
		if realLineLength > maxLineLength {
			maxLineLength = realLineLength
		}
	}
	d, err := drawer.NewDrawer(maxLineLength, len(lines))
	if err != nil {
		log.Fatal(fmt.Errorf("error while allocating new drawer in NodeString.Draw: %v", err))
	}
	for y, line := range lines {
		// x is decleared outside and incremented manually because in the for range it would
		// be incremented depending on bytes and not runes
		x := 0
		for _, r := range line {
			err := d.DrawRune(r, x, y)
			if err != nil {
				log.Fatal(fmt.Errorf("error while drawing %d th rune of %s line %d in NodeString.Draw() method: %v", x, s, y, err))
			}
			x++
		}
	}
	return d
}

// NodeFloat64 is the default type for drawing float64s on the tree.
type NodeFloat64 float64

// Draw satisfies the NodeValue interface.
func (f NodeFloat64) Draw() *drawer.Drawer {
	return NodeString(fmt.Sprintf("%v", f)).Draw()
}

// NodeComplex128 is the default type for drawing complex128s on the tree.
type NodeComplex128 complex128

// Draw satisfies the NodeValue interface.
func (z NodeComplex128) Draw() *drawer.Drawer {
	return NodeString(fmt.Sprintf("%v", z)).Draw()
}
