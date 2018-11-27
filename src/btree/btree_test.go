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

func TestLink(t *testing.T) {
	log.Println("BTree test start")
	// 测试InsertValue at Node level
	bt := NewBTree(4, compare)
	root := NewNode(bt)
	bt.SetRoot(root)
	root.InsertValue(1)
	root.InsertValue(2)
	root.InsertValue(3)
	fmt.Println("Print-------------------")
	bt.root.Print(0, 0)
	root.InsertValue(4)
	bt.root.child[1].InsertValue(5)
	bt.root.child[1].InsertValue(6)
	fmt.Println("Print-------------------")
	bt.root.Print(0, 0)
	bt.root.child[1].InsertValue(100)
	fmt.Println("Print-------------------")
	bt.root.Print(0, 0)
	bt.root.child[1].InsertValue(14)
	bt.root.child[1].InsertValue(20)
	fmt.Println("Print-------------------")
	bt.root.Print(0, 0)
	bt.root.child[1].InsertValue(13)
	fmt.Println("Print-------------------")
	bt.root.Print(0, 0)
	bt.root.child[1].InsertValue(15)
	fmt.Println("Print-------------------")
	bt.root.Print(0, 0)
	bt.root.InsertValue(7)
	bt.root.InsertValue(8)
	fmt.Println("Print-------------------")
	bt.root.Print(0, 0)
	bt.root.InsertValue(9)
	fmt.Println("Print-------------------")
	bt.root.Print(0, 0)
	// root.InsertValue(0)
}
