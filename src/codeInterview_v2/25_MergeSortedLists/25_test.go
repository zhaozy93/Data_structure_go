package code_25

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	f := &Node{Val: 9}
	e := &Node{Val: 6, Next: f}
	d := &Node{Val: 4, Next: e}
	c := &Node{Val: 2, Next: d}
	b := &Node{Val: 0, Next: c}

	f2 := &Node{Val: 8}
	e2 := &Node{Val: 7, Next: f2}
	d2 := &Node{Val: 5, Next: e2}
	c2 := &Node{Val: 3, Next: d2}
	b2 := &Node{Val: 0, Next: c2}

	r := MergeSortedLists(b, b2)
	log.Println(r)
	log.Println(r.Next)
	log.Println(r.Next.Next)
	log.Println(r.Next.Next.Next)
	log.Println(r.Next.Next.Next.Next)
	log.Println(r.Next.Next.Next.Next.Next)

}
