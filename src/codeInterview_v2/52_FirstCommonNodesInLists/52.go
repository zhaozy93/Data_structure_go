package code_52

type Node struct {
	Val  int
	Next *Node
}

type Stack struct {
	vals []*Node
}

func (s *Stack) Push(val *Node) {
	if s.vals == nil {
		s.vals = make([]*Node, 0)
	}
	s.vals = append(s.vals, val)
}

func (s *Stack) Pop() (*Node, bool) {
	if s.vals == nil || len(s.vals) == 0 {
		return 0, false
	}
	val := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	if len(s.vals) == 0 {
		s.vals = nil
	}
	return val, true
}

func FirstCommonNodesInLists(root1 *Node, root2 *Node) {
	s1 := &Stack{}
	s2 := &Stack{}

	for root1 != nil {
		s1.Push(root1)
	}
}
