package code

func hammingDistance(x int, y int) int {
	distance := 0
	for x > 0 && y > 0 {
		if x%2 != y%2 {
			distance++
		}
		x = x >> 1
		y = y >> 1
	}
	return distance
}
func distance(x int) int {
	dis := 0
	for x > 0 {
		x = x >> 1
		dis++
	}
	return dis
}
