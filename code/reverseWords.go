package code

func reverseWords(s string) string {
	if len(s) < 1 {
		return s
	}
	result := []byte(s)
	start := 0
	for {
		end := findBlank(s, start)
		if end == -1 {
			break
		} else {
			reverse(result, start, end-1)
		}
		start = end + 1
	}
	//反转最后一个单词
	reverse(result, start, len(result)-1)
	return string(result)
}

//分词函数
func findBlank(s string, start int) int {
	for start < len(s) {
		if s[start] == ' ' {
			return start
		}
		start++
	}
	return -1
}

//单词反转函数
func reverse(s []byte, start, end int) {
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
