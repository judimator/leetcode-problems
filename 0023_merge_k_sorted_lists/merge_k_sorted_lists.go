package merge_k_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		merged := merge(lists[0], lists[1])
		lists = append(lists[2:], merged)
	}

	return lists[0]
}

func merge(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{
		Val:  -1,
		Next: nil,
	}

	tmp := head

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tmp.Next = list1
			list1 = list1.Next
		} else {
			tmp.Next = list2
			list2 = list2.Next
		}

		tmp = tmp.Next
	}

	if list1 == nil {
		tmp.Next = list2
	} else {
		tmp.Next = list1
	}

	return head.Next
}
