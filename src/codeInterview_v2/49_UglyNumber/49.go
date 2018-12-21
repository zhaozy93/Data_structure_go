package code_49

// 面试题49：丑数
// 题目：我们把只包含因子2、3和5的数称作丑数（Ugly Number）。求按从小到
// 大的顺序的第1500个丑数。例如6、8都是丑数，但14不是，因为它包含因子7。
// 习惯上我们把1当做第一个丑数。

func FindUglyNums(n int) []int {
	nums := make([]int, n)
	nums[0] = 1
	ulgyNumsIndex2 := 0
	ulgyNumsIndex3 := 0
	ulgyNumsIndex5 := 0
	for index := 1; index < n; {
		minIndex, val := min(nums[ulgyNumsIndex2]*2, nums[ulgyNumsIndex3]*3, nums[ulgyNumsIndex5]*5)
		switch minIndex {
		case 1:
			ulgyNumsIndex2++
		case 2:
			ulgyNumsIndex3++
		case 3:
			ulgyNumsIndex5++
		}
		if nums[index-1] == val {
			continue
		}
		nums[index] = val
		index++
	}
	return nums
}

func min(val1, val2, val3 int) (index, val int) {
	if val1 < val2 {
		index = 1
		val = val1
	} else {
		index = 2
		val = val2
	}

	if val3 < val {
		val = val3
		index = 3
	}
	return
}
