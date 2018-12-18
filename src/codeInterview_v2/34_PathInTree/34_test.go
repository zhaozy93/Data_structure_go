package code_34

import (
	"testing"
)

func TestSort(t *testing.T) {
	// h := &Node{Val: "h"}
	// i := &Node{Val: "i"}
	// e := &Node{Val: 7, LeftC: h, RightC: i}
	e := &Node{Val: 7}
	d := &Node{Val: 4}
	b := &Node{Val: 5, LeftC: d, RightC: e}

	// f := &Node{Val: "f"}
	// g := &Node{Val: "g"}
	// c := &Node{Val: 12, LeftC: f, RightC: g}
	c := &Node{Val: 12}
	a := &Node{Val: 10, LeftC: b, RightC: c}

	PathInTree(a, 22)
}
