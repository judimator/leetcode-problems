package combination_sum_ii

import (
	"reflect"
	"testing"
)

type TestData struct {
	input  []int
	target int
	output [][]int
}

var tests = []TestData{
	{
		[]int{10, 1, 2, 7, 6, 1, 5},
		8,
		[][]int{{1, 1, 6}, {1, 2, 5}},
	},
}

func TestCombinationSum2(t *testing.T) {
	for _, data := range tests {
		if result := combinationSum2(data.input, data.target); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, combinationSum2")
		}
	}
}
