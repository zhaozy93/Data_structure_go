package code_35

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	a := &ComplexListNode{Val: 1}
	b := &ComplexListNode{Val: 2}
	c := &ComplexListNode{Val: 3}
	d := &ComplexListNode{Val: 4}
	e := &ComplexListNode{Val: 5}
	f := &ComplexListNode{Val: 6}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f

	a.Sibling = c
	b.Sibling = d
	c.Sibling = e
	d.Sibling = f

	mirror := ComplexListNodeClone(a)
	log.Println(mirror == a)
	log.Println(mirror, mirror.Sibling)
	mirror = mirror.Next
	log.Println(mirror, mirror.Sibling)
	mirror = mirror.Next
	log.Println(mirror, mirror.Sibling)
	mirror = mirror.Next
	log.Println(mirror, mirror.Sibling)
	mirror = mirror.Next
	log.Println(mirror, mirror.Sibling)
	mirror = mirror.Next
	log.Println(mirror, mirror.Sibling)

}
