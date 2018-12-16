package code_18

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	d := &Node{Val: "d"}
	c := &Node{Val: "c", Next: d}
	b := &Node{Val: "b", Next: c}
	a := &Node{Val: "a", Next: b}

	root := DeleteNodeInList(a, a)
	log.Println(root)
	log.Println("---------------")

	d = &Node{Val: "d"}
	c = &Node{Val: "c", Next: d}
	b = &Node{Val: "b", Next: c}
	a = &Node{Val: "a", Next: b}
	root = DeleteNodeInList(a, b)
	log.Println(root)
	log.Println(root.Next)
	log.Println(root.Next.Next)
	log.Println("---------------")

	d = &Node{Val: "d"}
	c = &Node{Val: "c", Next: d}
	b = &Node{Val: "b", Next: c}
	a = &Node{Val: "a", Next: b}
	root = DeleteNodeInList(a, d)
	log.Println(root)
	log.Println(root.Next)
	log.Println(root.Next.Next)
}
