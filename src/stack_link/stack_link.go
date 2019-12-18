package stack_link

import (
	"linear_link"
)

type Stack struct{
	d *linear_link.Array
}

func NewStack() *Stack{
	return &Stack{
		d: linear_link.NewArray(),
	}
}

func (s *Stack) Push(val interface{}) *Stack{
	s.d.Push(val)
	return s
}

func (s *Stack) Pop() (val interface{}, ok bool){
	if s.d.Len() == 0 {
		return 
	} 
	val = s.d.DeleteWithIndex(s.d.Len()-1)
	ok = true 
	return 
}

func (s *Stack) Len() int{
	s.d.Print()
	return s.d.Len()
}


