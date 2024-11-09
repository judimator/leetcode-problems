package swap_nodes_in_pairs

import (
	"reflect"
	"testing"
)

type TestData struct {
	input1 *ListNode
	output *ListNode
}

func TestMergeTwoLists(t *testing.T) {
	node4 := ListNode{4, nil}
	node3 := ListNode{3, &node4}
	node2 := ListNode{2, &node3}
	node1 := ListNode{1, &node2}

	result3 := ListNode{3, nil}
	result4 := ListNode{4, &result3}
	result1 := ListNode{1, &result4}
	result2 := ListNode{2, &result1}

	var tests = []TestData{
		{
			&node1,
			&result2,
		},
	}

	for _, data := range tests {
		if result := swapPairs(data.input1); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, swapPairs")
		}
	}
}
