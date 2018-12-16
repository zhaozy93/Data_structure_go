package code_21

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	ReorderArray(s)

	log.Println(s)
}
