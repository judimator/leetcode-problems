package insert_interval

import (
	"reflect"
	"testing"
)

type TestData struct {
	intervals   [][]int
	newInterval []int
	output      [][]int
}

var tests = []TestData{
	{
		[][]int{{1, 3}, {4, 8}, {9, 12}},
		[]int{2, 10},
		[][]int{{1, 12}},
	},
	{
		[][]int{{1, 3}, {6, 9}},
		[]int{2, 5},
		[][]int{{1, 5}, {6, 9}},
	},
	{
		[][]int{{2, 3}, {4, 8}, {9, 12}},
		[]int{0, 1},
		[][]int{{0, 1}, {2, 3}, {4, 8}, {9, 12}},
	},
	{
		[][]int{{2, 3}, {4, 8}, {9, 12}},
		[]int{13, 15},
		[][]int{{2, 3}, {4, 8}, {9, 12}, {13, 15}},
	},
	{
		[][]int{{2, 3}, {4, 8}, {9, 12}},
		[]int{1, 15},
		[][]int{{1, 15}},
	},
	{
		[][]int{{2, 3}, {4, 8}},
		[]int{2, 3},
		[][]int{{2, 3}, {4, 8}},
	},
	{
		[][]int{{1, 5}},
		[]int{2, 3},
		[][]int{{1, 5}},
	},
	{
		[][]int{{1, 5}},
		[]int{2, 7},
		[][]int{{1, 7}},
	},
}

func TestMerge(t *testing.T) {
	for _, data := range tests {
		if result := insert(data.intervals, data.newInterval); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, insert")
		}
	}
}
