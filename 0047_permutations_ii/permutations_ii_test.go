package permutations_ii

import (
	"reflect"
	"testing"
)

type TestData struct {
	nums   []int
	output [][]int
}

var tests = []TestData{
	{
		[]int{1, 1, 2},
		[][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
	},
}

func TestPermute(t *testing.T) {
	for _, data := range tests {
		if result := permuteUnique(data.nums); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, permuteUnique")
		}
	}
}
