package code_12

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	path := "bfce"
	martix := [][]byte{
		[]byte{'a', 'b', 't', 'g'},
		[]byte{'c', 'f', 'c', 's'},
		[]byte{'j', 'd', 'e', 'h'},
	}
	log.Println(HasPath(path, martix))
	path = "abfcjdectgsc"
	log.Println(HasPath(path, martix))
}
