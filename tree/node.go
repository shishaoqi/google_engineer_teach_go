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

	/*nodes := []tree.Node {
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)*/

	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	//root.print()//print(root)
	fmt.Println()

	root.Traverse()

	nodeCount := 0
	root.TraverseFunc(func(abbc *tree.Node){
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)
}
