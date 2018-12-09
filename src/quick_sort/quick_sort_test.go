package quick_sort

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	in := []int{111, 99, 1, 45, 23, 67, 34, 67, 2, 3, 44, 3}
	log.Println(QuickSort(in))
}
