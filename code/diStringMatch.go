package code

/*
* @Author:cp
* @Parameter:字符串
* @Function:...难以描述
* @Return:数组
*/
func diStringMatch(S string) []int {
	min := 0
	max := len(S)
	result := make([]int, 1000)
	for i := 0; i < len(S); i++ {
		if S[i] == 'I' {
			result[i] = min
			min++
		} else {
			result[i] = max
			max--
		}
	}
	result[len(S)] = max
	return  result[0:len(S)]
}
