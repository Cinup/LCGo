package code

import "strings"

func findWords(words []string) []string {
	var keyboard = []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	var result []string
	for i := 0; i < len(words); i++ {
		//找出第一个字符的行数
		line := keyBoardLine(keyboard, string(words[i][0]))
		flag := true
		for _, j := range words[i] {
			//如果有元素不是同一行,直接结束
			if line^keyBoardLine(keyboard, strings.ToLower(string(j))) != 0 {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, words[i])
		}
	}
	return result
}
func keyBoardLine(keyboard []string, substr string) int {
	//返回字符在键盘中的行数
	if strings.Contains(keyboard[0], substr) {
		return 1
	}
	if strings.Contains(keyboard[1], substr) {
		return 2
	} else {
		return 3
	}
}
