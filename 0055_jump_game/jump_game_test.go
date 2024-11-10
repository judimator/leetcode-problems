package jump_game

import (
	"testing"
)

type TestData struct {
	nums   []int
	output bool
}

var tests = []TestData{
	{
		[]int{2, 3, 1, 1, 4},
		true,
	},
	{
		[]int{3, 2, 1, 0, 4},
		false,
	},
	{
		[]int{0},
		true,
	},
	{
		[]int{0, 1},
		false,
	},
	{
		[]int{2, 0, 0},
		true,
	},
}

func TestCanJump(t *testing.T) {
	for _, data := range tests {
		if result := canJump(data.nums); result != data.output {
			t.Error("Failed, canJump")
		}
	}
}
