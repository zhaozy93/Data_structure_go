package sort

import (
	"math"
)

// 都是从小到大排序
func BubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a)-i; j++ {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

func SelectionSort(a []int) {
	for i := 0; i < len(a); i++ {
		idx := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[idx] {
				idx = j
			}
		}
		a[i], a[idx] = a[idx], a[i]
	}
}

func InsertSort(a []int) {
	for i := 1; i < len(a); i++ {
		temp := a[i]
		for j := i - 1; j >= 0; j-- {
			if temp < a[j] {
				a[j+1] = a[j]
			} else {
				a[j+1] = temp
				break
			}
		}
	}
}

func ShellSort(a []int) {
	// 间隔序列定死
	// 动态间隔不考虑
	k := []int{4, 3, 2, 1}
	for i := 0; i < len(k); i++ {
		step := k[i]
		for m := step; m < len(a); m++ {
			temp := a[m]
			for n := m - step; n >= 0; n -= step {
				if temp < a[n] {
					a[n+step], a[n] = a[n], a[n+step]
				} else {
					a[n+step] = temp
					break
				}
			}
		}
	}
}

func MergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	mid := len(a) / 2
	return merge(MergeSort(a[:mid]), MergeSort(a[mid:]))
}

func merge(a, b []int) []int {
	i, j, m := 0, 0, 0
	s := make([]int, len(a)+len(b))
	for {
		if m == len(s) {
			break
		}
		if i == len(a) {
			s[m] = b[j]
			j++
			m++
		} else if j == len(b) {
			s[m] = a[i]
			i++
			m++
		} else if a[i] < b[j] {
			s[m] = a[i]
			i++
			m++
		} else {
			s[m] = b[j]
			j++
			m++
		}
	}
	return s
}

func QuickSort(a []int) {
	if len(a) < 2 {
		return
	}
	temp, l := a[0], 0
	for i := 1; i < len(a); i++ {
		if a[i] < temp {
			a[l] = a[i]
			a[i] = a[l+1]
			l++
		}
	}
	a[l] = temp
	QuickSort(a[:l])
	QuickSort(a[l+1:])
}

func HeapSort(a []int) {
	if len(a) < 2 {
		return
	}
	// 构建最大堆
	for i := 1; i < len(a)-1; i++ {
		j := i
		for {
			if j <= 0 {
				break
			}
			p := int(math.Ceil(float64(j)/2 - 1))
			if a[j] > a[p] {
				a[j], a[p] = a[p], a[j]
			}
			j = p
		}
	}

	// 从堆上一项项取出来
	for i := len(a) - 1; i >= 0; i-- {
		a[i], a[0] = a[0], a[i]
		p := 0
		for {
			l := (p+1)*2 - 1
			r := (p + 1) * 2
			if l >= i {
				break
			}
			if r == i {
				if a[p] < a[l] {
					a[l], a[p] = a[p], a[l]
				}
				break
			}

			if a[p] > a[r] && a[p] > a[l] {
				break
			}
			if a[l] < a[r] {
				a[p], a[r] = a[r], a[p]
				p = r
			} else {
				a[p], a[l] = a[l], a[p]
				p = l
			}
		}
	}
}
