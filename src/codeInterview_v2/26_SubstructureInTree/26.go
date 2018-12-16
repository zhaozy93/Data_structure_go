package code_26

// 面试题26：树的子结构
// 题目：输入两棵二叉树A和B，判断B是不是A的子结构。

type Node struct {
	Val    string
	LeftC  *Node
	RightC *Node
}

func SubstructureInTree(root, subRoot *Node) bool {
	if root == nil || subRoot == nil {
		return false
	}

	if root.Val == subRoot.Val {
		if hasSameContent(root, subRoot) {
			return true
		}
	}

	if SubstructureInTree(root.LeftC, subRoot) {
		return true
	}
	if SubstructureInTree(root.RightC, subRoot) {
		return true
	}
	return false
}

func hasSameContent(root, subRoot *Node) bool {
	if subRoot == nil {
		return true
	}
	if root == nil {
		return false
	}
	if root.Val != subRoot.Val {
		return false
	}

	if !hasSameContent(root.LeftC, subRoot.LeftC) {
		return false
	}
	if !hasSameContent(root.RightC, subRoot.RightC) {
		return false
	}
	return true
}
