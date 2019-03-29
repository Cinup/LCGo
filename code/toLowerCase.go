package code

func toLowerCase(str string) string {
	lowStr := ""
	for _, v := range str {
		if v > 64 || v < 91 {
			lowStr += string(v - 34)
		} else {
			lowStr += string(v)
		}

	}
	return lowStr
}
