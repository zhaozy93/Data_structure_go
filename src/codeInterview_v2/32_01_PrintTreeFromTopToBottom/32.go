package code_32

import (
	"errors"
	"fmt"
)

// 面试题32（一）：不分行从上往下打印二叉树
// 题目：从上往下打印出二叉树的每个结点，同一层的结点按照从左到右的顺序打印。

type Node struct {
	Val    string
	LeftC  *Node
	RightC *Node
}

type QueeuNode struct {
	Val  interface{}
	Next *QueeuNode
}
type Queeu struct {
	root *QueeuNode
	end  *QueeuNode
	len  int
}

func (q *Queeu) Push(val interface{}) {
	node := &QueeuNode{Val: val}
	if q.len == 0 {
		q.root = node
		q.end = node
		q.len = 1
		return
	}
	q.end.Next = node
	q.len++
	q.end = node
}

func (q *Queeu) Pop() (val interface{}, err error) {
	if q.len == 0 {
		err = errors.New("empty")
		return
	}
	val = q.root.Val
	q.root = q.root.Next
	q.len--
	if q.len == 0 {
		q.end = nil
	}
	return
}

func PrintTreeFromTopToBottom(root *Node) {
	queeu := &Queeu{}
	queeu.Push(root)
	queeu.Push(nil)
	lastIsNil := false
	for {
		val, err := queeu.Pop()
		if err != nil {
			break
		}
		// if err2, ok := val.(error); ok && err2 == nil {
		if val == nil {
			fmt.Print("\n")
			queeu.Push(nil)
			if lastIsNil {
				break
			}
			lastIsNil = true
			continue
		}
		if node, ok := val.(*Node); ok {
			lastIsNil = false
			fmt.Print(node.Val, ", ")
			if node.LeftC != nil {
				queeu.Push(node.LeftC)
			}
			if node.RightC != nil {
				queeu.Push(node.RightC)
			}
			continue
		}
	}
}
