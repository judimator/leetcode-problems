package remove_duplicates_from_sorted_array

import "testing"

type TestData struct {
	input  []int
	output int
}

var tests = []TestData{
	{
		[]int{1, 1, 2},
		2,
	},
	{
		[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		5,
	},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, data := range tests {
		if result := removeDuplicates(data.input); result != data.output {
			t.Error("Failed, removeDuplicates")
		}
	}
}
