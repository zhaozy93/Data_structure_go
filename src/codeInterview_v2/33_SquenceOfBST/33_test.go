package code_33

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	test := []int{5, 7, 6, 9, 11, 10, 8}
	log.Println(VerifySquenceOfBST(test))

	test = []int{7, 4, 6, 5}
	log.Println(VerifySquenceOfBST(test))
}
