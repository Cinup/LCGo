package code
//滑动窗口法
func lengthOfLongestSubstring_2(s string) int {
	len := len(s)
	isexit := [128]int8{0}
	max := 0
	i, j := 0, 0
	for i < len && j < len {
		if isexit[s[j]] == 0 {
			isexit[s[j]] = 1
			j++
			if j-i > max {
				max = j - i
			}
		} else {
			isexit[s[i]] = 0
			i++
		}
	}
	return max
}
