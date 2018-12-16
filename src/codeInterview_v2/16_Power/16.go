package code_16

// 面试题16：数值的整数次方
// 题目：实现函数double Power(double base, int exponent)，求base的exponent
// 次方。不得使用库函数，同时不需要考虑大数问题。

func Power(base float64, exponent int) float64 {
	if exponent == 0 {
		return 1
	}
	if base == 0 {
		return 0
	}
	absExponent := exponent
	if exponent < 0 {
		absExponent = -absExponent
	}
	result := float64(1)
	for i := 0; i < absExponent; i++ {
		result *= base
	}
	if exponent < 0 {
		result = 1 / result
	}
	return result
}
