package permutations

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
		[]int{1, 2, 3},
		[][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
	},
}

func TestPermute(t *testing.T) {
	for _, data := range tests {
		if result := permute(data.nums); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, permute")
		}
	}
}
