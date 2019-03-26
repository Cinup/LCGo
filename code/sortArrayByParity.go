package code

/*
* @Author:cp
* @Parameter:非负整数数组
* @Function:将数组中的奇数和偶数分开
* @Return:数组
*/
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
