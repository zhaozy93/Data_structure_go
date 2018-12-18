package code_30

import (
	"math"
)

// 面试题30：包含min函数的栈
// 题目：定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的min
// 函数。在该栈中，调用min、push及pop的时间复杂度都是O(1)。

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

type MinStack struct {
	dataStack  *Stack
	minStack   *Stack
	currentMin int
}

func (m *MinStack) Push(val int) {
	if m.dataStack == nil {
		m.dataStack = &Stack{}
		m.minStack = &Stack{}
		m.dataStack.Push(val)
		m.minStack.Push(val)
		m.currentMin = val
		return
	}
	m.dataStack.Push(val)
	if val < m.currentMin {
		m.currentMin = val
	}
	m.minStack.Push(m.currentMin)
}
func (m *MinStack) Pop() (val int, ok bool) {
	if m.dataStack == nil {
		return
	}
	val, ok = m.dataStack.Pop()
	if ok {
		m.minStack.Pop()
		min, ok := m.minStack.Pop()
		if !ok {
			m.dataStack = nil
			m.minStack = nil
		} else {
			m.currentMin = min
			m.minStack.Push(min)
		}
	}
	return
}

func (m *MinStack) Min() (val int, ok bool) {
	if m.dataStack == nil {
		return
	}
	return m.currentMin, true
}

func NewMinStcak() *MinStack {
	stack1 := &Stack{}
	stack2 := &Stack{}
	return &MinStack{stack1, stack2, math.MaxInt32}
}
