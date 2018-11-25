package stack

import (
	"log"
	"testing"
)

func TestStack(t *testing.T) {
	log.Println("stack test start")
	s := NewStack(2)
	log.Println("EMPTY STACK:", s.Empty())
	res, err := s.Pop()
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(res)
	}
	s.Push(1)
	s.Push(2)
	_, err = s.Push(3)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(s.GetLen())
	a, _ := s.Pop()
	log.Println(a)
	a, _ = s.Pop()
	log.Println(a)
	a, err = s.Pop()
	log.Println(err.Error())
}
