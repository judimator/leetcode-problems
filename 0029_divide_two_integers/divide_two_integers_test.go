package divide_two_integers

import "testing"

type TestData struct {
	dividend int
	divisor  int
	output   int
}

var tests = []TestData{
	{
		10,
		3,
		3,
	},
	{
		7,
		-3,
		-2,
	},
	{
		-7,
		-3,
		2,
	},
	{
		-7,
		3,
		-2,
	},
}

func TestDivide(t *testing.T) {
	for _, data := range tests {
		if result := divide(data.dividend, data.divisor); result != data.output {
			t.Error("Failed, divide")
		}
	}
}
