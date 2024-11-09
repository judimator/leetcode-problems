package powx_n

import (
	"testing"
)

type TestData struct {
	x      float64
	n      int
	output float64
}

var tests = []TestData{
	{
		2.00000,
		10,
		1024.00000,
	},
}

func TestMyPow(t *testing.T) {
	for _, data := range tests {
		if result := myPow(data.x, data.n); result != data.output {
			t.Error("Failed, myPow")
		}
	}
}
