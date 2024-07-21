package next_permutation

import (
	"reflect"
	"testing"
)

type TestData struct {
	input  []int
	output []int
}

var tests = []TestData{
	{
		[]int{1, 2, 3},
		[]int{1, 3, 2},
	},
	{
		[]int{1, 1, 5},
		[]int{1, 5, 1},
	},
	{
		[]int{3, 2, 1},
		[]int{1, 2, 3},
	},
	{
		[]int{2, 3, 1},
		[]int{3, 2, 1},
	},
}

func TestNextPermutation(t *testing.T) {
	for _, data := range tests {
		nextPermutation(data.input)

		if !reflect.DeepEqual(data.input, data.output) {
			t.Error("Failed, nextPermutation")
		}
	}
}
