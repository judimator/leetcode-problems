package reverse_integer

import "math"

func reverse(x int) int {
	outputSlice := make([]int, 0)
	tmpVar := x

	if x < 0 {
		tmpVar = tmpVar * (-1)
	}

	for {
		rightMost := tmpVar % 10
		outputSlice = append(outputSlice, rightMost)
		tmpVar = tmpVar / 10

		if tmpVar <= 0 {
			break
		}
	}

	result := 0
	i := len(outputSlice)
	z := 0

	for _, val := range outputSlice {
		if val == 0 && z == 0 {
			i--

			continue
		}

		tmp := 1

		for j := 1; j < i; j++ {
			tmp = tmp * 10
		}

		result += val * tmp
		i--
		z++
	}

	if result > math.MaxInt32 {
		return 0
	}

	if x < 0 {
		result = result * (-1)
	}

	return result
}
