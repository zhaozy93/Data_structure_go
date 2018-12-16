package code_10

func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	left := 0
	right := 1
	result := 1
	for i := 2; i <= n; i++ {
		result = left + right
		left = right
		right = result
	}
	return result
}
