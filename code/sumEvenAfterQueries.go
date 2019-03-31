package code

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	result := make([]int, 10000)
	sum := 0
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			sum += A[i]
		}
	}
	for i := 0; i < len(queries); i++ {
		val := queries[i][0]
		idx := queries[i][1]

		if A[idx]%2 != 0 {
			//奇+奇
			if val%2 != 0 {
				sum = sum + A[idx] + val
			}
		} else {
			//偶+偶
			if val%2 == 0 {
				sum += val
			} else { //偶+奇
				sum -= A[idx]
			}
		}
		A[idx] += val
		result[i] = sum
	}
	return result[:len(queries)]
}
