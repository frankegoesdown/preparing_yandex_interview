package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	data1 := []int{1, 2, 4}
	data2 := []int{1, 3, 4}
	fmt.Println(mergeTwoLists(MakeListNote(data1), MakeListNote(data2)))
}

func MakeListNote(x []int) *ListNode {
	list := &ListNode{}
	head := list

	list.Val = x[0]
	for i := 1; i < len(x); i++ {
		list.Next = &ListNode{}
		list = list.Next
		list.Val = x[i]
	}
	return head
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tmp := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	if l1 != nil {
		tmp.Next = l1
	} else {
		tmp.Next = l2
	}

	return head.Next
}
