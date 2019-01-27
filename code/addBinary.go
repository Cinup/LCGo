package code

func addBinary(a string, b string) string {
	lengthA := len(a)
	lengthB := len(b)
	minLength := 0
	maxLength := 0
	var temp string
	if lengthA < lengthB {
		minLength = lengthA
		maxLength = lengthB
		temp = b[0 : maxLength-minLength]
	} else {
		maxLength = lengthA
		minLength = lengthB
		temp = a[0 : maxLength-minLength]
	}
	s := make([]uint8, maxLength+1)
	indexA := lengthA - 1
	indexB := lengthB - 1
	index := maxLength
	//ACSII 48 = char 0
	var carry uint8 = 48
	for i := minLength - 1; i >= 0; i-- {
		switch a[indexA] + b[indexB] + carry {
		case 144:
			s[index] = 48
			carry = 48
		case 145:
			s[index] = 49
			carry = 48
		case 146:
			s[index] = 48
			carry = 49
		case 147:
			s[index] = 49
			carry = 49
		}
		indexA--
		indexB--
		index--
	}
	i := maxLength - minLength - 1
	for ; i >= 0; i-- {
		switch temp[i] + carry {
		case 96:
			s[index] = 48
			carry = 48
		case 97:
			s[index] = 49
			carry = 48
		case 98:
			s[index] = 48
			carry = 49
		}
		index--
	}
	s[0] = carry
	if s[0] == 48 {
		return string(s[1:])
	}
	return string(s)
}
