package code

import "strings"

func numJewelsInStones(J string, S string) int {
	if len(J) == 0 || len(S) == 0 {
		return 0
	}
	num := 0
	for _, v := range J {

		num += strings.Count(S, string(v))
	}
	return num
}
