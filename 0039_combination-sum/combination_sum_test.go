package combination_sum

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
		[]int{2, 3, 6, 7},
		7,
		[][]int{{2, 2, 3}, {7}},
	},
}

func TestCombinationSum(t *testing.T) {
	for _, data := range tests {
		if result := combinationSum(data.input, data.target); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, combinationSum")
		}
	}
}
