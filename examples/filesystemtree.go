package main

import (
	"fmt"
	"log"
	"os"

	"github.com/m1gwings/treedrawer/tree"
)

func main() {
	t := tree.NewTree(tree.NodeString(""))

	var f func(string, string, *tree.Tree)
	f = func(previousPath, name string, t *tree.Tree) {
		t.SetVal(tree.NodeString(name))
		file, err := os.Open(previousPath + name)
		if err != nil {
			log.Fatal(err)
		}
		fileNames, err := file.Readdirnames(-1)
		if err != nil {
			return
		}
		for _, newName := range fileNames {
			// Ignoring hidden files
			if newName[0] == '.' {
				return
			}
			f(previousPath+name+"/", newName, t.AddChild(tree.NodeString("")))
		}
	}
	// Starting from FirstChild because the DocumentRoot has an empty Val
	f("../../", "treedrawer", t)

	fmt.Println(t)
}
