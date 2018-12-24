package code_60

// 面试题60：n个骰子的点数
// 题目：把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s
// 的所有可能的值出现的概率。

func PrintProbability_Solution1(n int) map[int]float64 {
	p := make(map[int]float64)
	Core(n, 1, 0, p)
	totalTimes := float64(0)
	for _, times := range p {
		totalTimes += times
	}
	for sum, times := range p {
		p[sum] = times / totalTimes
	}
	return p
}

func Core(n int, step int, currentSum int, p map[int]float64) {
	if step > n {
		p[currentSum]++
		return
	}
	for i := 1; i <= 6; i++ {
		Core(n, step+1, currentSum+i, p)
	}
}
