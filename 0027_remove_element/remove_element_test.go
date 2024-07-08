package remove_element

import "testing"

type TestData struct {
	input  []int
	val    int
	output int
}

var tests = []TestData{
	{
		[]int{3, 2, 2, 3},
		3,
		2,
	},
	{
		[]int{0, 1, 2, 2, 3, 0, 4, 2},
		2,
		5,
	},
}

func TestRemoveElement(t *testing.T) {
	for _, data := range tests {
		if result := removeElement(data.input, data.val); result != data.output {
			t.Error("Failed, removeElement")
		}
	}
}
