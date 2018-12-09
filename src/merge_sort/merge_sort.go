package merge_sort

func MergeSort(in []int) []int {
	mergeSort(in, 0, len(in)-1)
	return in
}
func mergeSort(in []int, p, q int) {
	if q-p > 1 {
		mid := (p + q) / 2
		mergeSort(in, p, mid)
		mergeSort(in, mid+1, q)
		merge(in, p, q, mid)
	} else {
		merge(in, p, q, -1)
	}
	return
}

func merge(in []int, p, q, mid int) {
	if p == q {
		return
	} else if mid == -1 {
		if in[p] > in[q] {
			in[p], in[q] = in[q], in[p]
		}
		return
	} else {
		temp := make([]int, q-p+1)
		left := 0
		right := 0
		for i := 0; i < len(temp); i++ {
			if left == mid+1-p {
				temp[i] = in[mid+1+right]
				right++
				continue
			}
			if right == q-mid {
				temp[i] = in[p+left]
				left++
				continue
			}
			leftV := in[p+left]
			rightV := in[mid+1+right]
			if leftV < rightV {
				temp[i] = leftV
				left++
				continue
			}
			temp[i] = rightV
			right++
		}
		copy(in[p:q+1], temp)
	}
}
