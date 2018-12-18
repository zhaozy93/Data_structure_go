package code_35

// 面试题35：复杂链表的复制
// 题目：请实现函数 ComplexListNode* Clone(ComplexListNode* pHead)，复
// 制一个复杂链表。在复杂链表中，每个结点除了有一个m_pNext指针指向下一个
// 结点外，还有一个m_pSibling 指向链表中的任意结点或者nullptr。

type ComplexListNode struct {
	Val     int
	Next    *ComplexListNode
	Sibling *ComplexListNode
}

func ComplexListNodeClone(root *ComplexListNode) (mirror *ComplexListNode) {
	if root == nil {
		return
	}
	node := root
	// 复制链表 形成 a --> a_ --> b --> b_
	for {
		if node == nil {
			break
		}
		nodeMirror := cloneNode(node)
		nodeMirror.Next = node.Next
		node.Next = nodeMirror
		node = nodeMirror.Next
	}
	// 修正a_的sibling属性
	node = root
	index := 0
	for {
		if node == nil {
			break
		}
		if index%2 == 0 {
			index++
			node = node.Next
			continue
		}
		if node.Sibling != nil {
			node.Sibling = node.Sibling.Next
		}
		node = node.Next
	}
	// 拆分链表
	node = root
	mirror = root.Next
	for {
		mirror := node.Next
		next := node.Next.Next
		if next == nil {
			node.Next = nil
			break
		}
		mirrorNext := node.Next.Next.Next
		node.Next = next
		mirror.Next = mirrorNext
		node = next
	}
	return
}

func cloneNode(node *ComplexListNode) *ComplexListNode {
	return &ComplexListNode{
		Val:     node.Val,
		Next:    node.Next,
		Sibling: node.Sibling,
	}
}
