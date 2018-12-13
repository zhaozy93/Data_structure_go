package code_04

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	martix := [][]int{
		[]int{1, 2, 8, 9},
		[]int{2, 4, 9, 12},
		[]int{4, 7, 10, 13},
		[]int{7, 8, 11, 15},
	}
	log.Println(FindInPartiallySortedMatrix(martix, 9))
	log.Println(FindInPartiallySortedMatrix(martix, 7))
	log.Println(FindInPartiallySortedMatrix(martix, 6))
}
