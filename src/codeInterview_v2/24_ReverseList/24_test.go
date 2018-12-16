package code_24

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	f := &Node{Val: "f"}
	e := &Node{Val: "e", Next: f}
	d := &Node{Val: "d", Next: e}
	c := &Node{Val: "c", Next: d}
	b := &Node{Val: "b", Next: c}
	a := &Node{Val: "a", Next: b}

	r := ReverseList(a)
	log.Println(r)
	log.Println(r.Next)
	log.Println(r.Next.Next)
	log.Println(r.Next.Next.Next)
	log.Println(r.Next.Next.Next.Next)
	log.Println(r.Next.Next.Next.Next.Next)

	r = ReverseList(a)
	log.Println(r)
}
