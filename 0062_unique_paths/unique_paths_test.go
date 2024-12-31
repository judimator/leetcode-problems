package unique_paths

import (
	"testing"
)

type TestData struct {
	m      int
	n      int
	output int
}

var tests = []TestData{
	{
		3,
		7,
		28,
	},
	{
		3,
		2,
		3,
	},
}

func TestUniquePaths(t *testing.T) {
	for _, data := range tests {
		if result := uniquePaths(data.m, data.n); result != data.output {
			t.Error("Failed, uniquePaths")
		}

		if result := uniquePaths2(data.m, data.n); result != data.output {
			t.Error("Failed, uniquePaths2")
		}
	}
}
