package code

/*
* @Author:cp
* @Parameter:一个字符串数组,数组的每个元素都是一个邮件地址
* @Function:计算实际的邮件地址
* @Return:邮件地址数
*/
func numUniqueEmails(emails []string) int {

	var email = make(map[string]int)

	for i := 0; i < len(emails); i++ {
		s := emails[i]
		for j := 0; j < len(s); j++ {
			switch s[i] {
			case '.':
			case '+':
			case '@':

			}
		}
		email[s[0:]] = 1

	}
}
