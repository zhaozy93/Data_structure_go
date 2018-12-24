package code_63

func maxProfit(nums []int) int {
	diff := make([]int, len(nums))
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - min
		if nums[i] < min {
			min = nums[i]
		}
	}
	max := diff[1]
	for _, val := range diff {
		if val > max {
			max = val
		}
	}
	return max
}
