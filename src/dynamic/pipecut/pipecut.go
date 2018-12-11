package pipecut

import "fmt"

var Price map[int]int = map[int]int{
	1:  1,
	2:  5,
	3:  8,
	4:  9,
	5:  10,
	6:  17,
	7:  17,
	8:  20,
	9:  24,
	10: 30,
}

func RecursiveMethod(n int) int {
	if n == 0 {
		return 0
	}
	max := 0
	for i := 1; i <= n; i++ {
		price := Price[i] + RecursiveMethod(n-i)
		if price > max {
			max = price
		}
	}
	return max
}

var memNote map[int]int = map[int]int{
	0: 0,
}

func MemorandumMethod(n int) int {
	if _, ok := memNote[n]; ok {
		fmt.Printf("命中缓存%d, price: %d\n", n, memNote[n])
		return memNote[n]
	}
	max := 0
	for i := 1; i <= n; i++ {
		price := Price[i] + MemorandumMethod(n-i)
		if price > max {
			max = price
		}
	}
	if _, ok := memNote[n]; !ok {
		memNote[n] = max
	}
	return max
}

// 动态规划思想， 记录你之前做了什么 记录之前的结果 不要进行重复计算
func DynamicMethod(n int) int {
	note := make(map[int]int)
	note[0] = 0
	note[1] = 1
	for i := 1; i <= n; i++ {
		price := 0
		for j := 1; j <= i; j++ {
			if Price[j]+note[i-j] > price {
				price = Price[j] + note[i-j]
			}
		}
		note[i] = price
	}
	return note[n]
}
