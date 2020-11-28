package main

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

func (t *Tree) Add(value int) {
	t.root = addUtil(t.root, value)
}

func addUtil(n *Node, value int) *Node {
	if n == nil {
		n = new(Node)
		n.value = value
		return n
	}

	if value < n.value {
		n.left = addUtil(n.left, value)
	} else {
		n.right = addUtil(n.right, value)
	}
	return n
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := CreateBinaryTree(arr)
}
