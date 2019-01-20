package code

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}
	if length == 1 {
		return strs[0]
	}
	//找到最短的长度,公共前缀不会超过这个长度
	minLength := len(strs[0])
	for i := 1; i < length; i++ {
		if len(strs[i]) < minLength {
			minLength = len(strs[i])
		}
	}
	//循环遍历
	result := ""
	for i := 0; i < minLength; i++ {
		tmp := strs[0][i]
		for j := 1; j < length; j++ {
			if strs[j][i] != tmp {
				return result
			}
		}
		result += string(tmp)
	}
	return result
}
