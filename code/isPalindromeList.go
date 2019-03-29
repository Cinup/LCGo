package code

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
			return false
		}
		i++
		j--
	}
	return true
}
