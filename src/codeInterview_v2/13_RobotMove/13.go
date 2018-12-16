package code_13

// 面试题13：机器人的运动范围
// 题目：地上有一个m行n列的方格。一个机器人从坐标(0, 0)的格子开始移动，它
// 每一次可以向左、右、上、下移动一格，但不能进入行坐标和列坐标的数位之和
// 大于k的格子。例如，当k为18时，机器人能够进入方格(35, 37)，因为3+5+3+7=18。
// 但它不能进入方格(35, 38)，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

func RobertMove(k int, martix [][]int) int {
	visited := make([][]bool, len(martix))
	for i := 0; i < len(martix); i++ {
		visited[i] = make([]bool, len(martix[0]))
	}
	count := 0
	RobertMoveCore(k, martix, 0, 0, visited, &count)
	return count
}

func RobertMoveCore(k int, martix [][]int, row, col int, visited [][]bool, count *int) {
	if row < 0 || row == len(martix) {
		return
	}
	if col < 0 || col == len(martix[0]) {
		return
	}
	if visited[row][col] == true {
		return
	}
	visited[row][col] = true
	if check(k, row, col) {
		(*count)++
	}
	RobertMoveCore(k, martix, row-1, col, visited, count)
	RobertMoveCore(k, martix, row+1, col, visited, count)
	RobertMoveCore(k, martix, row, col-1, visited, count)
	RobertMoveCore(k, martix, row, col+1, visited, count)

}

func check(k, i, j int) bool {
	sum := 0
	for {
		if i == 0 {
			break
		}
		sum += i % 10
		i = i / 10
	}
	for {
		if j == 0 {
			break
		}
		sum += j % 10
		j = j / 10
	}
	if sum <= k {
		return true
	}
	return false
}
