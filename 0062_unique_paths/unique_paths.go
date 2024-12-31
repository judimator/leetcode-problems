package unique_paths

func uniquePaths2(m int, n int) int {
	result := 0
	var backtrace func(currM, currN int)

	backtrace = func(currM, currN int) {
		if currM >= m || currN >= n {
			return
		}

		if m-1 == currM && n-1 == currN {
			result++

			return
		}

		backtrace(currM+1, currN)
		backtrace(currM, currN+1)
	}

	backtrace(0, 0)

	return result
}

func uniquePaths(m int, n int) int {
	matrix := make([][]int, m)

	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		matrix[i][0] = 1
	}

	for i := 0; i < n; i++ {
		matrix[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			matrix[i][j] = matrix[i-1][j] + matrix[i][j-1]
		}
	}

	return matrix[m-1][n-1]
}
