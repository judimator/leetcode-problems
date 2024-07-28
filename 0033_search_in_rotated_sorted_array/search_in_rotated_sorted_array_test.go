package search_in_rotated_sorted_array

import "testing"

type TestData struct {
	input  []int
	target int
	output int
}

var tests = []TestData{
	{
		[]int{4, 5, 6, 7, 0, 1, 2},
		0,
		4,
	},
	{
		[]int{4, 5, 6, 7, 0, 1, 2},
		3,
		-1,
	},
}

func TestSearch(t *testing.T) {
	for _, data := range tests {
		if result := search(data.input, data.target); result != data.output {
			t.Error("Failed, search")
		}
	}
}
