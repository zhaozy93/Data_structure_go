package code_23

// 面试题23：链表中环的入口结点
// 题目：一个链表中包含环，如何找出环的入口结点？例如，在图3.8的链表中，
// 环的入口结点是结点3。

type Node struct {
	Val  string
	Next *Node
}

func FindLoopEntryNode(root *Node) *Node {
	p1 := root
	p2 := root
	for {
		temp := p2.Next
		if temp == nil {
			return nil
		}
		if temp == p1 {
			p2 = temp
			break
		}
		temp = temp.Next
		if temp == nil {
			return nil
		}
		if temp == p1 {
			p2 = temp
			break
		}
		p2 = temp
		p1 = p1.Next
	}
	cnt := 0
	for {
		cnt++
		p1 = p1.Next
		if p1 == p2 {
			break
		}
	}

	p1 = root
	p2 = root

	for i := 0; i < cnt; i++ {
		p1 = p1.Next
	}

	for {
		if p1 == p2 {
			break
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}
