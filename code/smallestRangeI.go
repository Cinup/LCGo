package code

import "sort"

func smallestRangeI(A []int, K int) int {
	if len(A) < 2 {
		return 0
	}
	sort.Ints(A)
	min, max := A[0], A[len(A)-1]
	if (max - min) > 2*K {
		return max - min - 2*K
	}
	return 0
}
