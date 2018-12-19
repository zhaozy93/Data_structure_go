package code_45

import (
	"strings"
)

// 面试题45：把数组排成最小的数
// 题目：输入一个正整数数组，把数组里所有数字拼接起来排成一个数，打印能拼
// 接出的所有数字中最小的一个。例如输入数组{3, 32, 321}，则打印出这3个数
// 字能排成的最小数字321323。

func SortArrayForMinNumber(nums []string) string {
	mirror := make([]string, len(nums))
	copy(mirror, nums)
	Core(mirror, 0, len(mirror)-1)
	return strings.Join(mirror, "")
}

func Core(nums []string, start, end int) {
	if start == end {
		return
	}
	left := start
	for i := start; i < end; i++ {
		if compare(nums[i], nums[end]) {
			continue
		}
		if i != left {
			swap(nums, i, left)
		}
		left++
	}
	swap(nums, left, end)

	if left > start {
		Core(nums, start, left-1)
	}
	if left < end {
		Core(nums, left+1, end)
	}
}

// i> j true
func compare(i string, j string) bool {
	lenI := len(i)
	lenJ := len(j)
	indexI := 0
	indexJ := 0
	for indexI < lenI && indexJ < lenJ {
		if i[indexI] == j[indexJ] {
			indexI++
			indexJ++
			continue
		}
		if i[indexI] > j[indexJ] {
			return true
		}
		return false
	}
	if lenI < lenJ {
		indexI = 0
	} else {
		indexJ = 0
	}
	for indexI < lenI && indexJ < lenJ {
		if i[indexI] == j[indexJ] {
			indexI++
			indexJ++
			continue
		}
		if i[indexI] > j[indexJ] {
			return true
		}
		return false
	}
	return false
}
func swap(nums []string, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

var Atoi map[byte]int = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}
