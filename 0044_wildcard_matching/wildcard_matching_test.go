package wildcard_matching

import (
	"testing"
)

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
		"*",
		true,
	},
	{
		"bbbbbbbabaabaaabaaaaaabbbabbabbaaabbbabaabbababbabaaabbbbaaabbabbbaabbabbabaaaaaaaaabbbabbabaaabbbbaabbbbaabbabbabbbaabababbabaaababaaaaaabbabaaabbbbbbababbbbbaaabbabbaaaaababaaabbbaaaababb",
		"**b*bb*aa*baaaa*aaa*b*baaa*a*aaa*b*a*ba**ba*ba*ba*b*b****a*ba*b**a*****ba*bb*a***abb***a*bb***b**abb*",
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
