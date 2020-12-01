package main

type Graph struct {
	count      int
	VertexList [](*EdgeList)
}

func (g *Graph) DFSStack() {
	count := g.count
	visited := make([]int, count)
	var curr int
	stk := new(stack.Stack)
	for i := 0; i < count; i++ {
		visited[i] = 0
	}
	visited[1] = 1
	stk.Push(1)

	for stk.Len() != 0 {
		curr = stk.Pop().(int)
		head := g.VertexList[curr].head
		for head != nil {
			if visited[head.destination] == 0 {
				visited[head.destination] = 1
				stk.Push(head.destination)
			}
			head = head.next
		}
	}
}
