package code

import (
	"strings"
)

func shortestToChar(S string, C byte) []int {
	var distance = make([]int, 10000)
	firstIndex := strings.Index(S, string(C))
	for i := 0; i <= firstIndex; i++ {
		distance[i] = firstIndex - i
	}
	dis := 0
	for i := firstIndex + 1; i < len(S); i++ {
		if S[i] == C {
			dis = 0
		} else {
			dis++
		}
		distance[i] = dis

	}
	lastIndex := strings.LastIndex(S, string(C))
	for i := len(S) - 1; i >= lastIndex; i-- {
		distance[i] = len(S) - 1 - i
	}
	dis = 0
	for i := lastIndex - 1; i >= 0; i-- {
		if S[i] == C {
			dis = 0
		} else {
			dis++
		}
		if dis < distance[i] {
			distance[i] = dis
		}

	}
	//fmt.Println(strings.Index(S,string(C)))
	return distance[:len(S)+1]
}
