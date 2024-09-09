package trapping_rain_water

func trap(height []int) int {
	l := len(height)
	left := make([]int, l)
	right := make([]int, l)

	left[0] = height[0]
	right[l-1] = height[l-1]

	for i := 1; i < l; i++ {
		left[i] = max(left[i-1], height[i])
	}

	for i := l - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}

	result := 0

	for i := 0; i < l; i++ {
		result = result + (min(left[i], right[i]) - height[i])
	}

	return result
}
