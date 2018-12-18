package code_36

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	e := &Node{Val: 8}
	d := &Node{Val: 4}
	b := &Node{Val: 6, LeftC: d, RightC: e}

	c := &Node{Val: 12}
	a := &Node{Val: 10, LeftC: b, RightC: c}

	ConvertBinarySearchTree(a)
	log.Println(d)
	log.Println(d.RightC)
	log.Println(d.RightC.RightC)
	log.Println(d.RightC.RightC.RightC)
	log.Println(d.RightC.RightC.RightC.RightC)

	log.Println(d.RightC.RightC.RightC.RightC)
	log.Println(d.RightC.RightC.RightC.RightC.LeftC)
	log.Println(d.RightC.RightC.RightC.RightC.LeftC.LeftC)
	log.Println(d.RightC.RightC.RightC.RightC.LeftC.LeftC.LeftC)
	log.Println(d.RightC.RightC.RightC.RightC.LeftC.LeftC.LeftC.LeftC)
}
