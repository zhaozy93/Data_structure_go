package code_19

// 面试题19：正则表达式匹配
// 题目：请实现一个函数用来匹配包含'.'和'*'的正则表达式。模式中的字符'.'
// 表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）。在本题
// 中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"
// 和"ab*ac*a"匹配，但与"aa.a"及"ab*a"均不匹配。

func RegularExpressions(s string, pattern string) bool {
	return RegularExpressionsCore(s, pattern, 0, 0)
}

func RegularExpressionsCore(s, pattern string, sIndex, pIndex int) bool {
	// check end condition
	if sIndex == len(s) || pIndex == len(pattern) {
		if sIndex == len(s) && pIndex == len(pattern) {
			return true
		}
		if sIndex != len(s) {
			return false
		}
		if pIndex != len(pattern) {
			return checkZeroLen(pattern, pIndex)
		}
	}

	// 判断下一个是不是*
	nextIs8 := false
	if pIndex < len(pattern)-1 {
		nextIs8 = (pattern[pIndex+1] == '*')
	}
	// 下一个字符不是*
	if !nextIs8 {
		if pattern[pIndex] == s[sIndex] || pattern[pIndex] == '.' {
			return RegularExpressionsCore(s, pattern, sIndex+1, pIndex+1)
		}
		return false
	}

	// 下一个字符是*
	// choice one 跳过这次 ch*组合  即*=0
	if pIndex+2 < len(pattern) {
		if RegularExpressionsCore(s, pattern, sIndex, pIndex+2) {
			return true
		}
	}

	// choice two *重复N次
	if pattern[pIndex] == s[sIndex] || pattern[pIndex] == '.' {
		return RegularExpressionsCore(s, pattern, sIndex+1, pIndex)
	}
	return false
}

func checkZeroLen(pattern string, pIndex int) bool {
	char := true
	for {
		if pIndex == len(pattern) {
			break
		}
		if char {
			pIndex++
			char = false
			continue
		}
		if !char {
			if pattern[pIndex] == '*' {
				pIndex++
				char = true
				continue
			}
			return false
		}
	}
	return char
}
