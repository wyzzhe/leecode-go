package main

import (
	"fmt"

	"github.com/wyzzhe/leecode-go/tree"
)

func main() {
	runTree()
}

func runTree() {
	flag := "Ddd"
	switch flag {
	case "tree":
		tree := tree.NewTreeNode()
		fmt.Printf("%+v\n", tree)
	case "Ddd":
		result := tree.Ddd()
		fmt.Println(result)
	}
}
