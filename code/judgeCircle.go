package code

func judgeCircle(moves string) bool {
	//机器人坐标轴的位置
	x, y := 0, 0
	for _, v := range moves {
		switch v {
		case 'U':
			y++
		case 'D':
			y--
		case 'L':
			x--
		case 'R':
			x++
		}
	}
	if x == 0 && y == 0 {
		return true
	}
	return false
}
