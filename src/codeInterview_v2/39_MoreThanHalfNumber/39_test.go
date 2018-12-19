package code_39

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	a := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	log.Println(MoreThanHalfNumber(a))
}
