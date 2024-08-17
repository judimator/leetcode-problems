package count_and_say

import "testing"

type TestData struct {
	input  int
	output string
}

var tests = []TestData{
	{
		1,
		"1",
	},
	{
		2,
		"11",
	},
	{
		3,
		"21",
	},
	{
		4,
		"1211",
	},
	{
		5,
		"111221",
	},
}

func TestCountAndSay(t *testing.T) {
	for _, data := range tests {
		if result := countAndSay(data.input); result != data.output {
			t.Error("Failed, countAndSay")
		}
	}
}
