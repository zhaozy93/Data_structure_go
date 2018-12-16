package code_23

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
	f.Next = b

	log.Println(FindLoopEntryNode(a))

}
