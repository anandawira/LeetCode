package daily

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	l := 0

	for p := head; p != nil; p = p.Next {
		l++
	}

	n, r := l/k, l%k

	ans := make([]*ListNode, k)

	p := head

	for i := 0; i < k && p != nil; i++ {
		ans[i] = p

		var rc int
		if r > 0 {
			rc = 1
			r--
		}

		for j := 1; j < n+rc; j++ {
			p = p.Next
		}

		t := p.Next
		p.Next = nil
		p = t
	}

	return ans
}
