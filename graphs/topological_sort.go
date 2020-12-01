package main

import (
	"fmt"
)

type Graph struct {
	count      int
	VertexList [](*EdgeList)
}

func (g *Graph) BFS() {
	count := g.count
	visited := make([]int, count)
	for i := 0; i < count; i++ {
		visited[i] = 0
	}
	for i := 0; i < count; i++ {
		if visited[i] == 0 {
			g.BFSQueue(i, visited)
		}
	}
}

func (g *Graph) BFSQueue(index int, visited []int) {
	var curr int
	que := new(queue.Queue)
	visited[index] = 1
	que.Enqueue(index)

	for que.Len() != 0 {
		curr = que.Dequeue().(int)
		head := g, VertexList[curr].head
		for head != nil {
			if visited[head.destination] == 0 {
				visited[head.destination] = 1
				que.Enqueue(head.destination)
			}
			head = head.next
		}
	}
}

func (g *Graph) TopologicalSort() {
	fmt.Print("Toplogical order of a given graph is:")
	var count = g.count
	stk := new(stack.Stack)
	visited := make([]int, count)
	for i := 0; i < count; i++ {
		visited[i] = 0
	}
	for i := 0; i < count; i++ {
		if visited[i] == 0 {
			visited[i] = 1
			g.TopologicalSortDFS(i, visited, stk)
		}
	}
	for stk.Len() != 0 {
		fmt.Print(stk.Pop().(int), "")
	}
	fmt.Println()
}

func (g *Graph) TopologicalSortDFS(index int, visited []int, stk *stack.Stack) {
	head := g.VertexList[index].head
	for head != nil {
		if visited[head.destination] == 0 {
			visited[head.destination] = 1
			g.TopologicalSortDFS(head.destination, visited, stk)
		}
		head = head.next
	}
	stk.Push(index)
}
