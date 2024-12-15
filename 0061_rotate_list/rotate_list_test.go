package rotate_list

import (
	"reflect"
	"testing"
)

type TestData struct {
	input  *ListNode
	k      int
	output *ListNode
}

func TestRotateRight1(t *testing.T) {
	node5 := ListNode{5, nil}
	node4 := ListNode{4, &node5}
	node3 := ListNode{3, &node4}
	node2 := ListNode{2, &node3}
	node1 := ListNode{1, &node2}

	node33 := ListNode{3, nil}
	node22 := ListNode{2, &node33}
	node11 := ListNode{1, &node22}
	node55 := ListNode{5, &node11}
	node44 := ListNode{4, &node55}

	var tests = []TestData{
		{
			&node1,
			2,
			&node44,
		},
	}

	for _, data := range tests {
		if result := rotateRight(data.input, data.k); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, rotateRight1")
		}
	}
}
