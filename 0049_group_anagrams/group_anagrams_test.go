package group_anagrams

import (
	"reflect"
	"testing"
)

type TestData struct {
	input  []string
	output [][]string
}

var tests = []TestData{
	{
		[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
		[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
	},
}

func TestGroupAnagrams(t *testing.T) {
	for _, data := range tests {
		if result := groupAnagrams(data.input); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, groupAnagrams")
		}
	}
}
