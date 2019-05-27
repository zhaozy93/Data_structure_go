package btree

import (
	"fmt"
	"log"
	"testing"
)

func compare(i, j interface{}) int {
	if i.(int) > j.(int) {
		return 1
	}
	if i.(int) < j.(int) {
		return -1
	}
	return 0
}

// func TestBTree(t *testing.T) {
// 	return
// 	log.Println("BTree test start")
// 	// 测试AddValue at Node level
// 	bt := NewBTree(4, compare)
// 	root := NewNode(bt)
// 	bt.SetRoot(root)
// 	root.AddValue(1)
// 	root.AddValue(2)
// 	root.AddValue(3)
// 	fmt.Println("Print-------------------")
// 	bt.root.Print(0, 0)
// 	root.AddValue(4)
// 	bt.root.child[1].AddValue(5)
// 	bt.root.child[1].AddValue(6)
// 	fmt.Println("Print-------------------")
// 	bt.root.Print(0, 0)
// 	bt.root.child[1].AddValue(100)
// 	fmt.Println("Print-------------------")
// 	bt.root.Print(0, 0)
// 	bt.root.child[1].AddValue(14)
// 	bt.root.child[1].AddValue(20)
// 	fmt.Println("Print-------------------")
// 	bt.root.Print(0, 0)
// 	bt.root.child[1].AddValue(13)
// 	fmt.Println("Print-------------------")
// 	bt.root.Print(0, 0)
// 	bt.root.child[1].AddValue(15)
// 	fmt.Println("Print-------------------")
// 	bt.root.Print(0, 0)
// 	bt.root.AddValue(7)
// 	bt.root.AddValue(8)
// 	fmt.Println("Print-------------------")
// 	bt.root.Print(0, 0)
// 	bt.root.AddValue(9)
// 	fmt.Println("Print-------------------")
// 	bt.root.Print(0, 0)
// 	// root.AddValue(0)
// 	fmt.Println("Start Search-------------------")
// 	fmt.Println(bt.Search(1))
// 	fmt.Println(bt.Search(2))
// 	fmt.Println(bt.Search(7))
// 	fmt.Println(bt.Search(8))
// 	fmt.Println(bt.Search(13))
// 	fmt.Println(bt.Search(10))
// 	fmt.Println(bt.Search(20))
// 	fmt.Println(bt.Search(30))
// 	fmt.Println(bt.Search(100))
// 	fmt.Println("Start Insert Test---------")
// 	bt = NewBTree(4, compare)
// 	root = NewNode(bt)
// 	bt.SetRoot(root)
// 	bt.Insert(1)
// 	bt.Insert(6)
// 	bt.Insert(10)
// 	fmt.Println("Print-------------------")
// 	bt.root.Print(0, 0)
// 	bt.Insert(4)
// 	fmt.Println("Print-------------------")
// 	bt.root.Print(0, 0)
// 	bt.Insert(9)
// 	fmt.Println("Print-------------------")
// 	bt.root.Print(0, 0)
// 	bt.Insert(10)
// 	bt.Insert(10)
// 	bt.Insert(4)
// 	bt.Insert(1)
// 	bt.Insert(6)
// 	bt.Insert(9)
// 	fmt.Println("Print-------------------")
// 	bt.root.Print(0, 0)
// 	bt.Insert(7)
// 	fmt.Println("Print-------------------")
// 	bt.root.Print(0, 0)
// }

func TestBTreeUp(t *testing.T) {
	log.Println("BTree test start")
	// 测试AddValue at Node level
	bt := NewBTree(4, compare)
	log.Println(bt.Insert(1))
	log.Println(bt.Insert(2))
	log.Println(bt.Insert(3))
	log.Println(bt.Insert(4))
	log.Println(bt.Insert(5))
	log.Println(bt.Insert(6))
	log.Println(bt.Insert(7))
	log.Println(bt.Insert(8))
	// log.Println(bt.Insert(4))
	// log.Println(bt.Insert(4))
	// log.Println(bt.root)
	bt.root.Print(0, 0)
	fmt.Println(bt.root.bitValue)
	fmt.Println(bt.root.child[0].bitValue)
	fmt.Println(bt.root.child[1].bitValue)
	// fmt.Println(bt.root.child[2].bitValue)
}
