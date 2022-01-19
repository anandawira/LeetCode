package random

func isLengthEven(head *ListNode) bool {
	isEven := true

	for head != nil {
		isEven = !isEven
		head = head.Next
	}

	return isEven
}
