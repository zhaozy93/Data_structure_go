package code_38

import (
	"fmt"
	"strings"
)

// 面试题38：字符串的排列
// 题目：输入一个字符串，打印出该字符串中字符的所有排列。例如输入字符串abc，
// 则打印出由字符a、b、c所能排列出来的所有字符串abc、acb、bac、bca、cab和cba。

func Permutation(s string) {
	if len(s) == 0 {
		return
	}
	ss := strings.Split(s, "")
	Core(ss, 0)
}

func Core(s []string, index int) {
	if index == len(s)-1 {
		fmt.Println(strings.Join(s, ""))
		return
	}

	for j := index; j < len(s); j++ {
		swap(s, index, j)
		Core(s, index+1)
		swap(s, index, j)
	}
}

func swap(s []string, i, j int) {
	if i == j {
		return
	}
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
}
