package code

func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		for j, k := 0, len(A[i])-1; j <= k; {
			A[i][j], A[i][k] = A[i][k]^1, A[i][j]^1
			j++
			k--
		}
	}
	return A
}
