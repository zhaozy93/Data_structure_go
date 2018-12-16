package code_11

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, start, end int) {
	if start == end {
		return
	}
	mid := (start + end) / 2
	midVal := nums[mid]
	swap(nums, mid, end)
	small := start
	for index := start; index < end; index++ {
		if nums[index] <= midVal {
			if small != index {
				swap(nums, index, small)
			}
			small++
		}
	}
	swap(nums, small, end)

	if small-1 > start {
		quickSort(nums, start, small-1)
	}
	if small+1 < end {
		quickSort(nums, small+1, end)
	}
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
