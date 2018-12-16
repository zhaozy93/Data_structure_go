package code_07

// 重构二叉树
type Node struct {
	Val    string
	leftC  *Node
	rightC *Node
}

func ConstructBinaryTreeByPreandInOrder(preOrder []string, inOrder []string) *Node {
	return constructBinaryTreeByPreandInOrder(preOrder, 0, len(preOrder)-1, inOrder, 0, len(preOrder)-1)
}

func constructBinaryTreeByPreandInOrder(preOrder []string, preIndexStart int, preIndexEnd int, inOrder []string, inOrderIndexStart int, inOrderIndexEnd int) *Node {
	root := &Node{Val: preOrder[preIndexStart]}
	i := inOrderIndexStart
	for ; i <= inOrderIndexEnd; i++ {
		if inOrder[i] == preOrder[preIndexStart] {
			break
		}
	}

	if i-inOrderIndexStart > 0 {
		root.leftC = constructBinaryTreeByPreandInOrder(preOrder, preIndexStart+1, preIndexStart+i-inOrderIndexStart, inOrder, inOrderIndexStart, i-1)
	}

	if inOrderIndexEnd-i > 0 {
		root.rightC = constructBinaryTreeByPreandInOrder(preOrder, preIndexStart+i+1-inOrderIndexStart, preIndexEnd, inOrder, i+1, inOrderIndexEnd)
	}
	return root
}
