package code_09

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	q := NewQueeu()
	q.AppendTail(1)
	q.AppendTail(2)
	q.AppendTail(3)
	log.Println(q.DeleteHead())
	log.Println(q.DeleteHead())
	q.AppendTail(4)
	log.Println(q.DeleteHead())
	log.Println(q.DeleteHead())
}
