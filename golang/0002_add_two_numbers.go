package main

func main() {

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	p := &ListNode{}
	res := p
	for l1 != nil || l2 != nil {
		s := carry
		if l1 != nil {
			s += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			s += l2.Val
			l2 = l2.Next
		}

		res.Next = &ListNode{Val: s % 10}

		carry = s / 10
		res = res.Next
	}

	if carry > 0 {
		res.Next = &ListNode{Val: carry}
	}

	return p.Next
}
