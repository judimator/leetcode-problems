package roman_ro_integer

import "testing"

type TestData struct {
	input  string
	output int
}

var tests = []TestData{
	{
		"III",
		3,
	},
	{
		"LVIII",
		58,
	},
	{
		"MCMXCIV",
		1994,
	},
}

func TestRomanToInteger(t *testing.T) {
	for _, data := range tests {
		if result := romanToInt(data.input); result != data.output {
			t.Error("Failed, romanToInt")
		}
	}
}
