package lcs

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	// log.Println(RecursiveMethod(4))
	// log.Println(MemorandumMethod(5))
	log.Println(DynamicLcs("GTACCGTCA", "CATCGA"))
}
