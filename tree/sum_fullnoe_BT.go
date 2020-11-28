package main

func LevelOrderBinaryTree(arr []int) *Tree {
	tree := new(Tree)
	tree.root = LevelOrderBinaryTree(arr, 0, len(arr))
	return tree
}

func LevelOrderBinaryTree(arr []int, start int, size int) *Node {
	curr := &Node{arr[start], nil, nil}

	left := 2*start + 1
	right := 2*start + 2

	if left < size {
		curr.left = LevelOrderBinaryTree(arr, left, size)
	}

	if right < size {
		curr.right = LevelOrderBinaryTree(arr, right, size)
	}
	return curr
}

func (t *Tree) SumAllBT() int {
	return sumAllBT(t.root)
}

func sumAllBT(curr *Node) int {
	var sum, leftsum, rightsum, int
	if curr == nil {
		return 0
	} 

	rightsum = sumAllBT(curr.right)
	leftsum = sumAllBT(curr.left)
	sum = rigthsum + leftsum + curr.value
	return sum
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t2 := LevelOrderBinaryTree(arr)
	t2.PreOrder()
}
