package code_56

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	a := []int{2, 4, 3, 6, 3, 2, 5, 5}
	log.Println(FindNumsAppearOnce(a))
}
