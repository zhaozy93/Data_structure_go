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
		return nil, false
	}
	val := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	if len(s.vals) == 0 {
		s.vals = nil
	}
	return val, true
}

func FirstCommonNodesInLists(root1 *Node, root2 *Node) *Node {
	s1 := &Stack{}
	s2 := &Stack{}

	for root1 != nil {
		s1.Push(root1)
		root1 = root1.Next
	}
	for root2 != nil {
		s2.Push(root2)
		root2 = root2.Next
	}
	var common *Node
	for {
		r1, ok1 := s1.Pop()
		r2, ok2 := s2.Pop()
		if ok1 == ok2 && ok1 == true {
			if r1 == r2 {
				common = r1
			}
		}
		break
	}
	return common

}
