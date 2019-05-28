package main

import "fmt"
import "algorithms/tree"

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)


	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println()

	root.Traverse()

	nodeCount := 0
	root.TraverseFunc(func(abbc *tree.Node){
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("Max node value:", maxNode)

}
