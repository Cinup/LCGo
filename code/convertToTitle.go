package code

/*
* @Author:cp
* @Parameter:一个整数
* @Function:给定一个正整数,返回它在Excel表中相对应的列名称
* @Return:列名称
*/
func convertToTitle(n int) string {
	//s := ""
	var b []byte
	for n > 26 {
		b = append(b, byte(65+n%26))
		n = n / 26
	}
	b = append(b,byte(n))
	return string(b)
}
