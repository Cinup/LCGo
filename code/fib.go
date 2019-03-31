package code

func fib(N int) int {
	f1, f2 := 1, 1
	if N == 0 || N == 1 {
		return N
	}
	for i := 2; i <= N; i++ {
		tmp := f2
		f2 = f1 + f2
		f1 = tmp
	}
	return f2
}
