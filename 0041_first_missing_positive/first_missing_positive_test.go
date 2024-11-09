package first_missing_positive

import (
	"testing"
)

type TestData struct {
	input  []int
	output int
}

var tests = []TestData{
	{
		[]int{1, 2, 0},
		3,
	},
	{
		[]int{3, 4, -1, 1},
		2,
	},
	{
		[]int{7, 8, 9, 11, 12},
		1,
	},
}

func TestFirstMissingPositive(t *testing.T) {
	for _, data := range tests {
		if result := firstMissingPositive(data.input); result != data.output {
			t.Error("Failed, firstMissingPositive")
		}
	}
}
