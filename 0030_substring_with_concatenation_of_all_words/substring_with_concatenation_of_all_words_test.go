package substring_with_concatenation_of_all_words

import (
	"reflect"
	"testing"
)

type TestData struct {
	s      string
	words  []string
	output []int
}

var tests = []TestData{
	{
		"barfoothefoobarman",
		[]string{"foo", "bar"},
		[]int{0, 9},
	},
	{
		"wordgoodgoodgoodbestword",
		[]string{"word", "good", "best", "word"},
		[]int{},
	},
}

func TestFindSubstring(t *testing.T) {
	for _, data := range tests {
		if result := findSubstring(data.s, data.words); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, findSubstring")
		}
	}
}
