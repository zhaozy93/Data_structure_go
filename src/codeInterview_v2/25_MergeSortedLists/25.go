package code_25

// 面试题25：合并两个排序的链表
// 题目：输入两个递增排序的链表，合并这两个链表并使新链表中的结点仍然是按
// 照递增排序的。例如输入图3.11中的链表1和链表2，则合并之后的升序链表如链
// 表3所示。

type Node struct {
	Val  int
	Next *Node
}

func MergeSortedLists(r1, r2 *Node) (root *Node) {
	if r1 == nil {
		return r2
	}
	if r2 == nil {
		return r1
	}

	if r1.Val < r2.Val {
		root = r1
		root.Next = MergeSortedLists(r1.Next, r2)
	} else {
		root = r2
		root.Next = MergeSortedLists(r1, r2.Next)
	}
	return root

}
