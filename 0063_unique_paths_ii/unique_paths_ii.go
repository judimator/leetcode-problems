package _063_unique_paths_ii

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	matrix := make([][]int, m)

	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 0 {
			matrix[i][0] = 1
		} else {
			break
		}
	}
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 0 {
			matrix[0][j] = 1
		} else {
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				matrix[i][j] = matrix[i-1][j] + matrix[i][j-1]
			}
		}
	}

	return matrix[m-1][n-1]
}
