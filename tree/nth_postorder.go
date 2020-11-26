package main

import (
	"fmt"
)

func (t *Tree) NthPostOrder(index int) {
	var counter int
	nthPostOrder(t.root, index, &counter)
}

func nthPostOrder(node *Node, index int, counter *int) {
	if node != nil {
		nthPostOrder(node.left, index, counter)
		nthPostOrder(node.right, index, counter)
		(*counter)++
		if *counter == index {
			fmt.Println(node.value)
		}
	}
}

func main() {
	nthPostOrder(node.left)
	nthPostOrder(node.right)
	fmt.Println()
}
