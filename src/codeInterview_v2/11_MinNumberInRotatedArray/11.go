package code_11

// 面试题11：旋转数组的最小数字
// 题目：把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
// 输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如数组
// {3, 4, 5, 1, 2}为{1, 2, 3, 4, 5}的一个旋转，该数组的最小值为1。

func FindMinNum(num []int) int {
	left := 0
	right := len(num) - 1
	// 实际未旋转
	if num[left] < num[right] {
		return num[left]
	}
	for {
		if right-left == 1 {
			break
		}
		mid := (left + right) / 2
		if num[mid] == num[left] && num[mid] == num[right] {
			min := num[left]
			for i := left; i <= right; i++ {
				if num[i] < min {
					min = num[i]
				}
			}
			return min
		}
		if num[mid] < num[left] {
			right = mid
			continue
		}
		if num[mid] > num[left] {
			left = mid
			continue
		}
	}
	if num[right] < num[left] {
		return num[right]
	}
	return num[left]
}
