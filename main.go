package main

import (
	"fmt"

	"github.com/x1n13y84issmd42/dm/graphs"
)

type MyNode struct {
	graphs.Node
	Name string
}

func main() {
	root := MyNode{
		Name: "A",
	}

	root.Add(&MyNode{
		Name: "B",
	}).Add(&MyNode{
		Name: "C",
	})

	root.Add(&MyNode{
		Name: "D",
	})

	for node := range graphs.DFS(&root) {
		fmt.Printf("%s\n", node.(*MyNode).Name)
	}
}
