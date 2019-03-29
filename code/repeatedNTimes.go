package code

func repeatedNTimes(A []int) int {
	i, length := 0, len(A)
	for {
		//n个数字有n+1个空位,连成环的话就只有n个空位,因此两个步长内必有相同的数
		if A[i] == A[(i+1)%length] || A[i] == A[(i+2)%length] {
			break
		}
		i++
	}
	return A[i]
}
