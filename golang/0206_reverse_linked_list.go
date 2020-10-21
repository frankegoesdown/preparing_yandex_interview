package main

func main() {

}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		prev, cur, cur.Next = cur, cur.Next, prev
	}

	return prev
}
