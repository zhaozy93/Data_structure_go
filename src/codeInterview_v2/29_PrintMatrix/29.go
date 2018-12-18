package code_29

import "fmt"

// 面试题29：顺时针打印矩阵
// 题目：输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

func PrintMartix(num [][]int) {
	if num == nil || len(num) == 0 || len(num[0]) == 0 {
		return
	}
	row := 0
	col := 0
	for {
		if row*2 < len(num) && col*2 < len(num[0]) {
			PrintCircle(num, row, col)
			row++
			col++
			continue
		}
		break
	}
}

func PrintCircle(num [][]int, row, col int) {
	// 顺时针打印
	rowIndex := row
	colIndex := col
	// 横着走 左上--> 右上
	for ; colIndex <= len(num[0])-col-1; colIndex++ {
		fmt.Print(num[rowIndex][colIndex], ",")
	}
	colIndex--

	if (rowIndex+1)*2 > len(num)-1 {
		return
	}

	// 竖着走 右上 --> 右下
	for rowIndex++; rowIndex <= len(num)-row-1; rowIndex++ {
		fmt.Print(num[rowIndex][colIndex], ",")
	}
	rowIndex--

	// 横着走 右下 --> 左下
	for colIndex--; colIndex >= col; colIndex-- {
		fmt.Print(num[rowIndex][colIndex], ",")
	}
	colIndex++

	if rowIndex-1 <= row {
		return
	}

	// 竖着走 左下 --> 左上
	for rowIndex--; rowIndex > row; rowIndex-- {
		fmt.Print(num[rowIndex][colIndex], ",")
	}

}
