package code_52

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	h := &Node{Val: 7}
	i := &Node{Val: 6, Next: h}
	e := &Node{Val: 5, Next: i}
	d := &Node{Val: 4, Next: e}
	b := &Node{Val: 3, Next: d}

	a := &Node{Val: 3, Next: b}

	log.Println(FirstCommonNodesInLists(a, b))
}
