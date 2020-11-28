package main

import "math"

func CreateBinaryTree(arr []int) *Tree {
	t := new(Tree)
	size := len(arr)
	t.root = createBinaryTreeUtil(arr, 0, size-1)
	return t
}

func createBinaryTreeUtil(arr []int, start int, end int) *Node {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	curr := new(Node)
	curr.value = arr[mid]
	curr.left = createBinaryTreeUtil(arr, start, mid-1)
	curr.right = createBinaryTreeUtil(arr, mid+1, end)
	return curr
}

func (t *Tree) FloorBST(val int) int {
	curr := t.root
	floor := math.MaxInt32

	for curr != nil {
		if curr.value == val {
			floor = curr.value
			break
		} else if curr.value > val {
			curr = curr.left
		} else {
			floor = curr.value
			curr = curr.right
		}
	}
	return floor
}

func (t *Tree) CeilBST(val int) int {
	curr := t.root
	ceil := math.MinInt32

	for curr != nil {
		if curr.value == val {
			ceil = curr.value
			break
		} else if curr.value > val {
			ceil = curr.value
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return ceil
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := CreateBinaryTree(arr)
}
