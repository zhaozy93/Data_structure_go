package sort

import (
	"fmt"
	"testing"
)

func TestLink(t *testing.T) {
	a := []int{3, 1, 100, 5, 3, 7, 44, 56, 23, 76, 45, 98, 99, 77, 77, 43, 0}
	BubbleSort(a)
	fmt.Println(a)

	a = []int{3, 1, 100, 5, 3, 7, 44, 56, 23, 76, 45, 98, 99, 77, 77, 43, 0}
	SelectionSort(a)
	fmt.Println(a)

	a = []int{3, 1, 100, 5, 3, 7, 44, 56, 23, 76, 45, 98, 99, 77, 77, 43, 0}
	InsertSort(a)
	fmt.Println(a)

	a = []int{3, 1, 100, 5, 3, 7, 44, 56, 23, 76, 45, 98, 99, 77, 77, 43, 0}
	ShellSort(a)
	fmt.Println(a)

	a = []int{3, 1, 100, 5, 3, 7, 44, 56, 23, 76, 45, 98, 99, 77, 77, 43, 0}
	fmt.Println(MergeSort(a))

	a = []int{3, 1, 100, 5, 3, 7, 44, 56, 23, 76, 45, 98, 99, 77, 77, 43, 0}
	QuickSort(a)
	fmt.Println(a)

	a = []int{3, 1, 100, 5, 3, 7, 44, 56, 23, 76, 45, 98, 99, 77, 77, 43, 0}
	HeapSort(a)
	fmt.Println(a)
}
