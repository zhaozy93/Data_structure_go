package code_39

// 面试题39：数组中出现次数超过一半的数字
// 题目：数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。例
// 如输入一个长度为9的数组{1, 2, 3, 2, 2, 2, 5, 4, 2}。由于数字2在数组中
// 出现了5次，超过数组长度的一半，因此输出2。

// 假设可以修改数据内容
func MoreThanHalfNumber(nums []int) int {
	nums_mirror := make([]int, len(nums))
	copy(nums_mirror, nums)

	midNum := len(nums_mirror) / 2
	index := -1
	start := 0
	end := len(nums_mirror) - 1
	for index != midNum {
		if index < midNum {
			start = index + 1
		} else {
			end = index - 1
		}
		index = Core(nums_mirror, start, end)
	}
	return nums_mirror[midNum]
}

func Core(nums []int, start, end int) int {
	if start == end {
		return 0
	}
	left := start
	for i := start; i < end; i++ {
		if nums[i] > nums[end] {
			continue
		}
		if i != left {
			swap(nums, i, left)
		}
		left++
	}
	swap(nums, left, end)
	return left
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
