package code_55

// 面试题55（一）：二叉树的深度
// 题目：输入一棵二叉树的根结点，求该树的深度。从根结点到叶结点依次经过的
// 结点（含根、叶结点）形成树的一条路径，最长路径的长度为树的深度。

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func TreeDepth(root *Node) int {
	if root == nil {
		return 0
	}
	leftDepth := TreeDepth(root.Left)
	rightDepth := TreeDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}
