package code_58

// 面试题58（一）：翻转单词顺序
// 题目：输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。
// 为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，
// 则输出"student. a am I"。

type Stack struct {
	vals []string
}

func (s *Stack) Push(val string) {
	if s.vals == nil {
		s.vals = make([]string, 0)
	}
	s.vals = append(s.vals, val)
}
func (s *Stack) Pop() (string, bool) {
	if s.vals == nil || len(s.vals) == 0 {
		return "", false
	}
	val := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	if len(s.vals) == 0 {
		s.vals = nil
	}
	return val, true
}

func ReverseSentence(str string) (reverse string) {
	s1 := &Stack{}
	start := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			s1.Push(str[start:i])
			start = i
		}
	}
	s1.Push(str[start:len(str)])
	for {
		val, ok := s1.Pop()
		if !ok {
			break
		}
		reverse = reverse + " " + val
	}
	return
}
