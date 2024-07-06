package reverse_nodes_in_k_group

import (
	"reflect"
	"testing"
)

type TestData struct {
	input1 *ListNode
	k      int
	output *ListNode
}

func TestReverseKGroup1(t *testing.T) {
	node5 := ListNode{5, nil}
	node4 := ListNode{4, &node5}
	node3 := ListNode{3, &node4}
	node2 := ListNode{2, &node3}
	node1 := ListNode{1, &node2}

	result5 := ListNode{5, nil}
	result3 := ListNode{3, &result5}
	result4 := ListNode{4, &result3}
	result1 := ListNode{1, &result4}
	result2 := ListNode{2, &result1}

	var tests = []TestData{
		{
			&node1,
			2,
			&result2,
		},
	}

	for _, data := range tests {
		if result := reverseKGroup(data.input1, data.k); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, reverseKGroup")
		}
	}
}

func TestReverseKGroup2(t *testing.T) {
	node5 := ListNode{5, nil}
	node4 := ListNode{4, &node5}
	node3 := ListNode{3, &node4}
	node2 := ListNode{2, &node3}
	node1 := ListNode{1, &node2}

	result5 := ListNode{5, nil}
	result4 := ListNode{4, &result5}
	result1 := ListNode{1, &result4}
	result2 := ListNode{2, &result1}
	result3 := ListNode{3, &result2}

	var tests = []TestData{
		{
			&node1,
			3,
			&result3,
		},
	}

	for _, data := range tests {
		if result := reverseKGroup(data.input1, data.k); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, reverseKGroup")
		}
	}
}

func TestReverseKGroup3(t *testing.T) {
	node2 := ListNode{2, nil}
	node1 := ListNode{1, &node2}

	result1 := ListNode{1, nil}
	result2 := ListNode{2, &result1}

	var tests = []TestData{
		{
			&node1,
			2,
			&result2,
		},
	}

	for _, data := range tests {
		if result := reverseKGroup(data.input1, data.k); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, reverseKGroup")
		}
	}
}
