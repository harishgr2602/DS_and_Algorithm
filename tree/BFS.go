package main

import (
	"fmt"
)

func (t *Tree) PrintBredthFirst() {
	que := new(queue Queue)
	var temp *Node

	if t.root != nil {
		que.Put(t.root)
	}

	for que.Empty() == false {
		temp2, _ := que.Get(1)
		temp = temp2[0].(*Node)
		fmt.Print(temp.value, "") 
		if temp.left != nil {
			que.Put(temp.left)
		}
		if temp.rigth != nil {
			que.Put(temp.right)
		}
	}
	fmt.Println()
}

func main() {
	fmt.Println(temp)
}