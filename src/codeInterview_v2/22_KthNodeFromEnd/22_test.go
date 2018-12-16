package code_22

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	d := &Node{Val: "d"}
	c := &Node{Val: "c", Next: d}
	b := &Node{Val: "b", Next: c}
	a := &Node{Val: "a", Next: b}

	log.Println(FindKthToTail(a, 1))
	log.Println(FindKthToTail(a, 6))

	log.Println(FindKthToTail(a, 2))

}
