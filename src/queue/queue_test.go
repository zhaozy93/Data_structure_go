package queue

import (
	"log"
	"testing"
)

func TestQueue(t *testing.T) {
	log.Println("queue test start")
	s := NewQueue(2)
	log.Println("EMPTY Queue:", s.Empty())
	res, err := s.Shift()
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
	log.Println("Len is: ", s.GetLen())
	a, _ := s.Shift()
	log.Println(a)
	a, _ = s.Shift()
	log.Println(a)
	a, err = s.Shift()
	log.Println(err.Error())
}
