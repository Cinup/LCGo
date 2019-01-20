package code

func reverse(x int) int {
	isNegative := false
	max := 1<<31 - 1
	min := -max - 1
	//result := make([]int, 32)
	if x < 0 {
		x = -x
		isNegative = true
	}
	num := 0
	sum := 0
	for x > 0 {
		sum = sum*10 + x%10
		x = x / 10
		num++
		if num > 32 {
			return 0
		}
	}
	if isNegative {
		sum = -sum
	}
	if sum > max || sum < min {
		return 0
	}
	return sum
}
