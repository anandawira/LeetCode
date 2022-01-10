package shopee

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := new(ListNode)
	cur := res

	for list1 != nil || list2 != nil {
		if list1 == nil {
			cur.Next = list2
			list2 = nil
			continue
		}

		if list2 == nil {
			cur.Next = list1
			list1 = nil
			continue
		}

		if list1.Val > list2.Val {
			cur.Next = list2
			cur, list2 = cur.Next, list2.Next
		} else {
			cur.Next = list1
			cur, list1 = cur.Next, list1.Next
		}
	}

	return res.Next
}
