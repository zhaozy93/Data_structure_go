package code_36

// 面试题36：二叉搜索树与双向链表
// 题目：输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。要求
// 不能创建任何新的结点，只能调整树中结点指针的指向。
type Node struct {
	Val    int
	LeftC  *Node // 向前指针
	RightC *Node // 向后指针
}

func ConvertBinarySearchTree(root *Node) {
	var pre **Node
	n := &Node{}
	pre = &n
	*pre = nil
	Core(root, pre)
}

func Core(node *Node, pre **Node) {
	if node == nil {
		return
	}
	if node.LeftC != nil {
		Core(node.LeftC, pre)
	}

	if pre != nil && *pre != nil {
		node.LeftC = *pre
		(*pre).RightC = node
	}
	*pre = node
	if node.RightC != nil {
		Core(node.RightC, pre)
	}
}
