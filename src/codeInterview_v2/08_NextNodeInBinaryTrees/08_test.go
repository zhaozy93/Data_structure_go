package code_08

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	h := &Node{Val: "h"}
	i := &Node{Val: "i"}
	e := &Node{Val: "e", LeftC: h, RightC: i}
	d := &Node{Val: "d"}
	b := &Node{Val: "b", LeftC: d, RightC: e}

	f := &Node{Val: "f"}
	g := &Node{Val: "g"}
	c := &Node{Val: "c", LeftC: f, RightC: g}
	a := &Node{Val: "a", LeftC: b, RightC: c}

	h.Parent = e
	i.Parent = e
	e.Parent = b
	d.Parent = b
	b.Parent = a
	c.Parent = a
	g.Parent = c
	f.Parent = c

	log.Println(GetNextInOrderNodeInBinaryTrees(i).Val) // a
	log.Println(GetNextInOrderNodeInBinaryTrees(a).Val) // c
	log.Println(GetNextInOrderNodeInBinaryTrees(c).Val) // g
	log.Println(GetNextInOrderNodeInBinaryTrees(g))     // nil
}
