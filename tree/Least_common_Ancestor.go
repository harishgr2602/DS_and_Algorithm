package main

import (
	"fmt"
)

func (t *Tree) LCA(first int, second int) (int, bool) {
	ans := LCAUtil(t.root, first, second)
	if ans != nil {
		return ans.value, true
	}
	fmt.Println("NotFoundError")
	return 0, false
}

func LCAUtil(curr *Node, first int, second int) *Node {
	var left, right *Node

	if curr == nil {
		return nil
	}

	if curr.value == first || curr.value == second {
		return curr
	}

	left = LCAUtil(curr.left, first, second)
	right = LCAUtil(curr.right, first, second)

	if left != nil && right != nil {
		return curr
	} else if left != nil {
		return left
	} else {
		return right
	}
}

func main() {
	LCAUtil(curr.left)
	LCAUtil(curr.right)
	fmt.Println()
}
