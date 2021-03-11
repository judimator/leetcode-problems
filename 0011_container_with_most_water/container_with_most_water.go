package container_with_most_water

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	mArea := 0

	for left < right {
		var minHeight int

		if height[left] < height[right] {
			minHeight = height[left]
		} else {
			minHeight = height[right]
		}

		if mArea < (minHeight * (right - left)) {
			mArea = minHeight * (right - left)
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return mArea
}
