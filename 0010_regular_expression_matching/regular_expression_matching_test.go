package regular_expression_matching

import "testing"

type TestData struct {
	s      string
	p      string
	output bool
}

var tests = []TestData{
	{
		"aa",
		"a",
		false,
	},
	{
		"aa",
		"a*",
		true,
	},
	{
		"ab",
		".*",
		true,
	},
	{
		"aaabbbbbccccddd",
		"aaab*cccc*ddd",
		true,
	},
}

func TestIsMatch(t *testing.T) {
	for _, data := range tests {
		if result := isMatch(data.s, data.p); result != data.output {
			t.Error("Failed, isMatch")
		}
	}
}
