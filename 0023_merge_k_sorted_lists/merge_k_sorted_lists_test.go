package merge_k_sorted_lists

import (
	"reflect"
	"testing"
)

type TestData struct {
	input1 []*ListNode
	output *ListNode
}

func TestMergeKSortedLists(t *testing.T) {
	node5 := ListNode{5, nil}
	node4 := ListNode{4, &node5}
	node1 := ListNode{1, &node4}

	node44 := ListNode{4, nil}
	node33 := ListNode{3, &node44}
	node11 := ListNode{1, &node33}

	node666 := ListNode{6, nil}
	node222 := ListNode{2, &node666}

	result6 := ListNode{6, nil}
	result5 := ListNode{5, &result6}
	result4 := ListNode{4, &result5}
	result44 := ListNode{4, &result4}
	result3 := ListNode{3, &result44}
	result2 := ListNode{2, &result3}
	result1 := ListNode{1, &result2}
	result11 := ListNode{1, &result1}

	var tests = []TestData{
		{
			[]*ListNode{&node1, &node11, &node222},
			&result11,
		},
	}

	for _, data := range tests {
		if result := mergeKLists(data.input1); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, mergeKLists")
		}
	}
}
