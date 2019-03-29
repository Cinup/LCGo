package code

func sortArrayByParityII(A []int) []int {
	var array = make([]int, 20000)
	arrLength := len(A)
	i, j := 0, 1
	for k := 0; k < arrLength; k++ {
		if A[k]%2 == 0 {
			array[i] = A[k]
			i += 2
		} else {
			array[j] = A[k]
			j += 2
		}
	}
	return array[0:arrLength]
}
