package code_50

func FirstNotRepeatingChar(s string) rune {
	cnt := make(map[rune]int)
	for _, v := range s {
		cnt[v]++
	}
	for _, v := range s {
		if cnt[v] == 1 {
			return v
		}
	}
	return '0'
}
