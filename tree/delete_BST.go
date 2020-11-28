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

func (t *Tree) DeleteNode(value int) {
	t.root = DeleteNode(t.root, value)
}

func DeleteNode(node *Node, value int) *Node {
	var temp *Node = nil
	if node != nil {
		if node.value == value {
			if node.left == nil && node.right == nil {
				return nil
			}
			if node.left == nil {
				temp = node.right
				return temp
			}
			if node.right == nil {
				temp = node.left
				return temp
			}
			maxNode := FindMax(node.left)
			maxValue := maxNode.value
			node.value = maxValue
			node.left = DeleteNode(node.left, maxValue)
		} else {
			if node.value > value {
				node.left = DeleteNode(node.left, value)
			} else {
				node.right = DeleteNode(node.right, value)
			}
		}
	}
	return node
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := CreateBinaryTree(arr)
}
