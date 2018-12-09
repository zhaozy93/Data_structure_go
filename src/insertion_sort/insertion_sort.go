package insertion_sort

func InsertionSort(in []int) []int {
	for i := 1; i < len(in); i++ {
		temp := in[i]
		j := i - 1
		for ; j >= 0; j-- {
			if temp > in[j] {
				break
			}
		}
		j++
		copy(in[j+1:i+1], in[j:i])
		in[j] = temp
	}
	return in
}
