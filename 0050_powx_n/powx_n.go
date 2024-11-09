package powx_n

func myPow(x float64, n int) float64 {
	var tmp float64
	var result float64

	if n < 0 {
		x = 1 / x
		n = -n
	}

	result = 1
	tmp = x

	for n > 0 {
		if n%2 == 1 {
			result = result * tmp
		}

		tmp = tmp * tmp
		n = n / 2
	}

	return result
}
