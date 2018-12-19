package code_45

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	a := []string{"4", "5", "1", "6", "2", "7", "3", "8"}
	log.Println(SortArrayForMinNumber(a))

}
