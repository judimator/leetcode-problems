package longest_valid_parentheses

import (
	"testing"
)

type TestData struct {
	s      string
	output int
}

var tests = []TestData{
	{
		"(()",
		2,
	},
	{
		")()())",
		4,
	},
}

func TestFindSubstring(t *testing.T) {
	for _, data := range tests {
		if result := longestValidParentheses(data.s); result != data.output {
			t.Error("Failed, longestValidParentheses")
		}
	}
}
