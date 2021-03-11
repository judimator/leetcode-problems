package threesum

import (
	"reflect"
	"testing"
)

type TestData struct {
	input  []int
	output [][]int
}

var tests = []TestData{
	{
		[]int{-1, 0, 1, 2, -1, -4},
		[][]int{{-1, -1, 2}, {-1, 0, 1}},
	},
	{
		[]int{0, 0, 0, 0},
		[][]int{{0, 0, 0}},
	},
	{
		[]int{1, -1, -1, 0},
		[][]int{{-1, 0, 1}},
	},
}

func TestThreeSum(t *testing.T) {
	for _, data := range tests {
		if result := threeSum(data.input); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, threeSum")
		}
	}
}
