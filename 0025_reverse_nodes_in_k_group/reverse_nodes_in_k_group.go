package reverse_nodes_in_k_group

type ListNode struct {
	Val  int
	Next *ListNode
}

type Stack struct {
	nodes []*ListNode
}

func (s *Stack) push(node *ListNode) {
	s.nodes = append(s.nodes, node)
}

func (s *Stack) len() int {
	return len(s.nodes)
}

func (s *Stack) pop() *ListNode {
	node := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[0 : len(s.nodes)-1]

	return node
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	node := &ListNode{
		-1,
		nil,
	}
	tmp := node
	s := Stack{}

	for head != nil {
		if s.len() == k {
			for s.len() > 0 {
				tmp.Next = s.pop()
				tmp.Next.Next = nil
				tmp = tmp.Next
			}
		}

		s.push(head)
		head = head.Next
	}

	var last *ListNode

	if s.len() == k {
		for s.len() > 0 {
			tmp.Next = s.pop()
			tmp.Next.Next = nil
			tmp = tmp.Next
		}
	} else {
		for s.len() > 0 {
			last = s.pop()
		}
	}

	tmp.Next = last

	return node.Next
}
