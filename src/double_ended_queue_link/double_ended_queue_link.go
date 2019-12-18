package double_ended_queue_link

import (
	"linear_link"
)

type DEQueue struct{
	d *linear_link.Array
}

func NewDEQueue() *DEQueue{
	return &DEQueue{
		d: linear_link.NewArray(),
	}
}

func (s *DEQueue) UnShift(val interface{}) *DEQueue{
	s.d.PushWithIndex(val, 0)
	return s
}

func (s *DEQueue) Push(val interface{}) *DEQueue{
	s.d.Push(val)
	return s
}

func (s *DEQueue) Shift() (val interface{}, ok bool){
	if s.d.Len() == 0 {
		return 
	} 
	val = s.d.DeleteWithIndex(0)
	ok = true 
	return 
}

func (s *DEQueue) Pop() (val interface{}, ok bool){
	if s.d.Len() == 0 {
		return 
	}
	val = s.d.DeleteWithIndex(s.d.Len()-1)
	ok = true 
	return 
}

func (s *DEQueue) Len() int{
	s.d.Print()
	return s.d.Len()
}


