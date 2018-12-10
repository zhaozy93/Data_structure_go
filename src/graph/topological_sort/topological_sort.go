package topological_sort

type Vector string
type Line struct {
	Vectors []*Vector
	Value   int
}
type Graph struct {
	Vectors []*Vector
	Lines   []*Line
}

func TopologicalSort(graph *Graph) [][]*Point {
	stack := make([]*Point, 0)
	topologicalSort(graph, stack)
}

func topologicalSort(grpah *Graph, stack []*Point) {
	inDegree := make(map[Vector]int)
	for _, l := range grpah.Lines {
		f := l.Vectors[0]
		e := l.Vectors[1]
		if _, ok := inDegree[f]; !ok {
			inDegree[f] = 0
		}
		inDegree[e]++
	}
	for v,in:=range inDegree{
		if in == 0{
			vs:=make([]*Vector, len(inDegree)-1)
			for i:=0;i<
			topologicalSort()
		}
	}
}
