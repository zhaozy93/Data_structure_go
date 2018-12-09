package quick_sort

func QuickSort(in []int) []int {
	quicksort(in, 0, len(in)-1)
	return in
}
func quicksort(in []int, p, r int) {
	if p >= r {
		return
	}

	mid := in[r]
	s := p
	e := r
	isSearchBig := true
	for {
		if s == e {
			in[s] = mid
			break
		}
		if isSearchBig {
			if in[s] < mid {
				s++
				continue
			} else {
				in[e] = in[s]
				e--
				isSearchBig = false
				continue
			}
		} else {
			if in[e] > mid {
				e--
				continue
			} else {
				in[s] = in[e]
				s++
				isSearchBig = true
				continue
			}
		}
	}
	quicksort(in, p, s-1)
	quicksort(in, s+1, r)
}
