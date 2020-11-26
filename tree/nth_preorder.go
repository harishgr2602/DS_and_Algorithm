package main

import (
	"fmt"
)

func (t *Tree) NthPreOrder(index int) {
	var counter int
	nthPreOrder(t.root, index, &counter)
}

func nthPreOrder(node *Node, index int, counter *int) {
	if node != nil {
		(*counter)++
		if *counter == index {
			fmt.Println(node.value)
		}
		nthPreOrder(node.left, index, counter)
		nthPreOrder(node.right, index, counter)
	}
}

func main() {
	fmt.Println(node.left)
	fmt.Println(node.right)
}
