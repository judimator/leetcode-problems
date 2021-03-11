package threesum_closest

import (
	"testing"
)

type TestData struct {
	input  []int
	target int
	output int
}

var tests = []TestData{
	{
		[]int{-1, 2, 1, -4},
		1,
		2,
	},
}

func TestThreeSumClosest(t *testing.T) {
	for _, data := range tests {
		if result := threeSumClosest(data.input, data.target); result != data.output {
			t.Error("Failed, threeSumClosest")
		}
	}
}
