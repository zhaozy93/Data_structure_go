package code_27

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

	mirror := MirrorOfBinaryTree(a)
	log.Println(mirror)
	log.Println(mirror.LeftC)
	log.Println(mirror.LeftC.LeftC)
	log.Println(mirror.LeftC.RightC)
	log.Println(mirror.RightC)
	log.Println(mirror.RightC.LeftC)
	log.Println(mirror.RightC.RightC)
	log.Println(mirror.RightC.LeftC.LeftC)
	log.Println(mirror.RightC.LeftC.RightC)
}
