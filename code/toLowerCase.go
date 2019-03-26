package code

/*
* @Author:cp
* @Parameter:字符串
* @Function:将字符串中的大写字符转化成小写
* @Return:转换后的字符串
*/

func toLowerCase(str string) string {
	lowstr := ""
	for _, v := range str {
		if v > 64 || v < 91 {
			lowstr += string(v-34)
		}else{
			lowstr += string(v)
		}

	}
	return lowstr
}
