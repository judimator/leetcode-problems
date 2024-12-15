package rotate_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	oldTail, l := head, 1

	for oldTail.Next != nil {
		oldTail = oldTail.Next
		l++
	}

	k = k % l
	if k == 0 {
		return head
	}

	tail := head

	for i := 0; i < l-k-1; i++ {
		tail = tail.Next
	}

	newHead := tail.Next
	tail.Next = nil
	oldTail.Next = head

	return newHead
}
