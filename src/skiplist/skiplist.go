package skiplist

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"sync"
)

var ERREXISTED = errors.New("elem has existed")

type Skiplist struct {
	maxLevel    int
	head        *SkiplistNode
	len         int64
	cap         int64
	step        int
	locker      *sync.Mutex
	compareFunc func(interface{}, interface{}) int
}

type SkiplistNode struct {
	Value interface{}
	Level int
	Right []*SkiplistNode // 0层-Level层 每层的右侧node
}

func NewSkiplist(cap int64, lev int, step int, compareFunc func(interface{}, interface{}) int) *Skiplist {
	if cap <= 0 {
		cap = math.MaxInt64
	}
	head := newNode(lev, nil)
	fmt.Println(head)
	return &Skiplist{
		maxLevel:    lev,
		head:        head,
		len:         0,
		cap:         cap,
		step:        step,
		locker:      new(sync.Mutex),
		compareFunc: compareFunc,
	}
}

func (sk *Skiplist) Insert(elem interface{}) error {
	sk.locker.Lock()
	defer sk.locker.Unlock()
	level := sk.randomLevel()
	node := newNode(level, elem)
	if sk.search(elem) != nil {
		return ERREXISTED
	}
	for i := 0; i <= level; i++ {
		head := sk.head
		for {
			if head.Right[i] == nil {
				head.Right[i] = node
				break
			}
			_com := sk.compareFunc(elem, head.Right[i].Value)
			if _com > 0 {
				head = head.Right[i]
				continue
			} else {
				node.Right[i] = head.Right[i]
				head.Right[i] = node
				break
			}
		}
	}
	return nil
}

func (sk *Skiplist) Search(elem interface{}) bool {
	sk.locker.Lock()
	defer sk.locker.Unlock()
	node := sk.search(elem)
	if node == nil {
		return false
	}
	return true
}

func (sk *Skiplist) search(elem interface{}) *SkiplistNode {
	tempNode := sk.head
	for i := sk.maxLevel; i >= 0; i-- {
		for {
			if tempNode.Right[i] == nil {
				break
			}
			_com := sk.compareFunc(elem, tempNode.Right[i].Value)
			if _com == 0 {
				return tempNode.Right[i]
			}
			if _com == 1 {
				tempNode = tempNode.Right[i]
				continue
			}
			break
		}
	}
	return nil
}

func (sk *Skiplist) Delete(elem interface{}) bool {
	sk.locker.Lock()
	defer sk.locker.Unlock()
	node := sk.search(elem)
	if node == nil {
		return false
	}
	sk.delete(node)
	return true
}

func (sk *Skiplist) delete(node *SkiplistNode) {
	tempPreNode := sk.head
	for i := node.Level; i >= 0; i-- {
		for {
			if node != tempPreNode.Right[i] {
				tempPreNode = tempPreNode.Right[i]
				continue
			}
			tempPreNode.Right[i] = node.Right[i]
			break
		}
	}
}

func (sk *Skiplist) randomLevel() int {
	level := 0
	for rand.Int()%sk.step != 0 {
		level++
	}
	if level > sk.maxLevel {
		level = sk.maxLevel
	}
	return level
}

func newNode(lev int, elem interface{}) *SkiplistNode {
	right := make([]*SkiplistNode, lev+1)
	for i := 0; i <= lev; i++ {
		right[i] = nil
	}
	return &SkiplistNode{
		Value: elem,
		Level: lev,
		Right: right,
	}
}

// 测试辅助打印
func (sk *Skiplist) Print() {
	for i := 0; i <= sk.maxLevel; i++ {
		fmt.Print("line:", i, "----")
		head := sk.head
		_ = head
		for {
			if head.Right[i] == nil {
				break
			}
			fmt.Print(head.Right[i].Value, "-")
			head = head.Right[i]
		}
		fmt.Print("\n")
	}
}
