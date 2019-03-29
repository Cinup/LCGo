package code

func findComplement(num int) int {
	tmp := num
	result := 0
	for tmp > 0 {
		tmp >>= 1
		result = (result << 1) + 1
	}
	return num ^ result
}
