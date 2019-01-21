package code

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	s := "11"
	for i := 2; i < n; i++ {
		pre := s[0]
		num := 1
		var tmps []uint8
		index := 0
		for j := 1; j < len(s); j++ {
			if s[j] == pre {
				num++
			} else {
				tmps[index] = uint8(num)
				index++
				tmps[index] = pre
				pre = s[j]
			}
		}
		s = string(tmps)
	}
	return s
}
