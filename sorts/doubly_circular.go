type doublyCircular struct {
	head  *Node
	tail  *Node
	count int
}

type Node struct {
	value int
	next  *Node
	prev  *Node
}

func (list *doublyCircular) Size() int {
	return list.count
}

func (list *doublyCircular) IsEmpty() bool {
	return list.count == 0
}

func (list *doublyCircular) peekHead() (int, bool) {
	if list.IsEmpty() {
		fmt.Println(" Empty list Error")
		return 0, false
	}
	return list.head.value, true
}