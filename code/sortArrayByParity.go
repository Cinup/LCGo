package code

func sortArrayByParity(A []int) []int {
	var array = make([]int, 5000)
	arrLength := len(A)
	i, j := 0, arrLength-1
	//偶数在前,奇数在后
	for k := 0; k < arrLength; k++ {
		if A[k]%2 == 0 {
			array[i] = A[k]
			i++
		} else {
			array[j] = A[k]
			j--
		}
	}
	return array[:arrLength]
}
