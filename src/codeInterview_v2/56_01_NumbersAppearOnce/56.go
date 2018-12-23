package code_56

// 面试题56（一）：数组中只出现一次的两个数字
// 题目：一个整型数组里除了两个数字之外，其他的数字都出现了两次。请写程序
// 找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。

func FindNumsAppearOnce(nums []int) []int {
	result := 0
	for val := range nums {
		result = result ^ val
	}

	index := 0
	for {
		if result == (result>>1)*2 {
			index++
		} else {
			break
		}
		if result == 0 {
			break
		}
	}

	left := make([]int, 0)
	right := make([]int, 0)

	for _, val := range nums {
		num := val >> uint(index)
		if num&1 == 0 {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}

	num1 := 0
	num2 := 0
	for _, val := range left {
		num1 = num1 ^ val
	}
	for _, val := range right {
		num2 = num2 ^ val
	}
	return []int{num1, num2}
}
