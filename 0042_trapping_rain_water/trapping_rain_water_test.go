package trapping_rain_water

import (
	"reflect"
	"testing"
)

type TestData struct {
	input  []int
	output int
}

var tests = []TestData{
	{
		[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		6,
	},
	{
		[]int{4, 2, 0, 3, 2, 5},
		9,
	},
}

func TestTrap(t *testing.T) {
	for _, data := range tests {
		if result := trap(data.input); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, trap")
		}
	}
}
