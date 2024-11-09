package maximum_subarray

import (
	"testing"
)

type TestData struct {
	input  []int
	output int
}

var tests = []TestData{
	{
		[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		6,
	},
}

func TestMaxSubArray(t *testing.T) {
	for _, data := range tests {
		if result := maxSubArray(data.input); result != data.output {
			t.Error("Failed, maxSubArray")
		}
	}
}
