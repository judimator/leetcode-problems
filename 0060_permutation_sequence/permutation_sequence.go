package permutation_sequence

import "strconv"

func getPermutation(n int, k int) string {
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = i + 1
	}

	factorials := make([]int, n+1)
	for i := 0; i <= n; i++ {
		factorials[i] = factorial(i)
	}

	result := ""
	k--

	for n > 0 {
		index := k / factorials[n-1]
		result += strconv.Itoa(numbers[index])
		numbers = append(numbers[:index], numbers[index+1:]...)

		k %= factorials[n-1]
		n--
	}

	return result
}

func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
