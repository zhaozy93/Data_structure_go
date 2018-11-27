package tree

import (
	"log"
	"testing"
)

func TestLink(t *testing.T) {
	log.Println("Binary Tree test start")
	r := NewBinaryNode(1)
	r.SetLeftChild(NewBinaryNode(2))
	r.SetRightChild(NewBinaryNode(3))
	r.GetLeftChild().SetLeftChild(NewBinaryNode(4))
	r.GetLeftChild().SetRightChild(NewBinaryNode(5))
	r.GetRightChild().SetLeftChild(NewBinaryNode(6))
	r.GetRightChild().SetRightChild(NewBinaryNode(7))
	tree := NewBinaryTree(r)
	inorder := tree.InTraversal()
	log.Println("InTraversal", inorder)
	preorder := tree.PreTraversal()
	log.Println("PreTraversal", preorder)
	postorder := tree.PostTraversal()
	log.Println("PostTraversal", postorder)
}
