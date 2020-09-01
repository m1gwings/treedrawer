# treedrawer
**treedrawer** is a Go module that will help you drawing trees to console like the one below.
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
```sh
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
*tree.Tree implements the Stringer interface, just use fmt to draw trees to console
```go
fmt.Println(t)
```
### Implementing NodeValue interface
## Examples
## Known issues
- Emojis are larger than normal characters
```go
fmt.Println(tree.NewTree(tree.NodeString("emojis are buggy 🤪")))
```
```
╭──────────────────╮ 
│emojis are buggy 🤪│ 
╰──────────────────╯ 

```
