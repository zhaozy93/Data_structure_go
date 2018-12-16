package code_08

// 面试题8：二叉树的下一个结点
// 题目：给定一棵二叉树和其中的一个结点，如何找出中序遍历顺序的下一个结点？
// 树中的结点除了有两个分别指向左右子结点的指针以外，还有一个指向父结点的指针。

type Node struct {
	Val    string
	LeftC  *Node
	RightC *Node
	Parent *Node
}

func GetNextInOrderNodeInBinaryTrees(node *Node) (next *Node) {
	next = node.RightC
	if next != nil {
		for {
			if next.LeftC == nil {
				break
			}
			next = next.LeftC
		}
		return
	}

	next = node.Parent
	if next == nil {
		return
	}
	if next.LeftC == node {
		return
	}

	for {
		if next.Parent == nil {
			next = nil
			break
		}
		if next.Parent.LeftC == next {
			next = next.Parent
			break
		}
		next = next.Parent
	}

	return
}
