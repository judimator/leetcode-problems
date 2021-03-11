package longest_common_prefix

import "testing"

type TestData struct {
	input  []string
	output string
}

var tests = []TestData{
	{
		[]string{"testing", "test"},
		"test",
	},
	{
		[]string{"some string", "sorrow"},
		"so",
	},
	{
		[]string{"some string", "another"},
		"",
	},
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, data := range tests {
		if result := longestCommonPrefix(data.input); result != data.output {
			t.Error("Failed, longestCommonPrefix")
		}
	}
}
