package code

func plusOne(digits []int) []int {
	if digits[len(digits)-1] != 9 {
		digits[len(digits)-1] += 1
		return digits
	}
	carry := 1
	for i := len(digits) - 1; i > 0; i-- {
		digits[i] += carry
		if digits[i] == 10 {
			digits[i] = 0
			carry = 1
		} else {
			carry = 0
			break
		}
	}
	digits[0] += carry
	if digits[0] == 10 {
		digits[0] = 0
		s := [] int {1}
		digits = append(s,digits...)
	}
	return digits
}
