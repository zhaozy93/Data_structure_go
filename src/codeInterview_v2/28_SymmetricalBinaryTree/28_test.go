package code_28

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

	log.Println(CheckIsSymmetryTree(a))

	d1 := &Node{Val: "a"}
	e1 := &Node{Val: "a"}
	b1 := &Node{Val: "a", LeftC: e1}
	c1 := &Node{Val: "a", RightC: d1}
	a1 := &Node{Val: "a", LeftC: b1, RightC: c1}
	log.Println(CheckIsSymmetryTree(a1))

	d2 := &Node{Val: "a"}
	e2 := &Node{Val: "a"}
	b2 := &Node{Val: "a", LeftC: e2}
	c2 := &Node{Val: "a", LeftC: d2}
	a2 := &Node{Val: "a", LeftC: b2, RightC: c2}
	log.Println(CheckIsSymmetryTree(a2))

}
