package selection_sort

func SelectionSort(in []int) []int {
	for i := 0; i < len(in); i++ {
		smIndex := i
		for j := i + 1; j < len(in); j++ {
			if in[j] < in[smIndex] {
				smIndex = j
			}
		}
		if smIndex != i {
			temp := in[i]
			in[i] = in[smIndex]
			in[smIndex] = temp
		}
	}
	return in
}
