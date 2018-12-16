package code_09

// 面试题9：用两个栈实现队列
// 题目：用两个栈实现一个队列。队列的声明如下，请实现它的两个函数appendTail
// 和deleteHead，分别完成在队列尾部插入结点和在队列头部删除结点的功能。

type Stack struct {
	vals []int
}

func (s *Stack) Add(val int) {
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
	return val, true
}

type Queeu struct {
	s1 *Stack
	s2 *Stack
}

func NewQueeu() *Queeu {
	return &Queeu{s1: &Stack{}, s2: &Stack{}}
}

func (q *Queeu) AppendTail(val int) {
	q.s1.Add(val)
}

func (q *Queeu) DeleteHead() (int, bool) {
	if val, ok := q.s2.Pop(); ok {
		return val, ok
	}
	for {
		if val, ok := q.s1.Pop(); ok {
			q.s2.Add(val)
			continue
		}
		break
	}
	return q.s2.Pop()
}
