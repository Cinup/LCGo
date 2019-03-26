package code

/*
* @Author:cp
* @Parameter:机器人移动指令的字符串
* @Function:根据字符串判断机器人是否回到原地
* @Return:判断结果
*/
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
