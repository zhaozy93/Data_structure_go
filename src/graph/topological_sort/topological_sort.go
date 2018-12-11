package topological_sort

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

func TopologicalSort(graph *Graph) [][]Vector {
	resSort := make([][]Vector, 0)
	stack := make([]Vector, 0)
	topologicalSort(graph, stack, &resSort)
	return resSort
}

func topologicalSort(grpah *Graph, stack []Vector, resSort *[][]Vector) {
	if len(grpah.Vectors) == 0 {
		*resSort = append(*resSort, stack)
		return
	}
	inDegree, _ := getInDegreeVector(grpah)
	for v, in := range inDegree {
		if in == 0 {
			_graph := removeVectorFromGraph(grpah, v)
			_stack := append(stack, v)
			topologicalSort(_graph, _stack, resSort)
		}
	}
}
func getInDegreeVector(graph *Graph) (inDegree, outDegree map[Vector]int) {
	inDegree = make(map[Vector]int)
	outDegree = make(map[Vector]int)
	for _, v := range graph.Vectors {
		inDegree[v] = 0
		outDegree[v] = 0
	}
	for _, l := range graph.Lines {
		inDegree[l.EVector]++
		outDegree[l.FVector]++
	}
	return

}
func removeVectorFromGraph(graph *Graph, vector Vector) *Graph {
	_graph := &Graph{}
	_graph.Vectors = make([]Vector, len(graph.Vectors)-1)
	hasFound := false
	for i := 0; i < len(graph.Vectors); i++ {
		if graph.Vectors[i] == vector {
			hasFound = true
			continue
		}
		if hasFound {
			_graph.Vectors[i-1] = graph.Vectors[i]
		} else {
			_graph.Vectors[i] = graph.Vectors[i]
		}
	}

	lines := make([]Line, 0)
	for _, l := range graph.Lines {
		if l.EVector == vector || l.FVector == vector {
			continue
		}
		lines = append(lines, l)
	}
	_graph.Lines = lines
	return _graph
}
