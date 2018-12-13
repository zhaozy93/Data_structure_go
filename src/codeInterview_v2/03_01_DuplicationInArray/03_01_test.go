package code_03_01

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	log.Println(FindDuplication([]int{0, 1, 2, 3, 4, 5}))
	log.Println(FindDuplication([]int{1, 1, 2, 3, 4, 5}))
	log.Println(FindDuplication([]int{0, 1, 4, 3, 4, 5}))
	// log.Println(ViolenceSearchSubstring("CGTCCGTCA", "CGTCA"))
	// log.Println(KMPSubstring("CGTCCGTCA", "CGTCA"))
}
