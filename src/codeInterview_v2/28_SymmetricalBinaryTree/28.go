package code_28

// 面试题28：对称的二叉树
// 题目：请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和
// 它的镜像一样，那么它是对称的。

type Node struct {
	Val    string
	LeftC  *Node
	RightC *Node
}

func CheckIsSymmetryTree(root *Node) bool {
	return CheckIsSymmetryTreeCore(root, root)
	// preTraversalNode := make([]*string, 0)
	// preTraversal(root, &preTraversalNode)
	// reversePreTraversalNode := make([]*string, 0)
	// reversePreTraversal(root, &preTraversalNode)
	// for i, n := range preTraversalNode {
	// 	if preTraversalNode[i] != reversePreTraversalNode[i] {
	// 		return false
	// 	}
	// }
	// return true
}

func CheckIsSymmetryTreeCore(root *Node, root2 *Node) bool {
	if root == nil && root2 == nil {
		return true
	}
	if root == nil || root2 == nil {
		return false
	}

	if root.Val != root2.Val {
		return false
	}

	return CheckIsSymmetryTreeCore(root.LeftC, root2.RightC) && CheckIsSymmetryTreeCore(root.RightC, root2.LeftC)
}

// func preTraversal(n *Node, res *[]*string) {
// 	if n != nil {
// 		*res = append(*res, &n.Val)
// 		preTraversal(n.LeftC, res)
// 		preTraversal(n.RightC, res)
// 	} else {
// 		*res = append(*res, nil)
// 	}
// }

// func reversePreTraversal(n *Node, res *[]*string) {
// 	if n != nil {
// 		*res = append(*res, &n.Val)
// 		preTraversal(n.RightC, res)
// 		preTraversal(n.LeftC, res)
// 	} else {
// 		*res = append(*res, nil)
// 	}
// }
