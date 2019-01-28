package code

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	} else {
		n1 := 1
		n2 := 2
		res := 0
		for i := 3; i <= n; i++ {
			res = n1 + n2
			n1 = n2
			n2 = res
		}
		return res
	}
}
