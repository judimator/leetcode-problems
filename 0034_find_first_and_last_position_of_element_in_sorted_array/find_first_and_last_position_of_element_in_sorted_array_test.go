package find_first_and_last_position_of_element_in_sorted_array

import (
	"reflect"
	"testing"
)

type TestData struct {
	input  []int
	target int
	output []int
}

var tests = []TestData{
	{
		[]int{5, 7, 7, 8, 8, 10},
		8,
		[]int{3, 4},
	},
	{
		[]int{2, 2},
		2,
		[]int{0, 1},
	},
}

func TestSearchRange(t *testing.T) {
	for _, data := range tests {
		if result := searchRange(data.input, data.target); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, searchRange")
		}
	}
}
