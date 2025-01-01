package _063_unique_paths_ii

import (
	"testing"
)

type TestData struct {
	obstacleGrid [][]int
	n            int
}

var tests = []TestData{
	//{
	//	[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
	//	2,
	//},
	{
		[][]int{{0, 1}, {0, 0}},
		1,
	},
}

func TestUniquePathsWithObstacles(t *testing.T) {
	for _, data := range tests {
		if result := uniquePathsWithObstacles(data.obstacleGrid); result != data.n {
			t.Error("Failed, uniquePathsWithObstacles")
		}
	}
}
