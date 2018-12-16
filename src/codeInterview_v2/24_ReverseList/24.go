package code_24

// 面试题24：反转链表
// 题目：定义一个函数，输入一个链表的头结点，反转该链表并输出反转后链表的
// 头结点。

type Node struct {
	Val  string
	Next *Node
}

func ReverseList(root *Node) *Node {
	var cur, pre *Node
	if root == nil {
		return nil
	}
	cur = root
	for {
		realNext := cur.Next
		cur.Next = pre
		pre = cur
		cur = realNext
		if cur == nil {
			break
		}
	}
	return pre
}
