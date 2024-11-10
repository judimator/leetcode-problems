package jump_game

func canJump(nums []int) bool {
	currentPos := 0
	distance := 0

	for i, val := range nums {
		distance = max(distance, val+i)

		if i == currentPos {
			currentPos = distance
		}
	}

	return (currentPos == 0 && len(nums) == 1) || currentPos >= len(nums)-1
}
