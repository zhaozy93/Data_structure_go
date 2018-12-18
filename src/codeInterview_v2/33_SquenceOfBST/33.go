package code_33

// 面试题33：二叉搜索树的后序遍历序列
// 题目：输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。
// 如果是则返回true，否则返回false。假设输入的数组的任意两个数字都互不相同。

func VerifySquenceOfBST(num []int) bool {
	if len(num) == 0 {
		return true
	}
	root := num[len(num)-1]
	index := 0
	for ; index < len(num)-2; index++ {
		if num[index] > root {
			break
		}
	}

	for j := index; j < len(num)-2; j++ {
		if num[j] < root {
			return false
		}
	}

	return VerifySquenceOfBST(num[:index]) && VerifySquenceOfBST(num[index:len(num)-1])
}
