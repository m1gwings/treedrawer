# treedrawer
treedrawer is a simple go module that will help you to visualize binary trees in the command line.
# Motivation
I started to build this module by trying to solve an exercise on "The Go Programming Language". The task was to implement the String() function of a binary tree.
# Code structure
This module has two sub-packages:
* tree
* drawer

**tree** provides a simple representation of a binary tree:
```
// Tree describes the node of a tree with atmost two children.
type Tree struct {
	val                 int64
	left, right, parent *Tree
}
```

**drawer** is a very simple "ascii-canvas" on which you can draw bytes representing ascii characters or another whole canvas specifying the coordinates of the upper-left corner.

---
If you run the following in the root directory:
```
go build .
```

you will get a binary that prints random trees to give you an idea.
```
./treedrawer

        59         
         |         
        32         
        / \        
       /   \       
      /     \      
     /       \     
   7         77    
    |        / \   
   19      24   91 
```

With the **-l** flag you can specify the maximum number of layers of the random tree
```
./treedrawer -l 3

   70    
    |    
   28    
   / \   
 27   64 
```
# API
## Building the tree
### func NewTree(val int64) *Tree
NewTree is the default constructor for Tree.
```
t := tree.NewTree(2)
```
### func (t *Tree) AddLeft(val int64)
AddLeft adds a left child to the current node which will held val.
```
t.AddLeft(5)
```
### func (t *Tree) AddRight(val int64)
AddRight adds a right child to the current node which will held val.
```
t.AddRight(3)
```
## Printing the tree
Tree type satisfies the Stringer interface, you can easily use fmt package to get results in the console.
```
fmt.Println(t)

   2   
  / \  
 5   3 
```
## Retreiving values from the tree
### func (t *Tree) Val() int64
Val returns the value held by the current node of the tree.
## Navigating the tree
### func (t *Tree) Left() (ok bool)
Left moves the current node to its left child.
Returns false if there is no left child, otherwise it returns true.
### func (t *Tree) Right() (ok bool)
Right moves the current node to its right child.
Returns false if there is no right child, otherwise it returns true.
### func (t *Tree) Parent() (ok bool)
Parent moves the current node to its parent.
Returns false if this node is the root of the whole tree, otherwise it returns true.
## Extra
### func Rand(n int) *Tree
Rand returns the root of a random three with at most n layers.