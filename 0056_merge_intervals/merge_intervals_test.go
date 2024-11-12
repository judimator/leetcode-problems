package merge_intervals

import (
	"reflect"
	"testing"
)

type TestData struct {
	input  [][]int
	output [][]int
}

var tests = []TestData{
	{
		[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		[][]int{{1, 6}, {8, 10}, {15, 18}},
	},
	{
		[][]int{{1, 4}, {4, 5}},
		[][]int{{1, 5}},
	},
	{
		[][]int{{1, 3}},
		[][]int{{1, 3}},
	},
	{
		[][]int{{1, 3}, {2, 6}, {8, 10}, {9, 18}},
		[][]int{{1, 6}, {8, 18}},
	},
	{
		[][]int{{1, 4}, {0, 4}},
		[][]int{{0, 4}},
	},
	{
		[][]int{{1, 4}, {0, 1}},
		[][]int{{0, 4}},
	},
	{
		[][]int{{1, 4}, {0, 0}},
		[][]int{{0, 0}, {1, 4}},
	},
}

func TestMerge(t *testing.T) {
	for _, data := range tests {
		if result := merge(data.input); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, merge")
		}
	}
}
