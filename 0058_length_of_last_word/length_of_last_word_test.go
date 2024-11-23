package length_of_last_word

import (
	"testing"
)

type TestData struct {
	string string
	output int
}

var tests = []TestData{
	{
		"Hello World",
		5,
	},
	{
		"   fly me   to   the moon  ",
		4,
	},
	{
		"   ",
		0,
	},
	{
		"Hello",
		5,
	},
}

func TestLengthOfLastWord(t *testing.T) {
	for _, data := range tests {
		if result := lengthOfLastWord(data.string); result != data.output {
			t.Error("Failed, lengthOfLastWord")
		}
	}
}
