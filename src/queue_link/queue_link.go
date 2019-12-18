package queue_link

import (
	"linear_link"
)

type Queue struct{
	d *linear_link.Array
}

func NewQueue() *Queue{
	return &Queue{
		d: linear_link.NewArray(),
	}
}

func (s *Queue) Push(val interface{}) *Queue{
	s.d.Push(val)
	return s
}

func (s *Queue) Shift() (val interface{}, ok bool){
	if s.d.Len() == 0 {
		return 
	} 
	val = s.d.DeleteWithIndex(0)
	ok = true 
	return 
}

func (s *Queue) Len() int{
	s.d.Print()
	return s.d.Len()
}


