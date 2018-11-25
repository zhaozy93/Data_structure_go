package link

import (
	"log"
	"testing"
)

func TestLink(t *testing.T) {
	log.Println("Link test start")
	s := NewLink(3)
	log.Println("EMPTY Link:", s.Empty())
	res, err := s.Shift()
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(res)
	}
	res, err = s.Pop()
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(res)
	}
	s.Push(1)
	s.Push(2)
	_, err = s.Unshift(3)
	if err != nil {
		log.Println(err.Error())
	}
	_, err = s.Push(4)
	if err != nil {
		log.Println(err.Error())
	}
	_, err = s.Push(5)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Len is: ", s.GetLen())

	a, _ := s.Shift()
	log.Println(a)
	a, _ = s.Shift()
	log.Println(a)
	a, err = s.Shift()
	log.Println(a)
}
