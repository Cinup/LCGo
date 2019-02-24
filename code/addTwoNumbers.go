package code

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l := &ListNode{}
	index := l
	//进位
	carry := 0
	for l1 != nil && l2 != nil {
		tmp := new(ListNode)
		tmp.Val = (l1.Val + l2.Val + carry) % 10
		carry = (l1.Val + l2.Val + carry) / 10
		//
		l1 = l1.Next
		l2 = l2.Next
		//
		index.Next = tmp
		index = index.Next
	}
	//
	if l1 == nil {
		addRemainder(index, l2, carry)
	} else {
		addRemainder(index, l1, carry)
	}
	return l.Next
}
func addRemainder(l, r *ListNode, carry int) {
	for r != nil {
		tmp := new(ListNode)
		tmp.Val = (r.Val + carry) % 10
		carry = (r.Val + carry) / 10
		//
		r = r.Next
		//
		l.Next = tmp
		l = l.Next
	}
	//加上最后一个进位
	if carry == 1 {
		tmp := new(ListNode)
		tmp.Val = 1
		l.Next = tmp
	}
}
