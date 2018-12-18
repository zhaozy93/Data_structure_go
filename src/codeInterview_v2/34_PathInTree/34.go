package code_34

import (
	"fmt"
)

type Node struct {
	Val    int
	LeftC  *Node
	RightC *Node
}

type Stack struct {
	vals []int
}

func (s *Stack) Push(val int) {
	if s.vals == nil {
		s.vals = make([]int, 0)
	}
	s.vals = append(s.vals, val)
}

func (s *Stack) Pop() (int, bool) {
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

func PathInTree(root *Node, sum int) {
	path := &Stack{}
	PathInTreeCore(root, sum, 0, path)
}

func PathInTreeCore(node *Node, sum, current int, path *Stack) {
	if current > sum {
		return
	}
	current += node.Val
	// 叶子节点了
	if node.LeftC == nil && node.RightC == nil {
		if current == sum {
			fmt.Println(path.vals, node.Val)
			return
		}
	}
	path.Push(node.Val)
	if node.LeftC != nil {
		PathInTreeCore(node.LeftC, sum, current, path)
	}
	if node.RightC != nil {
		PathInTreeCore(node.RightC, sum, current, path)
	}
	path.Pop()
}
