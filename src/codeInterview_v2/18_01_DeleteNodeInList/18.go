package code_18

// 面试题18（一）：在O(1)时间删除链表结点
// 题目：给定单向链表的头指针和一个结点指针，定义一个函数在O(1)时间删除该
// 结点。

type Node struct {
	Val  string
	Next *Node
}

func DeleteNodeInList(root *Node, p *Node) *Node {
	if root == p {
		if p.Next != nil {
			root = p.Next
		} else {
			root = nil
		}
	} else if p.Next == nil {
		n := root
		for {
			if n.Next == p {
				n.Next = nil
				break
			}
			n = n.Next
		}
	} else {
		p.Val = p.Next.Val
		p.Next = p.Next.Next
	}

	return root
}
