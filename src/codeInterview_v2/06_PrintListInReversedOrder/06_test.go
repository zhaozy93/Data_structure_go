package code_06

import (
	"testing"
)

func TestSort(t *testing.T) {
	l4 := &Link{
		val:  4,
		next: nil,
	}
	l3 := &Link{
		val:  3,
		next: l4,
	}
	l2 := &Link{
		val:  2,
		next: l3,
	}
	l1 := &Link{
		val:  1,
		next: l2,
	}
	ReversePrintLink(l1)
}
