package reverse_integer

import (
	"testing"
)

type TestData struct {
	input  int
	output int
}

var tests = []TestData{
	{
		123,
		321,
	},
	{
		-123,
		-321,
	},
	{
		120,
		21,
	},
	{
		1534236469,
		0,
	},
}

func TestReverse(t *testing.T) {
	for _, data := range tests {
		if result := reverse(data.input); result != data.output {
			t.Error("Failed, reverse")
		}
	}
}
