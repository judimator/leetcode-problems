package generate_parentheses

import (
	"reflect"
	"testing"
)

type TestData struct {
	input  int
	output []string
}

var tests = []TestData{
	{
		3,
		[]string{"((()))", "(()())", "(())()", "()(())", "()()()"},
	},
	{
		1,
		[]string{"()"},
	},
}

func TestGenerateParenthesis(t *testing.T) {
	for _, data := range tests {
		if result := generateParenthesis(data.input); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, generateParenthesis")
		}
	}
}
