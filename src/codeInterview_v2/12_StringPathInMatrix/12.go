package code_12

// 面试题12：矩阵中的路径
// 题目：请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有
// 字符的路径。路径可以从矩阵中任意一格开始，每一步可以在矩阵中向左、右、
// 上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入
// 该格子。例如在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字
// 母用下划线标出）。但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个
// 字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。
// A B T G
// C F C S
// J D E H

func HasPath(path string, martix [][]byte) bool {
	for i := 0; i < len(martix); i++ {
		for j := 0; j < len(martix[0]); j++ {
			visited := make([][]bool, len(martix))
			for i := 0; i < len(visited); i++ {
				visited[i] = make([]bool, len(martix[0]))
			}
			if HasPathCore(path, martix, 0, i, j, visited) == true {
				return true
			}
		}
	}
	return false
}

func HasPathCore(path string, martix [][]byte, index, row, col int, visited [][]bool) bool {
	if martix[row][col] != path[index] {
		return false
	}
	if visited[row][col] == true {
		return false
	}
	if index == len(path)-1 {
		return true
	}
	visited[row][col] = true
	if row > 0 {
		if HasPathCore(path, martix, index+1, row-1, col, visited) {
			return true
		}
	}
	if row < len(martix)-1 {
		if HasPathCore(path, martix, index+1, row+1, col, visited) {
			return true
		}
	}
	if col > 0 {
		if HasPathCore(path, martix, index+1, row, col-1, visited) {
			return true
		}
	}
	if col < len(martix[0])-1 {
		if HasPathCore(path, martix, index+1, row, col+1, visited) {
			return true
		}
	}
	visited[row][col] = false
	return false
}
