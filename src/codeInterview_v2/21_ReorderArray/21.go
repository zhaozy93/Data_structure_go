package code_21

// 面试题21：调整数组顺序使奇数位于偶数前面
// 题目：输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有
// 奇数位于数组的前半部分，所有偶数位于数组的后半部分。

func ReorderArray(s []int) {
	i := 0
	j := len(s) - 1
	for {
		if i == j {
			break
		}

		for {
			if s[i]%2 == 0 || i == j {
				break
			}
			i++
		}
		for {
			if s[j]%2 == 1 || j == i {
				break
			}
			j--
		}
		if i != j {
			swap(s, i, j)
		}
	}
}

func swap(s []int, i, j int) {
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
}
