package merge_two_sorted_lists

import (
	"reflect"
	"testing"
)

type TestData struct {
	input1 *ListNode
	input2 *ListNode
	output *ListNode
}

func TestMergeTwoLists(t *testing.T) {
	node4 := ListNode{4, nil}
	node2 := ListNode{2, &node4}
	node1 := ListNode{1, &node2}

	node44 := ListNode{4, nil}
	node33 := ListNode{3, &node44}
	node11 := ListNode{1, &node33}

	result4 := ListNode{4, nil}
	result44 := ListNode{4, &result4}
	result3 := ListNode{3, &result44}
	result2 := ListNode{2, &result3}
	result1 := ListNode{1, &result2}
	result11 := ListNode{1, &result1}

	var tests = []TestData{
		{
			&node1,
			&node11,
			&result11,
		},
	}

	for _, data := range tests {
		if result := mergeTwoLists(data.input1, data.input2); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, mergeTwoLists")
		}
	}
}
