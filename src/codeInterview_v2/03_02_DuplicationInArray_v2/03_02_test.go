package code_03_02

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	log.Println(FindDuplicationNoEdit([]int{1, 1, 2, 3, 4, 5}))
	log.Println(FindDuplicationNoEdit([]int{1, 2, 2, 3, 4, 5}))
	log.Println(FindDuplicationNoEdit([]int{1, 2, 4, 3, 4, 4}))
}
