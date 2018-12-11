package lcs

func DynamicLcs(s1, s2 string) string {
	t := make([][]int, len(s1))

	for i := 0; i < len(s1); i++ {
		t[i] = make([]int, len(s2))
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				t[i][j] = getTableValue(t, i-1, j-1) + 1
			} else {
				if getTableValue(t, i-1, j) > getTableValue(t, i, j-1) {
					t[i][j] = getTableValue(t, i-1, j)
				} else {
					t[i][j] = getTableValue(t, i, j-1)
				}
			}
		}
	}

	if t[len(s1)-1][len(s2)-1] == 0 {
		return ""
	}
	return AssembleLcs(s1, s2, t, len(s1)-1, len(s2)-1, "")
}

func AssembleLcs(s1, s2 string, t [][]int, i1, i2 int, lcs string) string {
	if i1 == 0 || i2 == 0 && getTableValue(t, i1, i2) != 0 {
		return lcs
	}
	if getTableValue(t, i1, i2) == getTableValue(t, i1-1, i2) {
		return AssembleLcs(s1, s2, t, i1-1, i2, lcs)
	} else if getTableValue(t, i1, i2) == getTableValue(t, i1-1, i2) {
		return AssembleLcs(s1, s2, t, i1, i2-1, lcs)
	}
	lcs = string(s1[i1]) + lcs
	return AssembleLcs(s1, s2, t, i1-1, i2-1, lcs)
}

func getTableValue(t [][]int, i, j int) int {
	if i < 0 || j < 0 {
		return 0
	}
	return t[i][j]
}
