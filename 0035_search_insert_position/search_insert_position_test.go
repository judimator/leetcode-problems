package search_insert_position

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
		[]int{1, 3, 5, 6},
		5,
		2,
	},
	{
		[]int{1, 3, 5, 6},
		2,
		1,
	},
	{
		[]int{1, 3, 5, 6},
		7,
		4,
	},
}

func TestSearchInsert(t *testing.T) {
	for _, data := range tests {
		if result := searchInsert(data.input, data.target); result != data.output {
			t.Error("Failed, searchInsert")
		}
	}
}
