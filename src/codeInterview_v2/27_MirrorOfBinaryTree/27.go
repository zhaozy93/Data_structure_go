package code_27

// 面试题27：二叉树的镜像
// 题目：请完成一个函数，输入一个二叉树，该函数输出它的镜像。

type Node struct {
	Val    string
	LeftC  *Node
	RightC *Node
}

func MirrorOfBinaryTree(root *Node) *Node {
	temp := root.LeftC
	root.LeftC = root.RightC
	root.RightC = temp
	if root.LeftC != nil {
		MirrorOfBinaryTree(root.LeftC)
	}
	if root.RightC != nil {
		MirrorOfBinaryTree(root.RightC)
	}
	return root
}
