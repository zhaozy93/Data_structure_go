package code_11

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	a := []int{1, 6, 2, 9, 4, 7, 3, 0}
	QuickSort(a)
	log.Println(a)
}
