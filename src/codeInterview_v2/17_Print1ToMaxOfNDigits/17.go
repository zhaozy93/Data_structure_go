package code_17

import (
	"fmt"
	"strconv"
)

// 面试题17：打印1到最大的n位数
// 题目：输入数字n，按顺序打印出从1最大的n位十进制数。比如输入3，则
// 打印出1、2、3一直到最大的3位数即999。

func PrintMaxOfNDigits(n int) {
	if n <= 0 {
		fmt.Println("error")
	}
	num := make([]int, n)
	for {
		PrintNumReverse(num)
		overflow := false
		//	把数字倒叙排列
		increment := 1
		nTakeOver := 0
		index := 0
		for {
			incr := num[index] + increment + nTakeOver
			if incr >= 10 {
				if index == len(num)-1 {
					overflow = true
					break
				}
				num[index] = (incr) % 10
				nTakeOver = (incr) / 10
				increment = 0
				index++
			} else {
				num[index] = incr
				break
			}
		}

		if overflow {
			break
		}
	}
}

func PrintNumReverse(num []int) {
	str := ""
	for i := 0; i < len(num); i++ {
		str = strconv.Itoa(num[i]) + str
	}
	index := 0
	for ; index < len(str); index++ {
		if str[index] != '0' {
			break
		}
	}
	if index == len(str) {
		fmt.Println("0")
	} else {
		fmt.Println(str[index:])
	}
}

func PrintMaxOfNDigits_Recur(n int) {
	for i := 0; i < 10; i++ {
		PrintMaxOfNDigits_RecurCore(n, strconv.Itoa(i))
	}
}

func PrintMaxOfNDigits_RecurCore(n int, current string) {
	if len(current) == n {
		fmt.Println(current)
		return
	}
	for i := 0; i < 10; i++ {
		PrintMaxOfNDigits_RecurCore(n, current+strconv.Itoa(i))
	}
}
