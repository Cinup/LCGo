package code

func minDeletionSize(A []string) int {
	length := 0
	for j := 0; j < len(A[0]); j++ {
		for i := 1; i < len(A); i++ {
			if A[i][j] < A[i-1][j] {
				length++
				break
			}
		}
	}
	return length
}
