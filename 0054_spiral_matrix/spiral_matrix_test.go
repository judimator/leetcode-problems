package spiral_matrix

import (
	"reflect"
	"testing"
)

type TestData struct {
	input  [][]int
	output []int
}

var tests = []TestData{
	{
		[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
		[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
	},
	{
		[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
	},
	{
		[][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}},
		[]int{1, 2, 3, 4, 5, 10, 9, 8, 7, 6},
	},
	{
		[][]int{{3}, {2}},
		[]int{3, 2},
	},
	{
		[][]int{{2, 3}},
		[]int{2, 3},
	},
	{
		[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}, {17, 18, 19, 20}, {21, 22, 23, 24}},
		[]int{1, 2, 3, 4, 8, 12, 16, 20, 24, 23, 22, 21, 17, 13, 9, 5, 6, 7, 11, 15, 19, 18, 14, 10},
	},
}

func TestSpiralOrder(t *testing.T) {
	for _, data := range tests {
		if result := spiralOrder(data.input); !reflect.DeepEqual(data.output, result) {
			t.Error("Failed, spiralOrder")
		}
	}
}
