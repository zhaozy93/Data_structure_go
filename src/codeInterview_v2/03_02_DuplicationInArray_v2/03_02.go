package code_03_02

// 面试题3（二）：不修改数组找出重复的数字
// 题目：在一个长度为n+1的数组里的所有数字都在1到n的范围内，所以数组中至
// 少有一个数字是重复的。请找出数组中任意一个重复的数字，但不能修改输入的
// 数组。例如，如果输入长度为8的数组{2, 3, 5, 4, 3, 2, 6, 7}，那么对应的
// 输出是重复的数字2或者3。

func FindDuplicationNoEdit(nums []int) int {
	start := 1
	end := len(nums) - 1
	for start <= end {
		mid := (start + end) / 2
		cnt := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] <= mid && nums[i] >= start {
				cnt++
			}
		}
		if start == end {
			if cnt > 1 {
				return start
			} else {
				break
			}
		}
		if cnt <= mid-start+1 {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return -1
}
