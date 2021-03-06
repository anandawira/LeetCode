package random

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) (prev *ListNode) {
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return
}
