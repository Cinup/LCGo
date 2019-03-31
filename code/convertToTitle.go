package code

func convertToTitle(n int) string {
	//s := ""
	var b []byte
	for n > 26 {
		b = append(b, byte(65+n%26))
		n = n / 26
	}
	b = append(b, byte(n))
	return string(b)
}
