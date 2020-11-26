package main

import (
	"fmt"
)

func (t *Tree) NthInOrder(index int) {
	var counter int
	nthInOrder(t.root, index, &counter)
}

func nthInOrder(node *Node, index int, counter *int) {
	if node != nil {
		nthInOrder(node.left, index, counter)
		*counter++
		if *counter == index {
			fmt.Println(node.value)
		}
		nthInOrder(node.right, index, counter)
	}
}

func main() {
	nthInOrder(node.left)
	nthInOrder(node.right)
}
