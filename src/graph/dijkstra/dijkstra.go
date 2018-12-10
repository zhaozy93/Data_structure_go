package dijkstra

import (
	"math"
)

type Vector string
type Line struct {
	FVector Vector
	EVector Vector
	Value   int
}
type Graph struct {
	Vectors []Vector
	Lines   []Line
}

func Dijkstra(graph *Graph, vector Vector) (map[Vector]int, map[Vector]Vector) {
	shortest := make(map[Vector]int)
	for _, v := range graph.Vectors {
		shortest[v] = math.MaxInt32
	}
	shortest[vector] = 0
	predVector := make(map[Vector]Vector)
	dijkstra(graph, vector, shortest, predVector)
	return shortest, predVector
}
func dijkstra(graph *Graph, vector Vector, shortest map[Vector]int, predVector map[Vector]Vector) {
	for _, l := range graph.Lines {
		if l.FVector == vector {
			if shortest[vector]+l.Value < shortest[l.EVector] {
				shortest[l.EVector] = shortest[vector] + l.Value
				predVector[l.EVector] = vector
				dijkstra(graph, l.EVector, shortest, predVector)
			}
		}
	}
}
