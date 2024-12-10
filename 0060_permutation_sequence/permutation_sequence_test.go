package permutation_sequence

import (
	"testing"
)

type TestData struct {
	n      int
	k      int
	output string
}

var tests = []TestData{
	{
		3,
		3,
		"213",
	},
}

func TestGetPermutation(t *testing.T) {
	for _, data := range tests {
		if result := getPermutation(data.n, data.k); result != data.output {
			t.Error("Failed, getPermutation")
		}
	}
}
