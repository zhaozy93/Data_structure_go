package find_sub_string

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	log.Println(ViolenceSearchSubstring("CGTCCGTCA", "CGTCA"))
	log.Println(KMPSubstring("CGTCCGTCA", "CGTCA"))
}
