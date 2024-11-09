package swap_nodes_in_pairs

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

func (s *Stack) pop() *ListNode {
	item := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[0 : len(s.nodes)-1]

	return item
}

func (s *Stack) len() int {
	return len(s.nodes)
}

func swapPairs(head *ListNode) *ListNode {
	node := &ListNode{
		-1,
		nil,
	}
	tmp := node
	s := Stack{}

	for head != nil {
		if s.len() == 2 {
			for s.len() > 0 {
				tmp.Next = s.pop()
				tmp.Next.Next = nil
				tmp = tmp.Next
			}
		}

		s.push(head)
		head = head.Next
	}

	for s.len() > 0 {
		tmp.Next = s.pop()
		tmp.Next.Next = nil
		tmp = tmp.Next
	}

	return node.Next
}
