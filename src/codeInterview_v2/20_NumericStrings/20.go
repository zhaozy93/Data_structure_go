package code_20

// 面试题20：表示数值的字符串
// 题目：请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。例如，
// 字符串“+100”、“5e2”、“-123”、“3.1416”及“-1E-16”都表示数值，但“12e”、
// “1a3.14”、“1.2.3”、“+-5”及“12e+5.4”都不是

func isNumeric(s string) bool {
	hasSymbol := false
	hasInteger := false
	hasDecimal := false
	hasExpontentSymbol := false
	hasExpontent := false
	// 0-9  48-57
	// . 46
	// e 101 E 69
	// + 43  - 45
	for i := 0; i < len(s); i++ {
		if s[i] >= 48 && s[i] <= 57 {
			if hasDecimal {
				hasExpontent = true
			}
			continue
		}
		if s[i] == 43 || s[i] == 45 {
			if hasDecimal {
				if hasExpontentSymbol {
					return false
				}
				hasExpontentSymbol = true
				continue
			}
			if hasSymbol {
				return false
			}
			hasSymbol = true
			continue
		}
		if s[i] == 46 {
			if hasInteger {
				return false
			}
			hasSymbol = true
			hasInteger = true
			continue
		}
		if s[i] == 101 || s[i] == 69 {
			if hasDecimal {
				return false
			}
			hasSymbol = true
			hasInteger = true
			hasDecimal = true
			continue
		}
		return false
	}
	if hasDecimal && !hasExpontent {
		return false
	}
	return true
}
