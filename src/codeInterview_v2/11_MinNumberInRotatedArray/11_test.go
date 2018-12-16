package code_11

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	a := []int{3, 4, 5, 1, 2}
	log.Println(FindMinNum(a))
	b := []int{1, 1, 0, 1, 1, 1, 1}
	log.Println(FindMinNum(b))
}
