package code

//优化滑动窗口
func lengthOfLongestSubstring_3(s string) int {
	len := len(s)
	isexit := [128]int{0}
	max := 0
	i, j := 0, 0
	for ; j < len; j++ {
		i = mmax(isexit[s[j]], i)
		max = mmax(max, j-i+1)
		isexit[s[j]] = j + 1
	}
	return max
}
func mmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
