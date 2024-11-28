package spiral_matrix_ii

import (
	"reflect"
	"testing"
)

type TestData struct {
	n      int
	output [][]int
}

var tests = []TestData{
	{
		3,
		[][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
	},
	{
		4,
		[][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}},
	},
}

func TestGenerateMatrix(t *testing.T) {
	for _, data := range tests {
		if result := generateMatrix(data.n); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, generateMatrix")
		}
	}
}
