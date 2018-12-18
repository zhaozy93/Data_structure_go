package code_30

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	m := NewMinStcak()
	m.Push(1)
	m.Push(2)
	m.Push(3)
	log.Println(m.Min())
	log.Println(m.Pop())
	log.Println(m.Min())
	log.Println(m.Pop())
	log.Println(m.Min())
	log.Println(m.Pop())
	// false
	log.Println(m.Min())
	log.Println(m.Pop())
	m.Push(3)
	m.Push(2)
	log.Println(m.Min())
	m.Push(1)
	log.Println(m.Min())
}
