package skiplist

import (
	"log"
	"testing"
)

func Compare(l interface{}, r interface{}) int {
	if l.(int) > r.(int) {
		return 1
	}
	if l.(int) < r.(int) {
		return -1
	}
	return 0
}

func TestLink(t *testing.T) {
	log.Println("test skiplist start")

	sk := NewSkiplist(10, 4, Compare)
	sk.Insert(4)
	sk.Insert(3)
	sk.Insert(2)
	sk.Insert(1)

	for i := 100; i > 0; i-- {
		sk.Insert(i)
	}
	sk.Print()

	log.Println(sk.Search(101))
	log.Println(sk.Search(35))
	log.Println(sk.Search(60))
	sk.Delete(85)
	// sk.Print()

}
