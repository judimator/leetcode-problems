package container_with_most_water

import "testing"

type TestData struct {
	input  []int
	output int
}

var tests = []TestData{
	{
		[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		49,
	},
	{
		[]int{1, 1},
		1,
	},
}

func TestMaxArea(t *testing.T) {
	for _, data := range tests {
		if result := maxArea(data.input); result != data.output {
			t.Error("Failed, maxArea")
		}
	}
}
