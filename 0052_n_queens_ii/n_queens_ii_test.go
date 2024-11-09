package n_queens__ii

import (
	"reflect"
	"testing"
)

type TestData struct {
	n      int
	output int
}

var tests = []TestData{
	{
		4,
		2,
	},
}

func TestSolveNQueens(t *testing.T) {
	for _, data := range tests {
		if result := totalNQueens(data.n); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, totalNQueens")
		}
	}
}
