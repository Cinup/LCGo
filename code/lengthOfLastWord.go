package code

func lengthOfLastWord(s string) int {
	//空字符串
	if len(s) == 0 {
		return 0
	}
	num := 0
	length := 0
	//先过滤末尾的空格
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			length++
			continue
		} else {
			break
		}
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