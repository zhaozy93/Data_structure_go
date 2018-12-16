package code_20

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	log.Println(isNumeric("123"))
	log.Println(isNumeric("123.1"))
	log.Println(isNumeric("+100"))
	log.Println(isNumeric("5e2"))
	log.Println(isNumeric("-123"))
	log.Println(isNumeric("3.1416"))
	log.Println(isNumeric("-1E16"))
	log.Println(isNumeric("-1E-16"))
	log.Println(isNumeric("-1E"))
	log.Println(isNumeric("-1.."))
}
