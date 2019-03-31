package code

func lengthOfLastWord(s string) int {
	num := 0
	length := 0
	//先过滤末尾的空格
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			length++
		} else {
			break
		}
	}
	//空字符串
	if len(s)-length == 0 {
		return 0
	}
	for i := len(s) - length - 1; i >= 0; i-- {
		if s[i] != ' ' {
			num++
		} else {
			return num
		}
	}
	return num
}
