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

// 方法二 推荐方法
func QuickSort2(nums []int, start, end int) {
	fmt.Println(start, end)
	if start >= end {
		return
	}
	midValue, i := nums[start], start+1
	idx1, idx2 := start, end
	for i <= end {
		fmt.Println(start, end, i, nums)
		if nums[i] > midValue {
			nums[i], nums[end] = nums[end], nums[i]
			end--
		} else {
			nums[i], nums[start] = nums[start], nums[i]
			i++
			start++
		}
	}
	QuickSort2(nums, idx1, start-1)
	QuickSort2(nums, start+1, idx2)
}
