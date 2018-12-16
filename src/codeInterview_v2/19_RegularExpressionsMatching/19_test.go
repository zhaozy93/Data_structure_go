package code_19

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	log.Println(RegularExpressions("aaa", "..."))
	log.Println(RegularExpressions("aaa", "a.a"))
	log.Println(RegularExpressions("aaa", "a.."))
	log.Println(RegularExpressions("aaa", "..."))
	log.Println(RegularExpressions("aaa", "...."))

	log.Println(RegularExpressions("aaa", "a*"))
	log.Println(RegularExpressions("aaa", "a*b*"))
	log.Println(RegularExpressions("aaa", ".b*a*aa*"))
	log.Println(RegularExpressions("aaa", ".b*a*aa*"))
	log.Println(RegularExpressions("aaa", "ab*a"))
}
