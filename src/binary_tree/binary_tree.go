package binary_tree

import (
	"double_ended_queue_link"
)

type Node struct{
	Left *Node
	Right *Node
	Value interface{}
}

func NewNode(l,r *Node, v interface{})*Node{
	return &Node{l,r,v}
}

type Tree struct {
	Root *Node
}

func NewTree(Root *Node)*Tree{
	return &Tree{Root}
}

// Don't use loop
func (t *Tree) PreTraversal() (vals []*Node){
	if t.Root == nil{
		return
	}
	vals = make([]*Node, 0)
	q := double_ended_queue_link.NewDEQueue()
	q.Push(t.Root)
	var val *Node
	for {
		if tmp, ok := q.Shift(); !ok{
			break
		}else{
			val = tmp.(*Node)
		}
		vals = append(vals, val)
		if val.Right != nil{
			q.UnShift(val.Right)
		}
		if val.Left != nil{
			q.UnShift(val.Left)
		}
	}
	return 
}

func (t *Tree) PostTraversal() (vals []*Node){
	if t.Root == nil{
		return
	}
	vals = make([]*Node, 0)
	q := double_ended_queue_link.NewDEQueue()
	q.Push(t.Root)
	var val *Node
	for {
		if tmp, ok := q.Shift(); !ok{
			break
		}else{
			val = tmp.(*Node)
		}
		if val.Right == nil && val.Left == nil{
			vals = append(vals, val)
			continue
		}
		if len(vals) != 0 && val.Right == vals[len(vals)-1]{
			vals = append(vals, val)
			continue
		}
		q.UnShift(val)
		if val.Right != nil{
			q.UnShift(val.Right)
		}
		if val.Left != nil{
			q.UnShift(val.Left)
		}
	}
	return
}

func (t *Tree) InTraversal() (vals []*Node){
	if t.Root == nil{
		return
	}
	vals = make([]*Node, 0)
	q := double_ended_queue_link.NewDEQueue()
	node := t.Root
	q.UnShift(node)
	for {
		if node.Left != nil {
			q.UnShift(node.Left)
			node = node.Left
		}else{
			break
		}
	}
	var val *Node
	for {
		if tmp, ok := q.Shift(); !ok{
			break
		}else{
			val = tmp.(*Node)
		}
		vals = append(vals, val)
		if val.Right != nil{
			node := val.Right
			q.UnShift(node)
			for {
				if node.Left != nil {
					q.UnShift(node.Left)
					node = node.Left
				}else{
					break
				}
			}
		}
	}
	return 
}