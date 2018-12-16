package code_26

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

	t2 := &Node{Val: "g"}
	t3 := &Node{Val: "i"}
	t1 := &Node{Val: "e", LeftC: t2, RightC: t3}

	log.Println(SubstructureInTree(a, e))
	log.Println(SubstructureInTree(a, t1))
}
