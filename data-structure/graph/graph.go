package graph

import (
	"fmt"
)

type Graph struct {
	vertices []*Vertice
}

type Vertice struct {
	key       int
	adjacents []*Vertice
}

func (g *Graph) getVertex(vertex int) *Vertice {
	for i, v := range g.vertices {
		if v.key == vertex {
			return g.vertices[i]
		}
	}
	return nil
}

func (g *Graph) AddVertex(vertex int) error {
	// 1. I need to check to if vertex is already exists or not
	// 2. if not then i need to append the vertex

	if contains(g.vertices, vertex) {
		err := fmt.Errorf("vertex %d is already exists", vertex)
		return err
	}

	v := &Vertice{
		key: vertex,
	}

	g.vertices = append(g.vertices, v)

	return nil

}

func (g *Graph) AddEdge(to, from int) error {

	//1. first I need to check vertex from [to] and [from]

	toVertex := g.getVertex(to)
	fromVertex := g.getVertex(from)

	if toVertex == nil || fromVertex == nil {
		return fmt.Errorf("not a valid edge from %d --> %d", to, from)
	}

	if contains(fromVertex.adjacents, toVertex.key) {
		return fmt.Errorf("edge from vertex %d  ---> %d is already exists", fromVertex.key, toVertex.key)
	}

	fromVertex.adjacents = append(fromVertex.adjacents, toVertex)
	return nil

}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("%d :", v.key)
		for _, g := range v.adjacents {
			fmt.Printf("%d", g.key)
		}
		fmt.Println()
	}
}

func contains(v []*Vertice, key int) bool {
	for _, v := range v {
		if v.key == key {
			return true
		}
	}

	return false
}
