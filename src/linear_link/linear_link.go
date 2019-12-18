package linear_link

import "fmt"

type node struct{
	val interface{}
	next *node
}

type Array struct{
	d *node 
	len int
}

func NewArray() *Array {
	return &Array{
		d: &node{val:nil,next:nil},
		len: 0,
	}
}

func (a *Array) Len()int{
	return a.len
}

func (a *Array) Push(val interface{})*Array{
	return a.PushWithIndex(val, a.len)
}

func (a *Array) PushWithIndex(val interface{}, index int)*Array{
	if index > a.len{
		index = a.len
	}
	n := a.d
	for i:=0; i< index; i++{
		n = n.next
	}
	next := n.next
	n.next = &node{
		val: val,
		next: next,
	}
	a.len += 1
	return a
}

func (a *Array) DeleteWithIndex(index int) (val interface{}){
	if index >= a.len{
		return
	}
	node := a.d
	for i:=0; i< index; i++{
		node = node.next
	}
	val = node.next.val
	node.next = node.next.next
	a.len -= 1
	return 
}

func (a *Array) Print(){
	n := a.d.next
	for i:=0; i < a.len; i++{
		if n != nil {
			fmt.Print(n.val, ",")
		}
		n = n.next
	}
	fmt.Println(a.len)
}