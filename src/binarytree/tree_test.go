package tree

import (
	"log"
	"testing"
)

type TestVal int

func (t TestVal) Less(c Value) int {
	return int(t) - int(c.(TestVal))
}

func TestLink(t *testing.T) {
	log.Println("Binary Tree test start")
	r := NewBinaryNode(TestVal(1))
	r.SetLeftChild(NewBinaryNode(TestVal(2)))
	r.SetRightChild(NewBinaryNode(TestVal(3)))
	r.GetLeftChild().SetLeftChild(NewBinaryNode(TestVal(4)))
	r.GetLeftChild().SetRightChild(NewBinaryNode(TestVal(5)))
	r.GetRightChild().SetLeftChild(NewBinaryNode(TestVal(6)))
	r.GetRightChild().SetRightChild(NewBinaryNode(TestVal(7)))
	tree := NewBinaryTree()
	tree.root = r
	inorder := tree.InTraversal()
	log.Println("InTraversal", inorder)
	preorder := tree.PreTraversal()
	log.Println("PreTraversal", preorder)
	postorder := tree.PostTraversal()
	log.Println("PostTraversal", postorder)

	tree = NewBinaryTree()
	log.Println("Insert 1", tree.Insert(TestVal(4)))
	log.Println("Insert 1", tree.Insert(TestVal(1)))
	tree.Insert(TestVal(2))
	tree.Insert(TestVal(3))
	tree.Insert(TestVal(9))
	tree.Insert(TestVal(7))
	tree.Insert(TestVal(8))
	inorder = tree.InTraversal()
	log.Println("InTraversal", inorder)
}
