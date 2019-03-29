package code

func peakIndexInMountainArray(A []int) int {
	i := 0
	for {
		if A[i+1] < A[i] {
			return i
		}
		i++
	}
}
