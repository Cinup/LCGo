package code

func selfDividingNumbers(left int, right int) []int {
	var result []int
	for i := left; i <= right; i++ {
		tmp := i
		for tmp > 0 {
			c := tmp % 10
			if c == 0 || i%c != 0 {
				break
			}
			tmp = tmp / 10
		}
		if tmp == 0 {
			result = append(result, i)
		}
	}
	return result
}
