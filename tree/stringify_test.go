package tree

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/m1gwings/treedrawer/drawer"
)

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

// countNodes returns the total number of nodes in a tree with l layers
// each node has exactly c children except for leaf nodes
func countNodes(l, c int) (count int) {
	for i := 0; i < l; i++ {
		count += int(math.Pow(float64(c), float64(i)))
	}
	return
}

func benchmarkDrawing(layers, nChildren int, b *testing.B) {
	value := "*"

	t := NewTree(NodeString(value))
	var addChildren func(*Tree, int, int)
	addChildren = func(t *Tree, currentLayer, lastLayer int) {
		if currentLayer == lastLayer {
			return
		}
		for c := 0; c < nChildren; c++ {
			addChildren(t.AddChild(NodeString(value)), currentLayer+1, lastLayer)
		}
	}
	addChildren(t, 0, layers-1)
	devNull, err := os.OpenFile(os.DevNull, os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(fmt.Errorf("couldn't open %s to simulate printing: %v", os.DevNull, err))
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		fmt.Fprintf(devNull, "%v\n", t)
	}
	b.ReportMetric(float64(layers), "layers")
	b.ReportMetric(float64(nChildren), "children")
	b.ReportMetric(float64(countNodes(layers, nChildren)), "nodes")
}

func BenchmarkDrawing3L3C(b *testing.B)    { benchmarkDrawing(3, 3, b) }
func BenchmarkDrawing100L1C(b *testing.B)  { benchmarkDrawing(100, 1, b) }
func BenchmarkDrawing6L3C(b *testing.B)    { benchmarkDrawing(6, 3, b) }
func BenchmarkDrawing1000L1C(b *testing.B) { benchmarkDrawing(1000, 1, b) }
func BenchmarkDrawing10L2C(b *testing.B)   { benchmarkDrawing(10, 2, b) }
func BenchmarkDrawing11L2C(b *testing.B)   { benchmarkDrawing(11, 2, b) }
func BenchmarkDrawing8L3C(b *testing.B)    { benchmarkDrawing(8, 3, b) }
func BenchmarkDrawing12L2C(b *testing.B)   { benchmarkDrawing(12, 2, b) }
