package find_the_index_of_the_first_occurrence_in_a_string

import "testing"

type TestData struct {
	haystack string
	needle   string
	output   int
}

var tests = []TestData{
	{
		"sadbutsad",
		"sad",
		0,
	},
	{
		"mississippi",
		"issip",
		4,
	},
	{
		"aaa",
		"aaaa",
		-1,
	},
}

func TestStrStr(t *testing.T) {
	for _, data := range tests {
		if result := strStr(data.haystack, data.needle); result != data.output {
			t.Error("Failed, strStr")
		}
	}
}
