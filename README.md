# treedrawer
**treedrawer** is a Go module for drawing trees on the terminal like the one below.
```
                                    ╭─╮                                    
                                    │9│                                    
                                    ╰┬╯                                    
                   ╭─────────────────┴─────────────┬─────────────┬───┬───╮ 
        ╭──────────┴─────────╮                    ╭┴╮           ╭┴╮ ╭┴╮ ╭┴╮
        │I can handle strings│                    │1│           │2│ │3│ │4│
        ╰──────────┬─────────╯                    ╰┬╯           ╰─╯ ╰─╯ ╰─╯
                   │                      ╭─────┬──┴─────╮                 
  ╭────────────────┴────────────────╮   ╭─┴─╮ ╭─┴╮  ╭────┴───╮             
  │with as many children as you want│   │124│ │13│  │a string│             
  ╰────────────────┬────────────────╯   ╰───╯ ╰──╯  ╰────────╯             
                   │                                                       
   ╭───────────────┴───────────────╮                                       
   │with as many layers as you want│                                       
   ╰───────────────┬───────────────╯                                       
                   │                                                       
 ╭─────────────────┴─────────────────╮                                     
 │actually I can handle everything...│                                     
 ╰─────────────────┬─────────────────╯                                     
                   │                                                       
╭──────────────────┴──────────────────╮                                    
│...that satisfies NodeValue interface│                                    
╰─────────────────────────────────────╯                                    
```
## Import
```go
import "github.com/m1gwings/treedrawer/tree"
```
## Quick start
```sh
# Assume the following code is in example.go file
$ cat example.go
```
```go
package main

import (
	"fmt"

	"github.com/m1gwings/treedrawer/tree"
)

func main() {
	// Creating a tree with 5 as the value of the root node
	t := tree.NewTree(tree.NodeInt64(5))

	// Adding children
	t.AddChild(tree.NodeString("adding a string"))
	t.AddChild(tree.NodeInt64(4))
	t.AddChild(tree.NodeInt64(3))

	// Drawing the tree
	fmt.Println(t)
}
```
```sh
$ go run example.go
```
```
           ╭─╮           
           │5│           
           ╰┬╯           
        ╭───┴──────┬───╮ 
╭───────┴───────╮ ╭┴╮ ╭┴╮
│adding a string│ │4│ │3│
╰───────────────╯ ╰─╯ ╰─╯

```
## Usage
### Building the tree
Creating the tree with 1 as the value of the root node
```go
t := tree.NewTree(tree.NodeInt64(1))
```
Adding the first child to t with value 2
```go
t.AddChild(tree.NodeInt64(2))
```
Adding more children
```go
t.AddChild(tree.NodeInt64(3))
t.AddChild(tree.NodeInt64(4))
t.AddChild(tree.NodeInt64(5))
```
We've just built the tree below
```
      ╭─╮      
      │1│      
      ╰┬╯      
 ╭───┬─┴─┬───╮ 
╭┴╮ ╭┴╮ ╭┴╮ ╭┴╮
│2│ │3│ │4│ │5│
╰─╯ ╰─╯ ╰─╯ ╰─╯

```
### Navigating the tree
Navigating to first child of t (we're still working with the tree above)
```go
// This method returns an error if the i-th child does not exist
// in this case i = 0
tFirstChild, err := t.Child(0)
```
Adding children to first child
```go
tFirstChild.AddChild(tree.NodeInt64(6))
tFirstChild.AddChild(tree.NodeInt64(7))
tFirstChild.AddChild(tree.NodeInt64(8))
```
Going back to parent
```go
// ok would be equal to false if tFirstChild were the root of the tree
tFirstChildParent, ok := tFirstChild.Parent()

_ := tFirstChildParent == t // true, we have gone back to the root of the tree
```
Navigating to third child of t
```go
tThirdChild, err := t.Child(2)
```
Adding a string child to third child
```go
tThirdChild.AddChild(tree.NodeString("I'm a string"))
```
Getting a pointer to the root of the tree
```go
tRoot := tThirdChild.Root()

_ := tRoot == t // true
```
Now the tree looks like this
```
                ╭─╮                
                │1│                
                ╰┬╯                
     ╭───────┬───┴─────┬─────────╮ 
    ╭┴╮     ╭┴╮       ╭┴╮       ╭┴╮
    │2│     │3│       │4│       │5│
    ╰┬╯     ╰─╯       ╰┬╯       ╰─╯
 ╭───┼───╮             │           
╭┴╮ ╭┴╮ ╭┴╮     ╭──────┴─────╮     
│6│ │7│ │8│     │I'm a string│     
╰─╯ ╰─╯ ╰─╯     ╰────────────╯     

```
### Getting and setting values from the tree
Getting the value of a node
```go
v := t.Val()
```
Setting the value of a node
```go
t.SetVal(tree.NodeInt64(3))
```
### Drawing the tree
*tree.Tree implements the Stringer interface, just use package fmt to draw trees to terminal
```go
fmt.Println(t)
```
### Implementing NodeValue interface
The tree can handle every type that satisfies the **NodeValue** interface
```go
// NodeValue is the interface that wraps the Draw method.
//
// The Draw method allows to convert data into its unicode canvas representation.
// With the Draw method you can control how your data is going to appear on the tree.
type NodeValue interface {
	Draw() *drawer.Drawer
}
```
The wrappers for built-in are defined inside the package treedrawer/tree like tree.NodeInt64 or tree.NodeString used above, so you don't need to worry about them.  
Continue reading this section if you want to draw custom types instead.
- Importing treedrawer/drawer

First of all we need access to the drawer.Drawer type. Just import the following
```go
import "github.com/m1gwings/treedrawer/drawer"
```
drawer.Drawer under the hood is just a 2D slice of runes on which you can draw a rune specifying its coordinates or another entire drawer.Drawer specifying the coordinates of its upper-left corner.
- Defining a custom type
```go
type NodeAsterisk struct {
	Width, Height int
}
```
NodeAsterisk represents a rectangle of width NodeAsterisk.Width and height NodeAsterisk.Height.  

- Implementing NodeAsterisk.Draw() in order to satisfy NodeValue interface
```go
func (nA NodeAsterisk) Draw() *drawer.Drawer {
	d, err := drawer.NewDrawer(nA.Width, nA.Height)
	if err != nil {
		log.Fatal(err)
	}
	for x := 0; x < nA.Width; x++ {
		for y := 0; y < nA.Height; y++ {
			err = d.DrawRune('*', x, y)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return d
}
```
The method allocates a new drawer with width nA.Width and height nA.Height, then loops over each cell and fills it with an '*'.  
You can implement this method to represent your data as you want.

- Adding instances of NodeAsterisk to a tree
```go
t := tree.NewTree(NodeAsterisk{3, 4})
t.AddChild(NodeAsterisk{1, 2})
t.AddChild(NodeAsterisk{3, 3})
```
- Drawing the tree
```go
fmt.Println(t)
```
```
  ╭───╮  
  │***│  
  │***│  
  │***│  
  │***│  
  ╰─┬─╯  
 ╭──┴─╮  
╭┴╮ ╭─┴─╮
│*│ │***│
│*│ │***│
╰─╯ │***│
    ╰───╯

```
## Examples
You can find these examples inside the **./examples** folder
### HTML tree
```sh
# Assume the following code is in htmltree.go file
$ cat htmltree.go
```
```go
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
```
```sh
$ go run htmltree.go
```
```
             ╭────╮              
             │html│              
             ╰──┬─╯              
   ╭────────────┴───╮            
╭──┴─╮           ╭──┴─╮          
│head│           │body│          
╰────╯           ╰──┬─╯          
            ╭───────┴────╮       
           ╭┴╮         ╭─┴╮      
           │p│         │ul│      
           ╰┬╯         ╰─┬╯      
            │       ╭────┴──╮    
        ╭───┴──╮  ╭─┴╮    ╭─┴╮   
        │Links:│  │li│    │li│   
        ╰──────╯  ╰─┬╯    ╰─┬╯   
                    │       │    
                   ╭┴╮     ╭┴╮   
                   │a│     │a│   
                   ╰┬╯     ╰┬╯   
                    │       │    
                  ╭─┴─╮ ╭───┴──╮ 
                  │Foo│ │BarBaz│ 
                  ╰───╯ ╰──────╯ 

```
### File system tree
```sh
# Assume the following code is in filesystemtree.go file
$ cat filesystemtree.go
```
```go
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
```
```sh
$ go run filesystemtree.go
```
```
                                                                                ╭──────────╮                                                                                 
                                                                                │treedrawer│                                                                                 
                                                                                ╰─────┬────╯                                                                                 
     ╭──────────────────────┬─────────────────────┬───────────────────┬───────────────┴───┬───────────────────────────────────────────╮                                      
╭────┴────╮            ╭────┴───╮             ╭───┴───╮           ╭───┴──╮            ╭───┴──╮                                     ╭──┴─╮                                    
│README.md│            │examples│             │LICENSE│           │drawer│            │go.sum│                                     │tree│                                    
╰─────────╯            ╰────┬───╯             ╰───────╯           ╰───┬──╯            ╰──────╯                                     ╰──┬─╯                                    
                  ╭─────────┴──────╮                            ╭─────┴────────╮                         ╭─────────────────┬──────────┴─┬────────────┬───────────────╮       
            ╭─────┴─────╮ ╭────────┴────────╮           ╭───────┴──────╮  ╭────┴────╮           ╭────────┴───────╮  ╭──────┴─────╮  ╭───┴───╮ ╭──────┴──────╮ ╭──────┴─────╮ 
            │htmltree.go│ │filesystemtree.go│           │drawer_test.go│  │drawer.go│           │examples_test.go│  │stringify.go│  │tree.go│ │nodevalues.go│ │tree_test.go│ 
            ╰───────────╯ ╰─────────────────╯           ╰──────────────╯  ╰─────────╯           ╰────────────────╯  ╰────────────╯  ╰───────╯ ╰─────────────╯ ╰────────────╯ 

```
## Benchmarks
You can find the code used for the benchmark inside **./tree/stringify_test.go**  
In order to profile the module we first create trees with **l** layers and **c** children for each node, except leaf nodes. Each node has a tree.NodeString("*") as value.  
For example the tree below has 3 layers and 2 children for each node.
```
      ╭─╮      
      │*│      
      ╰┬╯      
   ╭───┴───╮   
  ╭┴╮     ╭┴╮  
  │*│     │*│  
  ╰┬╯     ╰┬╯  
 ╭─┴─╮   ╭─┴─╮ 
╭┴╮ ╭┴╮ ╭┴╮ ╭┴╮
│*│ │*│ │*│ │*│
╰─╯ ╰─╯ ╰─╯ ╰─╯

```
In our benchmark function we print to **/dev/null** a tree with the specified **l** and **c** parameter.  
Name|Iterations|Time|Children per Node|Layers|Total of Nodes|Memory|Allocations
-|-|-|-|-|-|-|-|
BenchmarkDrawing3L3C-12|10000|100063 ns/op|3.00 children|3.00 layers|13.0 nodes|135576 B/op|722 allocs/op
BenchmarkDrawing100L1C-12|382|3096956 ns/op|1.00 children|100 layers|100 nodes|3727628 B/op|23297 allocs/op
BenchmarkDrawing6L3C-12|9|119789317 ns/op|3.00 children|6.00 layers|364 nodes|366549320 B/op|33606 allocs/op
BenchmarkDrawing1000L1C-12|7|161607620 ns/op|1.00 children|1000 layers|1000 nodes|373048737 B/op|2033004 allocs/op
BenchmarkDrawing10L2C-12|2|733952840 ns/op|2.00 children|10.0 layers|1023 nodes|4375790984 B/op|114909 allocs/op
BenchmarkDrawing11L2C-12|1|3661883138 ns/op|2.00 children|11.0 layers|2047 nodes|20557038432 B/op|249910 allocs/op
BenchmarkDrawing8L3C-12|1|8550947886 ns/op|3.00 children|8.00 layers|3280 nodes|50574173168 B/op|387828 allocs/op
BenchmarkDrawing12L2C-12|1|13559015034 ns/op|2.00 children|12.0 layers|4095 nodes|96166288000 B/op|538283 allocs/op
##### Generated using go version go1.15.1 linux/amd64
## Known issues 🐛
- Emojis are larger than normal characters
```go
fmt.Println(tree.NewTree(tree.NodeString("emojis are buggy 🤪")))
```
```
╭──────────────────╮ 
│emojis are buggy 🤪│ 
╰──────────────────╯ 

```
