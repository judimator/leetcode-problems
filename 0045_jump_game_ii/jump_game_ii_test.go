package jump_game_ii

import (
	"testing"
)

type TestData struct {
	nums   []int
	output int
}

var tests = []TestData{
	{
		[]int{1, 1, 0, 0, 4},
		2,
	},
}

func TestJump(t *testing.T) {
	for _, data := range tests {
		if result := jump(data.nums); result != data.output {
			t.Error("Failed, jump")
		}
	}
}
