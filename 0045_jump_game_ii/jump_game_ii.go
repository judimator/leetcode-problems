package jump_game_ii

import "math"

func jump(nums []int) int {
	jumps := 0
	maxDistance := 0
	currentPos := 0

	for i := 0; i < len(nums)-1; i++ {
		maxDistance = int(math.Max(float64(maxDistance), float64(nums[i]+i)))

		if currentPos == i {
			jumps++
			currentPos = maxDistance
		}
	}

	return jumps
}
