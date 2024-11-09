package multiply_strings

import (
	"testing"
)

type TestData struct {
	num1   string
	num2   string
	output string
}

var tests = []TestData{
	{
		"123",
		"456",
		"56088",
	},
}

func TestTrap(t *testing.T) {
	for _, data := range tests {
		if result := multiply(data.num1, data.num2); result != data.output {
			t.Error("Failed, multiply")
		}
	}
}
