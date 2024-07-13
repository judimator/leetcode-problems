package divide_two_integers

import "math"

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	var sign bool

	if (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0) {
		sign = true
	} else {
		sign = false
	}

	dividend = abs(dividend)
	divisor = abs(divisor)

	result := 0

	for dividend >= divisor {
		tmp := divisor
		count := 1

		for dividend >= tmp<<1 {
			tmp = tmp << 1
			count = count << 1
		}

		result = result + count
		dividend = dividend - tmp
	}

	if sign {
		return result
	}

	return result * -1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
