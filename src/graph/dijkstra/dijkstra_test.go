package dijkstra

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	a := Vector("a")
	b := Vector("b")
	c := Vector("c")
	d := Vector("d")
	e := Vector("e")
	f := Vector("f")
	g := Vector("g")
	line1 := Line{a, b, 1}
	line2 := Line{a, c, 1}
	line3 := Line{b, d, 1}
	line4 := Line{c, d, 1}
	line5 := Line{d, e, 1}
	line6 := Line{d, f, 1}
	line7 := Line{a, g, 1}
	line8 := Line{g, f, 1}
	graph := &Graph{}
	graph.Vectors = []Vector{a, b, c, d, e, f, g}
	graph.Lines = []Line{line1, line2, line3, line4, line5, line6, line7, line8}
	log.Println(Dijkstra(graph, a))
}
