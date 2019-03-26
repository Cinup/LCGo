package code

import "sort"

/*
* @Author:cp
* @Parameter:数组
* @Function:将数组中数字的平方排序
* @Return:平方后的排序数组
*/
func sortedSquares(A []int) []int {
	for i := 0; i < len(A); i++ {
		A[i] *= A[i]
	}
	sort.Ints(A)
	return A
}
