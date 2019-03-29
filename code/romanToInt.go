package code

func romanToInt(s string) int {
	m := make(map[byte]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000
	length := len(s)
	pre := 0
	//保存结果
	sum := 0
	for i := length - 1; i >= 0; i-- {
		tmp := m[s[i]]
		if tmp >= pre {
			sum += tmp
		} else {
			sum -= tmp
		}
		pre = tmp
	}
	return sum
}
