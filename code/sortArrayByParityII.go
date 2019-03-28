package code

/*
* @Author:cp
* @Parameter:数组
* @Function:将数组中的奇偶排序,奇数的下标也是奇数
* @Return:数组
*/
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
