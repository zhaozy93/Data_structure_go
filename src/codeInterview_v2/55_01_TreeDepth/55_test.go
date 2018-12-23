package code_55

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	h := &Node{Val: 7}
	i := &Node{Val: 6, Right: h}
	e := &Node{Val: 5, Left: i}
	d := &Node{Val: 4, Left: e}
	b := &Node{Val: 3, Left: d}

	a := &Node{Val: 3, Left: b}

	log.Println(TreeDepth(a))
}
