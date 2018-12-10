package bitmap

import (
	"log"
	"testing"
)

func TestLink(t *testing.T) {
	log.Println("BitMap test start")
	b := NewBitMap(10000)
	for i := 0; i < 32; i++ {
		b.Add(i)
	}
	for i := 0; i < 32; i++ {
		log.Println(b.Has(i + 1))
	}
	log.Println(b.String())
	// b.Add(31)
	// log.Println(b.Existed(1))
	// log.Println(b.Existed(2))
	// log.Println(b.Existed(31))
	// log.Println(b.Existed())

	// b.Add(0)
	// log.Println(b.Existed(1))
	// log.Println(b.Existed(0))
	// log.Println(b.Existed(2))
	log.Println(b.data)
}
