package code_14

// 面试题14：剪绳子
// 题目：给你一根长度为n绳子，请把绳子剪成m段（m、n都是整数，n>1并且m≥1）。
// 每段的绳子的长度记为k[0]、k[1]、……、k[m]。k[0]*k[1]*…*k[m]可能的最大乘
// 积是多少？例如当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此
// 时得到最大的乘积18。

// 动态规划
func CuttingRopeDynamic(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	cache := make(map[int]int)
	cache[1] = 1
	cache[2] = 2
	cache[3] = 3

	for i := 4; i <= n; i++ {
		max := 0
		for j := 1; j <= i; j++ {
			if cache[j]*cache[i-j] > max {
				max = cache[j] * cache[i-j]
			}
		}
		cache[i] = max
	}
	return cache[n]
}
