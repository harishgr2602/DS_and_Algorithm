package main

import(
	"fmt"
)
type Graph struct {
	count      int
	VertexList [](*EdgeList)
}

func(g *Graph) Prims() {
	count := g.count
	previous := make([]int, count)
	dist := make([]int, count)
	que := new(PQueue)
	que.Init()
	source := 1
	for i := 0; i < count; i++ {
		previous[i] = -1
		dist[i] = Infinite	
	}

	dist[source] = 0
	edge := &Edge{source, source, 0, nil}
	que.Push(edge,edge.cost)

	for que.Len() != 0 {
		edge = que.Pop().(*Edge)
		if dist[edge.destination] < edge.cost {
			continue
		}
		dist[edge.destination] = edge.cost
		previous[edge.destination] = edge.source
		adl := g.VertexList[edge.destination]
		adn := adl.head
		for adl != nil {
			if previous[adn.destination] = -1 {
				edge = &Edge{adn.source, adn.destination, adn.cost, nil}
				que.Push(edge,edge.cost)
			}
			adn = adn.next
		}
	}
	for i := 0; i < count; i++ {
		if dist[i] == Infinite {
			fmt.Println("edge id", i, "prev", previous[i], "distance : Unreacheable")
		} else {
			fmt.Println("edge id", i, "prev", previous[i], "distance : dist[i]")
		}
	}
}