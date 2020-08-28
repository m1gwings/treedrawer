package main

import (
	"flag"
	"fmt"

	"github.com/SmashXric/treedrawer/tree"
)

var lFlag = flag.Int("l", 4, "Max number of layers in the random tree")

func main() {
	flag.Parse()
	t := tree.Rand(*lFlag)
	fmt.Println(t)
}
