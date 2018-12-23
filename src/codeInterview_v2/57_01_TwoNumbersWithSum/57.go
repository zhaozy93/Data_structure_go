package code_57

// 面试题57（一）：和为s的两个数字
// 题目：输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们
// 的和正好是s。如果有多对数字的和等于s，输出任意一对即可。
func TwoNumbersWithSum(nums []int, sum int) (index1, index2 int, ok bool) {
	start := 0
	end := len(nums) - 1
	for {
		if start == end {
			break
		}
		if nums[start]+nums[end] > sum {
			end--
		} else if nums[start]+nums[end] < sum {
			start++
		} else {
			index1 = start
			index2 = end
			ok = true
			return
		}
	}
	return
}
