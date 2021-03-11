package integer_to_roman

import "testing"

type TestData struct {
	input  int
	output string
}

var tests = []TestData{
	{
		3749,
		"MMMDCCXLIX",
	},
	{
		58,
		"LVIII",
	},
	{
		1994,
		"MCMXCIV",
	},
}

func TestIntegerToRoman(t *testing.T) {
	for _, data := range tests {
		if result := intToRoman(data.input); result != data.output {
			t.Error("Failed, isPalindrome")
		}
	}
}
