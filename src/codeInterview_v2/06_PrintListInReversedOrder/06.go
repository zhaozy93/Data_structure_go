package code_06

import "fmt"

// 面试题6：从尾到头打印链表
// 题目：输入一个链表的头结点，从尾到头反过来打印出每个结点的值。

type Link struct {
	val  interface{}
	next *Link
}

func ReversePrintLink(l *Link) {
	if l.next != nil {
		ReversePrintLink(l.next)
	}
	fmt.Println("Val is ", l.val)
}
