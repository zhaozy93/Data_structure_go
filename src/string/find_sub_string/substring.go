package find_sub_string

func ViolenceSearchSubstring(text, sub string) (index int) {
	index = -1
	for i := 0; i < len(text)-len(sub)+1; i++ {
		for j := 0; j < len(sub); j++ {
			if text[i+j] != sub[j] {
				break
			}
			if j == len(sub)-1 {
				return i
			}
		}
	}
	return
}

func KMPSubstring(text, sub string) (index int) {
	index = -1
	pat := getNext(sub)
	j := 0
	for i := 0; i < len(text); {
		if j == -1 || text[i] == sub[j] {
			j++
			i++
			if j == len(sub)-1 {
				return i
			}
		} else {
			j = pat[j]
		}
	}
	return index
}

func getNext(sub string) []int {
	p := make([]int, len(sub))
	p[0] = -1
	// quick loop
	i := 1
	k := -1
	for i < len(p) {
		if k == -1 || sub[i] == sub[k] {
			k++
			p[i] = k
			i++
		} else {
			k = p[k]
		}
	}

	// naive loop
	// for i := 1; i < len(p); i++ {
	// 	if p[i-1] == -1 || sub[i] == sub[p[i-1]+1] {
	// 		p[i] = p[i-1] + 1
	// 	} else {
	// 		k := p[i-1]
	// 		for {
	// 			if k == -1 {
	// 				p[i] = 0
	// 				break
	// 			}
	// 			if sub[i] == sub[k+1] {
	// 				p[i] = p[k] + 1
	// 				break
	// 			} else {
	// 				k = p[k-1]
	// 			}
	// 		}
	// 	}
	// }
	return p
}

// 更容易理解的KMP
func Kmp(s, p string) int {
	pp := make([]int, len(p))
	pp[0] = 0
	for i := 1; i < len(p); i++ {
		j := pp[i-1]
		for {
			if p[i] == p[j] {
				pp[i] = j + 1
				break
			} else {
				if j == 0 {
					pp[i] = 0
					break
				}
				j = pp[j]
			}
		}
	}

	i, j := 0, 0
	for i < len(s) && j < len(p) {
		if s[i] == p[j] {
			i++
			j++
		} else {
			if pp[j] == 0 {
				i++
				continue
			}
			j = pp[j] - 1
		}
	}
	if j == len(p) {
		return i - j
	}
	return -1
}
