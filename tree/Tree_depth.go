package amin

import (
	"fmt"
)

func (t *Tree) TreeDepth() int {
	return TreeDepth(t.root)
}

func TreeDepth(root *Node) int {
	if root == nil {
		return 0
	}
	lDepth := TreeDepth(root.left)
	rDepth := TreeDepth(root.right)

	if lDepth > rDepth {
		return lDepth + 1
	}
	return rDepth + 1
}

func main() {
	fmt.Println(temp)
}
