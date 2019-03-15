package code

/*
* @Author:cp
* @Parameter:两个链表头节点
* @Function:找出两个链表相交的节点
* @Return:相交节点
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	lengthA := getLength(headA)
	lengthB := getLength(headB)
	if lengthA > lengthB {
		for i := 0; i < lengthA-lengthB; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < lengthB-lengthA; i++ {
			headB = headA.Next
		}
	}
	for headA != nil {
		if headA == headB {
			break
		}else {
			headA = headA.Next
			headB = headB.Next
		}
	}
	return headA
}
func getLength(node *ListNode) int {
	length := 0
	for node != nil {
		node = node.Next
		length++
	}
	return length
}
