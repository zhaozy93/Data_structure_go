package code_22

// 面试题22：链表中倒数第k个结点
// 题目：输入一个链表，输出该链表中倒数第k个结点。为了符合大多数人的习惯，
// 本题从1开始计数，即链表的尾结点是倒数第1个结点。例如一个链表有6个结点，
// 从头结点开始它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个结点是
// 值为4的结点。

type Node struct {
	Val  string
	Next *Node
}

func FindKthToTail(root *Node, index int) (n *Node) {
	if index <= 0 {
		return nil
	}
	if root == nil {
		return nil
	}
	n = root
	pre := root
	for i := 0; i < index-1; i++ {
		if pre.Next == nil {
			return nil
		}
		pre = pre.Next
	}

	for {
		if pre.Next == nil {
			break
		}
		pre = pre.Next
		n = n.Next
	}
	return
}