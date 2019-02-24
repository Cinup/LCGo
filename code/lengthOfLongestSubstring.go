package code
//暴力搜索法
func lengthOfLongestSubstring(s string) int {
	len := len(s)
	isexit := [128]int8{0}
	max := 0
	for i := 0; i < len; i++ {
		isexit[s[i]] = 1
		num := 1
		for j := i + 1; j < len; j++ {
			if isexit[s[j]] != 1 {
				isexit[s[j]] = 1
				num ++
			} else {
				break
			}
		}
		if num > max {
			max = num
		}
		isexit = [128]int8{0}
	}
	return max
}