package code

func isValid(s string) bool {
	//极端情况,s为空
	if s == "" {
		return true
	}
	m := make(map[uint8]uint8)
	m['('] = ')'
	m['['] = ']'
	m['{'] = m['}']
	end := len(s) - 1
	//直接先判断首尾是否有效
	if s[0] == ')' || s[0] == ']' || s[0] == '}' {
		return false
	} else if s[end] == '(' || s[end] == '[' || s[end] == '{' {
		return false
	}
	var result []uint8
	result[0] = s[0]
	next := m[s[0]]
	//
	for i := 1; i <= end; i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			result = append(result, s[i])
			next = m[s[i]]
		} else if s[i] == next {
			result = result[:len(result)-1]
		} else {
			return false
		}
		if len(result) == 0 {
			next = ' '
		}else {
			next = result[len(result)-1]
		}
	}
	if len(result) == 0 {
		return true

	}
	return  false
}
