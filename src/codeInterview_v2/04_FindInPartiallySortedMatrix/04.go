package code_04

// 面试题4：二维数组中的查找
// 题目：在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按
// 照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个
// 整数，判断数组中是否含有该整数。

// 1  2  8   9
// 2  4  9   12
// 4  7  10  13
// 6  8  11  15

func FindInPartiallySortedMatrix(martix [][]int, num int) bool {
	rows := len(martix)
	cols := len(martix[0])

	rowIdx := 0
	colIdx := cols - 1
	for {
		if martix[rowIdx][colIdx] > num {
			colIdx--
			if colIdx == -1 {
				break
			}
			continue
		}
		if martix[rowIdx][colIdx] < num {
			rowIdx++
			if rowIdx == rows {
				break
			}
			continue
		}
		if martix[rowIdx][colIdx] == num {
			return true
		}
	}
	return false
}
