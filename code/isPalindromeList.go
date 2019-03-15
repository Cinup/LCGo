package code

/*
* @Author:cp
* @Parameter:链表头节点
* @Function:判断链表是否为回文链表
* @Return:判断结果
*/
func isPalindromeList(head *ListNode) bool {
	if head == nil {
		return true
	}
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	i, j := 0, len(result)-1
	for i < j {
		if result[i] != result[j] {
			return  false
		}
		i++
		j--
	}
	return true
}
