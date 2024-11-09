package n_queens

import (
	"reflect"
	"testing"
)

type TestData struct {
	n      int
	output [][]string
}

var tests = []TestData{
	{
		4,
		[][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}},
	},
}

func TestSolveNQueens(t *testing.T) {
	for _, data := range tests {
		if result := solveNQueens(data.n); !reflect.DeepEqual(result, data.output) {
			t.Error("Failed, solveNQueens")
		}
	}
}
