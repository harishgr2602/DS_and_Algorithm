package main

import (
	"fmt"
)

type Edge struct {
	source      int
	destination int
	cost        int
	next        *Edge
}

type EdgeList struct {
	head *Edge
}

type Graph struct {
	count      int
	VertexList [](*EdgeList)
}

func (g *Graph) Init(cnt int) {
	g.count = cnt
	g.VertexList = make([]*EdgeList, cnt)
	for i := 0; i < cnt; i++ {
		g.VertexList[i] = new(EdgeList)
		g.VertexList[i].head = nil
	}
}

func (g *Graph) AddEdge(source int, destination int, cost int) {
	edge := &Edge{source, destination, cost, g.VertexList[source].head}
	g.VertexList[source].head = edge
}

func (g *Graph) AddEdgeUnweighted(source int, desitnation int, cost int) {
	g.AddEdge(source, destination, 1)
}

func (g *Graph) AddBiEdge(source int, destination int, cost int) {
	g.AddEdge(source, destination, cost)
	g.AddEdge(destination, source, cost)
}

func (g *Graph) AddBiEdgeUnweighted(source int, destination int) {
	g.AddBiEdge(source, destination, 1)
}

func (g *Graph) Print() {
	for i := 0; i < g.count; i++ {
		ad := g.VertexList[i].head
		if ad != nil {
			fmt.Print("Vertex", i, "is connected to :")
			for ad != nil {
				fmt.Print("[", ad.destination, ad.cost, "]")
				ad = ad.next
			}
			fmt.Println()
		}
	}
}
