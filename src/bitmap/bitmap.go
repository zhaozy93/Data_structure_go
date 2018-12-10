package bitmap

import (
	"bytes"
	"fmt"
)

var bitMap = map[int]int{
	0:  1 << 0,
	1:  1 << 1,
	2:  1 << 2,
	3:  1 << 3,
	4:  1 << 4,
	5:  1 << 5,
	6:  1 << 6,
	7:  1 << 7,
	8:  1 << 8,
	9:  1 << 9,
	10: 1 << 10,
	11: 1 << 11,
	12: 1 << 12,
	13: 1 << 13,
	14: 1 << 14,
	15: 1 << 15,
	16: 1 << 16,
	17: 1 << 17,
	18: 1 << 18,
	19: 1 << 19,
	20: 1 << 20,
	21: 1 << 21,
	22: 1 << 22,
	23: 1 << 23,
	24: 1 << 24,
	25: 1 << 25,
	26: 1 << 26,
	27: 1 << 27,
	28: 1 << 28,
	29: 1 << 29,
	30: 1 << 30,
	31: 1 << 31,
}

type BitMap struct {
	data []int
	len  int
}

func NewBitMap(len int) *BitMap {
	return &BitMap{
		len:  len,
		data: make([]int, len/32+1),
	}
}

func (b *BitMap) Add(i int) {
	index, bit, err := b.checkIndex(i)
	if err == false && b.data[index]&bitMap[bit] != bitMap[bit] {
		b.data[index] += bitMap[bit]
	}
}
func (b *BitMap) Has(i int) bool {
	index, bit, err := b.checkIndex(i)
	if err == true {
		return false
	}
	return b.data[index]&bitMap[bit] == bitMap[bit]
}
func (b *BitMap) String() string {
	var buf bytes.Buffer
	for i := len(b.data) - 1; i >= 0; i-- {
		for j := 31; j >= 0; j-- {
			if b.data[i]&bitMap[j] != 0 {
				buf.WriteByte('1')
			} else {
				buf.WriteByte('0')
			}
		}
		buf.WriteByte(' ')
	}
	return buf.String()
}

func (b *BitMap) checkIndex(i int) (index, bit int, err bool) {
	if i < 0 || i > b.len {
		err = true
		return
	}
	index = i / 32
	bit = (i - index*32) % 32 // i % 8
	fmt.Println("index: ", index, "bit: ", bit)
	return
}
